package location

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/types"
	"strconv"
	"strings"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

var fc = geojson.NewFeatureCollection()
var c jsoniter.API

func init() {
	f, err := os.Open("pkg/location/nepal.geojson")
	if err != nil {
		log.Fatal(err)
		return
	}
	geojsonData, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	c = jsoniter.Config{
		EscapeHTML:              true,
		SortMapKeys:             false,
		MarshalFloatWith6Digits: true,
	}.Froze()

	fc = geojson.NewFeatureCollection()
	if err = c.Unmarshal(geojsonData, &fc); err != nil {
		log.Fatal(err)
	}
}

type LocationExtractor interface {
	GetLocation() (Location, Location, error) // longitude latitude error
}

type Location struct {
	Degrees   int
	Minutes   int
	Seconds   float64
	Direction string
}

func (loc *Location) ToFloat() float64 {
	numericValue := float64(loc.Degrees) + float64(loc.Minutes)/60 + loc.Seconds/60

	if strings.ToUpper(loc.Direction) == "S" || strings.ToUpper(loc.Direction) == "W" {
		numericValue = -numericValue
	}

	return numericValue
}

func (loc *Location) FromFloat(numericValue float64, isLatitude bool) {
	if isLatitude {
		if numericValue > 0 {
			loc.Direction = "N"
		} else {
			loc.Direction = "S"
		}
	} else {
		if numericValue > 0 {
			loc.Direction = "E"
		} else {
			loc.Direction = "W"
		}
	}

	numericValue = math.Abs(numericValue)

	loc.Degrees = int(numericValue)
	numericValue -= float64(loc.Degrees)

	loc.Minutes = int(numericValue * 60)
	numericValue -= float64(loc.Minutes) / 60

	loc.Seconds = numericValue * 60
}

func GetDistrict(longitude float64, latitude float64) string {
	point := orb.Point{longitude, latitude}
	fmt.Println(point)

	for _, feature := range fc.Features {
		switch g := feature.Geometry.(type) {
		case orb.Polygon:
			if planar.PolygonContains(g, point) {
				x, ok := feature.Properties["DIST_EN"]
				if ok {
					return x.(string)
				} else {
					return "NOT FOUND"
				}
			}
		}
	}
	return "NOT FOUND"
}

func GetLocation(images *types.ImageUploadType, request *http.Request, userId uuid.UUID, db *ent.Client) string {
	for _, image := range images.Images {
		opened, err := image.Open()
		if err != nil {
			continue
		}
		defer opened.Close()
		fileBuffer := bytes.NewBuffer(nil)
		if _, err := io.Copy(fileBuffer, opened); err != nil {
			continue
		}
		fileBytes := fileBuffer.Bytes()
		fileLE := LocationExtractorFromFile{
			File: &fileBytes,
		}
		latitude, longitude, err := fileLE.GetLocation()
		if err == nil {
			log.Printf("[FileLocation]: %v %v\n", latitude, longitude)
			return fmt.Sprintf("%f %f", latitude.ToFloat(), longitude.ToFloat())
		}
	}

	requestLE := LocationExtractorFromRequest{
		Request: request,
	}
	latitude, longitude, err := requestLE.GetLocation()
	if err == nil {
		log.Printf("[RequestLocation]: %v %v\n", latitude, longitude)
		return fmt.Sprintf("%f %f", latitude.ToFloat(), longitude.ToFloat())
	}

	userLE := LocationExtractorFromUser{
		Userid: userId,
		Db:     db,
	}
	latitude, longitude, err = userLE.GetLocation()
	if err == nil {
		log.Printf("[UserLocation]: %v %v\n", latitude, longitude)
		return fmt.Sprintf("%f %f", latitude.ToFloat(), longitude.ToFloat())
	}

	return "0 0"
}

func ParseDbLocation(input string) (Location, Location, error) {
	emptyLocation := Location{}

	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		return emptyLocation, emptyLocation, fmt.Errorf("Invalid location format")
	}

	lat, err := strconv.ParseFloat(parts[0], len(parts[0]))
	if err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("Failed to parse latitude:%w", err)
	}

	lon, err := strconv.ParseFloat(parts[1], len(parts[1]))
	if err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("Failed to parse longitude:%w", err)
	}

	latStruct := Location{}
	latStruct.FromFloat(lat, true)

	lonStruct := Location{}
	lonStruct.FromFloat(lon, false)

	return latStruct, lonStruct, nil
}
