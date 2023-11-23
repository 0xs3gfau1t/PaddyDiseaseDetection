package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"
	"segFault/PaddyDiseaseDetection/types"

	"github.com/labstack/echo/v4"
)

type NoUserReturn struct {
	Error string `json:"error"`
}

func UploadHandler(c echo.Context) error {
	user, ok := c.Get("user").(types.AuthenticatedUserRequestValues)
	if !ok {
		return c.JSON(http.StatusUnauthorized, &NoUserReturn{
			Error: "Couldn't find user info in request",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSONBlob(http.StatusBadRequest, []byte(err.Error()))
	}

	images := types.ImageUploadType{
		Images: form.File["images"],
	}

	if err := client.Cli.IdentifiedDiseases.UploadImages(&images, user.Id); err != nil {
		return c.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
	}
	return nil
}
