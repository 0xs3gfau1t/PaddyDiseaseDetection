package location

import (
	"context"
	"fmt"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/ent/user"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type LocationExtractorFromUser struct {
	Userid uuid.UUID
	Db     *ent.Client
}

func (le *LocationExtractorFromUser) GetLocation() (Location, Location, error) {
	emptyLocation := Location{}

	user, err := le.Db.User.Query().Unique(true).Where(user.ID(le.Userid)).Select(user.FieldLocation).First(context.Background())
	if err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("User not found:%w", err)
	}

	parts := strings.Split(user.Location, " ")
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
