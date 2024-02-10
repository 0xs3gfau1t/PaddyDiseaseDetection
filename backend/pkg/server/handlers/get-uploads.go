package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"
	"segFault/PaddyDiseaseDetection/types"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UploadsResponse struct {
	Uploads []*types.UploadedEntity `json:"data"`
}

type UploadResponse struct {
	Upload *types.UploadedEntity `json:"data"`
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

	return c.JSON(http.StatusOK, UploadsResponse{
		Uploads: uploads,
	})
}

func GetUploadHandler(c echo.Context) error {
	user, ok := c.Get("user").(types.AuthenticatedUserRequestValues)
	if !ok {
		return c.JSON(http.StatusUnauthorized, &NoUserReturn{
			Error: "Couldn't find user info in request",
		})
	}
	uploadId, err := uuid.Parse(c.QueryParam("itemId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &NoUserReturn{
			Error: "Invalid upload id",
		})
	}

	uploads, err := client.Cli.IdentifiedDiseases.GetUpload(&user.Id, &uploadId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, UploadResponse{
		Upload: uploads,
	})
}
