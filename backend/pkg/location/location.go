package location

import (
	"math"
	"strings"
)

type LocationExtractor interface {
	GetLocation() (Location, Location, error)
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
