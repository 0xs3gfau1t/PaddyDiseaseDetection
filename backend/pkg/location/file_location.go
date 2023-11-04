package location

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/barasher/go-exiftool"
)

type LocationExtractorFromFile struct {
	File []byte
}

func (le *LocationExtractorFromFile) GetLocation() (Location, Location, error) {
	emptyLocation := Location{}

	tmpFile, err := os.CreateTemp("", "location-extractor-")
	if err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("Failed to create temp file:%w", err)
	}
	filename := tmpFile.Name()
	defer tmpFile.Close()
	defer os.Remove(filename)

	if _, err := tmpFile.Write(le.File); err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("Failed to write data to temp file:%w", err)
	}

	et, err := exiftool.NewExiftool()
	if err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("Failed to instantiate exifTool:%w", err)
	}
	defer et.Close()

	fileInfo := et.ExtractMetadata(filename)[0]
	if fileInfo.Err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("Failed to extract metadata:%w", err)
	}

	longitudeStr := fileInfo.Fields["GPSLatitude"].(string)
	latitudeStr := fileInfo.Fields["GPSLongitude"].(string)

	regexPattern := `(\d+) deg (\d+)' ([\d.]+)" ([NSEW])`
	re := regexp.MustCompile(regexPattern)

	latitudeMatch := re.FindStringSubmatch(latitudeStr)
	longitudeMatch := re.FindStringSubmatch(longitudeStr)

	if latitudeMatch == nil || longitudeMatch == nil {
		return emptyLocation, emptyLocation, fmt.Errorf("Invalid GPS location format:%w", err)
	}

	latitudeLocation := Location{
		Degrees:   parseInt(latitudeMatch[1]),
		Minutes:   parseInt(latitudeMatch[2]),
		Seconds:   parseFloat(latitudeMatch[3]),
		Direction: latitudeMatch[4],
	}
	longitudeLocation := Location{
		Degrees:   parseInt(longitudeMatch[1]),
		Minutes:   parseInt(longitudeMatch[2]),
		Seconds:   parseFloat(longitudeMatch[3]),
		Direction: longitudeMatch[4],
	}

	fmt.Println(latitudeLocation, longitudeLocation)

	return latitudeLocation, longitudeLocation, nil
}

func parseInt(s string) int {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func parseFloat(s string) float64 {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
