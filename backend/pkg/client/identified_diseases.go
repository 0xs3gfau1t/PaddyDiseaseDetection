package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"segFault/PaddyDiseaseDetection/constants"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/ent/diseaseidentified"
	"segFault/PaddyDiseaseDetection/ent/user"

	"segFault/PaddyDiseaseDetection/pkg/helpers"
	"segFault/PaddyDiseaseDetection/pkg/location"
	"segFault/PaddyDiseaseDetection/pkg/storage"
	"segFault/PaddyDiseaseDetection/types"
	"sync"

	"github.com/google/uuid"
)

type IdentifiedDiseasesClient interface {
	UploadImage(*multipart.FileHeader, *uuid.UUID, chan ResultChannel, *sync.WaitGroup)
	UploadImages(*types.ImageUploadType, *uuid.UUID, *http.Request) error
	RollbackImageUploads([]string) error
	RemoveIdentifiedDisease(uuid.UUID, uuid.UUID) error
	GetUploads(*uuid.UUID) ([]*types.UploadListItemType, error)
	GetUpload(*uuid.UUID, *uuid.UUID) (*types.UploadedEntity, error)
	GetUploadStat(*uuid.UUID, *uuid.UUID) (*types.UploadListItemType, error)
	GetMapEntries() ([]types.HeatMapEntry, error)
}

type IdentifiedDiseases struct {
	dbDiseaseIdentified *ent.DiseaseIdentifiedClient
	dbImage             *ent.ImageClient
	storage             storage.Storage
	rabbitPublisher     func(string) error
	// TODO: Maintain a channel or queue where failed db inserts are retried
	// db_insert_failed_channel chan DbEntryType
}

type ResultChannel struct {
	Filename string
	Error    error
}

type DbEntryType struct {
	EntryDiseaseIdentified *ent.DiseaseIdentifiedCreate
	EntryImages            *ent.ImageCreateBulk
	Images                 []string
	Tries                  uint8
	MaxTries               uint8
}

func (idiseaseCli IdentifiedDiseases) UploadImage(image *multipart.FileHeader, userid *uuid.UUID, resultChan chan ResultChannel, wg *sync.WaitGroup) {
	defer wg.Done()

	if helpers.ToMB(image.Size) > constants.MAX_FILE_UPLOAD_LIMIT {
		resultChan <- ResultChannel{
			Filename: image.Filename,
			Error:    fmt.Errorf("Image size greater than limit of 5MB."),
		}
	}
	opened, err := image.Open()
	defer opened.Close()
	if err != nil {
		resultChan <- ResultChannel{
			Filename: image.Filename,
			Error:    fmt.Errorf("Couldn't read image"),
		}
		return
	}

	file := bytes.NewBuffer(nil)
	if _, err := io.Copy(file, opened); err != nil {
		resultChan <- ResultChannel{
			Filename: image.Filename,
			Error:    fmt.Errorf("Couldn't write to buffer"),
		}
		return
	}

	fileName, err := helpers.IsFileAllowed(image.Filename, file)
	if err != nil {
		resultChan <- ResultChannel{
			Filename: fileName,
			Error:    err,
		}
		return
	}
	fileName = fmt.Sprintf("%v/%v", userid, fileName)
	resultChan <- ResultChannel{
		Filename: fileName,
		Error:    idiseaseCli.storage.AddFile(fileName, file.Bytes()),
	}

}

func (idiseaseCli IdentifiedDiseases) UploadImages(images *types.ImageUploadType, userid *uuid.UUID, request *http.Request) error {
	resultChan := make(chan ResultChannel)
	var wg sync.WaitGroup
	for _, image := range images.Images {
		wg.Add(1)
		go func(uncapturedImage *multipart.FileHeader) {
			idiseaseCli.UploadImage(uncapturedImage, userid, resultChan, &wg)
		}(image)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	successfulUploads := make([]string, 0, len(images.Images))
	for result := range resultChan {
		if result.Error == nil {
			successfulUploads = append(successfulUploads, result.Filename)
		} else {
			log.Printf("[x] Failed uploading %v\n", result.Filename)
			log.Println(result.Error)
		}
	}
	if len(successfulUploads) == 0 {
		return types.ErrUploadFailed
	}
	newDbEnteryId := uuid.New()
	loc := location.GetLocation(images, request, *userid, Cli.db)
	dbEntry := &DbEntryType{
		EntryDiseaseIdentified: idiseaseCli.dbDiseaseIdentified.Create().SetID(newDbEnteryId).SetStatus("queued").SetLocation(loc).SetUploadedByID(*userid),
		EntryImages: idiseaseCli.dbImage.MapCreateBulk(successfulUploads, func(ic *ent.ImageCreate, i int) {
			ic.SetIdentifier(successfulUploads[i]).SetDiseaseIdentifiedID(newDbEnteryId)
		}),
		Images:   successfulUploads,
		Tries:    1,
		MaxTries: 5,
	}
	if err := dbEntry.EntryDiseaseIdentified.Exec(context.Background()); err != nil {
		log.Println(err)
		log.Println("[x] Inserting to db failed")
		go func() {
			// idiseaseCli.db_insert_failed_channel <- *dbEntry
			if err = idiseaseCli.RollbackImageUploads(successfulUploads); err != nil {
				// Images in supabase are now orphaned i.e no record on db
				// The only way they can now be removed is either
				// 1. Remove all user UploadImages
				// 2. Fetch list of all uploads -> compare with records on db -> remove those not in db
			}
		}()
		return err
	}

	// TODO: Use transactions for both of them
	if err := dbEntry.EntryImages.Exec(context.Background()); err != nil {
		if err := idiseaseCli.dbDiseaseIdentified.DeleteOneID(newDbEnteryId).Exec(context.Background()); err == nil {
			idiseaseCli.RollbackImageUploads(successfulUploads)
		} else {
			log.Printf("[x] Failed removing user submitted job %v\n", newDbEnteryId)
		}
	}

	// Publish a message for ML workers to work on identification
	if idiseaseCli.rabbitPublisher == nil {
		log.Println("[!] No publisher found")
		return types.ErrPublishFailed
	}
	if signedUrl, err := idiseaseCli.storage.GetFilePath(successfulUploads[0]); err == nil {
		if publishMsg, err := json.Marshal(types.PublishMessage{
			Id:   newDbEnteryId.String(),
			Link: signedUrl,
		}); err == nil {
			if idiseaseCli.rabbitPublisher(string(publishMsg)) != nil {
				return types.ErrPublishFailed
			}
		} else {
			return types.ErrPublishFailed
		}
	} else {
		return types.ErrPublishFailed
	}

	if len(successfulUploads) < len(images.Images) {
		return types.ErrUploadFailedPartial
	}
	return nil
}

func (idiseaseCli IdentifiedDiseases) RollbackImageUploads(images []string) error {
	log.Println("[i] Rolling back uploaded images")
	var wg sync.WaitGroup

	rollbackChan := make(chan error)
	defer close(rollbackChan)

	for _, image := range images {
		wg.Add(1)
		go func(image string) {
			defer wg.Done()
			if err := idiseaseCli.storage.RemoveFile(image); err != nil {
				rollbackChan <- err
			}
		}(image)
	}

	ctx, done := context.WithCancel(context.Background())
	go func(cancel *context.CancelFunc) {
		wg.Wait()
		done()
	}(&done)

	for {
		select {
		case err := <-rollbackChan:
			return err
		case <-ctx.Done():
			return nil
		}
	}
}

// Deletes entry on RemoveIdentifiedDisease table
// Fails if the ml model is still processing on this item
func (idiseaseCli IdentifiedDiseases) RemoveIdentifiedDisease(id uuid.UUID, user_id uuid.UUID) error {
	identifiedInstance, err := idiseaseCli.dbDiseaseIdentified.Query().Unique(true).Where(diseaseidentified.ID(id)).Where(diseaseidentified.HasUploadedByWith(user.ID(user_id))).Where(diseaseidentified.StatusNEQ("processing")).First(context.Background())
	if err != nil {
		log.Printf("[x] No entry found for %v. Maybe it's still being processed\n", id)

		// Here is a weird design decision TODO. This error is pure http.StatusLocked or other
		// Now how do we propel this error to response handler
		// Because currently we're sending 500 to any error that is not being explicitly  handled from handlers
		// i.e anything that raises error from client module is being treated as 500 but it's got more than that
		return err
	}

	photos, err := identifiedInstance.QueryImage().All(context.Background())
	photosStr := make([]string, len(photos))
	for i, photo := range photos {
		photosStr[i] = photo.Identifier
	}
	if err = idiseaseCli.RollbackImageUploads(photosStr); err == nil {
		if err = idiseaseCli.dbDiseaseIdentified.DeleteOne(identifiedInstance).Exec(context.Background()); err != nil {
			log.Printf("[x] Couldn't delete %v. DBERROR\n", id)
		}
	}

	if err != nil {
		log.Printf("[x] Unable to delete disease_identified entry: %v\n", id)
		return err
	}
	return nil
}

func (idiseaseCli IdentifiedDiseases) GetUploads(user_id *uuid.UUID) ([]*types.UploadListItemType, error) {
	diseases, err := idiseaseCli.dbDiseaseIdentified.Query().WithDisease().WithImage().Where(diseaseidentified.HasUploadedByWith(user.ID(*user_id))).All(context.Background())
	if err != nil {
		return nil, err
	}

	var cleanedUploads []*types.UploadListItemType

	for _, uploadItem := range diseases {

		imageLinks := make([]string, 0)
		for _, image := range uploadItem.Edges.Image {
			if imgLink, err := idiseaseCli.storage.GetFilePath(image.Identifier); err == nil {
				imageLinks = append(imageLinks, imgLink)
			}
		}

		diseaseName := make([]string, 0)
		for _, disease := range uploadItem.Edges.Disease {
			diseaseName = append(diseaseName, disease.Name)
		}

		cleanedUploads = append(cleanedUploads, &types.UploadListItemType{
			Id:       uploadItem.ID.String(),
			Name:     diseaseName,
			Severity: uploadItem.Severity,
			Status:   uploadItem.Status.String(),
			Images:   imageLinks,
		})
	}
	return cleanedUploads, nil
}

func (idiseaseCli IdentifiedDiseases) GetUploadStat(user_id *uuid.UUID, uploadId *uuid.UUID) (*types.UploadListItemType, error) {
	diseases, err := idiseaseCli.dbDiseaseIdentified.Query().WithDisease().WithImage().Where(diseaseidentified.HasUploadedByWith(user.ID(*user_id))).Where(diseaseidentified.ID(*uploadId)).First(context.Background())
	if err != nil {
		return nil, err
	}

	imageLinks := make([]string, 0)
	for _, image := range diseases.Edges.Image {
		if imgLink, err := idiseaseCli.storage.GetFilePath(image.Identifier); err == nil {
			imageLinks = append(imageLinks, imgLink)
		}
	}

	name := make([]string, 0)

	for _, disease := range diseases.Edges.Disease {
		name = append(name, disease.Name)
	}

	cleanedUploads := types.UploadListItemType{
		Id:       diseases.ID.String(),
		Name:     name,
		Status:   diseases.Status.String(),
		Severity: diseases.Severity,
		Images:   imageLinks,
	}
	return &cleanedUploads, nil
}

func (idiseaseCli IdentifiedDiseases) GetUpload(user_id *uuid.UUID, uploadId *uuid.UUID) (*types.UploadedEntity, error) {
	diseases, err := idiseaseCli.dbDiseaseIdentified.Query().WithDisease().WithImage().Where(diseaseidentified.HasUploadedByWith(user.ID(*user_id))).Where(diseaseidentified.ID(*uploadId)).First(context.Background())
	if err != nil {
		return nil, err
	}

	imageLinks := make([]string, 0)
	for _, image := range diseases.Edges.Image {
		if imgLink, err := idiseaseCli.storage.GetFilePath(image.Identifier); err == nil {
			imageLinks = append(imageLinks, imgLink)
		}
	}

	identified := make([]*types.IdentifiedDiseaseEntity, 0)
	name := make([]string, 0)

	for _, disease := range diseases.Edges.Disease {
		tmpSolution := make([]*types.SolutionEntity, 0)
		if solutionsFromDB, err := disease.QuerySolutions().All(context.Background()); err == nil {
			for _, solutionFromDB := range solutionsFromDB {
				solImages := make([]string, 0)
				for _, image := range solutionFromDB.Photos {
					if imgLink, err := idiseaseCli.storage.GetFilePath(image); err == nil {
						solImages = append(solImages, imgLink)
					}
				}
				tmpSolution = append(tmpSolution, &types.SolutionEntity{
					Id:          solutionFromDB.ID.String(),
					Name:        solutionFromDB.Name,
					Photos:      solImages,
					Description: solutionFromDB.Description,
					Ingredients: solutionFromDB.Ingredients,
				})
			}
		}
		tmpCauses := make([]*types.CausesEntity, 0)
		if causesFromDB, err := disease.QueryCauses().All(context.Background()); err == nil {
			for _, causeFromDB := range causesFromDB {
				if img, err := idiseaseCli.storage.GetFilePath(causeFromDB.Image); err == nil {
					tmpCauses = append(tmpCauses, &types.CausesEntity{
						Id:    causeFromDB.ID.String(),
						Name:  causeFromDB.Name,
						Image: img,
					})
				}
			}

		}
		name = append(name, disease.Name)
		c, err := disease.QueryCauses().All(context.Background())
		var causes []*types.CauseJson
		if err == nil {
			for _, cause := range c {
				imagePath, err := idiseaseCli.storage.GetFilePath(cause.Image)
				if err != nil {
					imagePath, _ = idiseaseCli.storage.GetFilePath("default.jpeg")
				}
				causes = append(causes, &types.CauseJson{
					Name:  cause.Name,
					Image: imagePath,
				})
			}
		}

		identified = append(identified, &types.IdentifiedDiseaseEntity{
			Name:      disease.Name,
			Id:        disease.ID.String(),
			Solutions: tmpSolution,
			Causes:    tmpCauses,
		})
	}

	cleanedUploads := types.UploadedEntity{
		Id:         diseases.ID.String(),
		Name:       name,
		Status:     diseases.Status.String(),
		Severity:   diseases.Severity,
		Images:     imageLinks,
		Identified: identified,
		Roi:        diseases.Roi,
	}
	return &cleanedUploads, nil
}

func (idiseaseCli IdentifiedDiseases) GetMapEntries() ([]types.HeatMapEntry, error) {
	data, err := idiseaseCli.dbDiseaseIdentified.Query().Where(diseaseidentified.StatusEQ("processed")).All(context.Background())
	if err != nil {
		return nil, err
	}

	var response []types.HeatMapEntry
	for _, entry := range data {
		latitude, longitude, _ := location.ParseDbLocation(entry.Location)
		response = append(response, types.HeatMapEntry{
			Id:        entry.ID.String(),
			Latitude:  latitude.ToFloat(),
			Longitude: longitude.ToFloat(),
			Weight:    entry.Severity,
		})
	}

	return response, nil
}
