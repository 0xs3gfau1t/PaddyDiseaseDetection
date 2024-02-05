package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/pkg/client"
	"segFault/PaddyDiseaseDetection/types"

	"github.com/labstack/echo/v4"
)

type UploadsResponse struct {
	Uploads []*ent.DiseaseIdentified `json:"uploads"`
}

func GetUploadsHandler(c echo.Context) error {
	user, ok := c.Get("user").(types.AuthenticatedUserRequestValues)
	if !ok {
		return c.JSON(http.StatusUnauthorized, &NoUserReturn{
			Error: "Couldn't find user info in request",
		})
	}

	uploads, err := client.Cli.IdentifiedDiseases.GetUploads(&user.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &UploadsResponse{
		Uploads: uploads,
	})
}
