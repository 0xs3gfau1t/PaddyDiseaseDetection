package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"segFault/PaddyDiseaseDetection/constants"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/ent/diseaseidentified"
	"segFault/PaddyDiseaseDetection/ent/user"

	// "segFault/PaddyDiseaseDetection/ent/user"
	"segFault/PaddyDiseaseDetection/pkg/helpers"
	"segFault/PaddyDiseaseDetection/pkg/storage"
	"segFault/PaddyDiseaseDetection/types"
	"sync"

	"github.com/google/uuid"
)

type IdentifiedDiseasesClient interface {
	UploadImage(*multipart.FileHeader, string, chan ResultChannel, *sync.WaitGroup)
	UploadImages(*types.ImageUploadType, string) error
	RollbackImageUploads([]string) error
	RemoveIdentifiedDisease(uuid.UUID, string) error
}

type IdentifiedDiseases struct {
	db      *ent.DiseaseIdentifiedClient
	storage storage.Storage
	// TODO: Maintain a channel or queue where failed db inserts are retried
	// db_insert_failed_channel chan DbEntryType
}

type ResultChannel struct {
	Filename string
	Error    error
}

type DbEntryType struct {
	Entry    *ent.DiseaseIdentifiedCreate
	Images   []string
	Tries    uint8
	MaxTries uint8
}

func (idiseaseCli IdentifiedDiseases) UploadImage(image *multipart.FileHeader, userid string, resultChan chan ResultChannel, wg *sync.WaitGroup) {
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

func (idiseaseCli IdentifiedDiseases) UploadImages(images *types.ImageUploadType, userid string) error {
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
	userId, _ := uuid.Parse(userid)
	dbEntry := &DbEntryType{
		Entry:    idiseaseCli.db.Create().SetID(uuid.New()).SetPhotos(successfulUploads).SetStatus("queued").SetLocation("Nepal").SetUploadedByID(userId),
		Images:   successfulUploads,
		Tries:    1,
		MaxTries: 5,
	}
	if err := dbEntry.Entry.Exec(context.Background()); err != nil {
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
func (idiseaseCli IdentifiedDiseases) RemoveIdentifiedDisease(id uuid.UUID, user_id string) error {
	userId, _ := uuid.Parse(user_id)
	photos, err := idiseaseCli.db.Query().Unique(true).Where(diseaseidentified.ID(id)).Where(diseaseidentified.HasUploadedByWith(user.ID(userId))).Where(diseaseidentified.StatusNEQ("processing")).First(context.Background())
	log.Println(photos)
	if err != nil {
		log.Printf("[x] No entry found for %v. Maybe it's still being processed", id)

		// Here is a weird design decision TODO. This error is pure http.StatusLocked or other
		// Now how do we propel this error to response handler
		// Because currently we're sending 500 to any error that is not being explicitly  handled from handlers
		// i.e anything that raises error from client module is being treated as 500 but it's got more than that
		return err
	}
	err = idiseaseCli.db.DeleteOne(photos).Exec(context.Background())

	if err != nil {
		log.Printf("[x] Unable to delete disease_identified entry: %v\n", id)
		return err
	}
	go func(photos []string) {
		idiseaseCli.RollbackImageUploads(photos)
	}(photos.Photos)
	return nil
}
