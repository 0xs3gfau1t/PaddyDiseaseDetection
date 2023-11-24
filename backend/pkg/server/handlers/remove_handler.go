package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"
	"segFault/PaddyDiseaseDetection/types"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RemoveIdentifiedDiseaseHandler(c echo.Context) error {
	user, ok := c.Get("user").(types.AuthenticatedUserRequestValues)
	if !ok {
		return c.JSON(http.StatusUnauthorized, &NoUserReturn{
			Error: "Couldn't find user info in request",
		})
	}

	var requestBody struct {
		Id uuid.UUID `form:"id"`
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return client.Cli.IdentifiedDiseases.RemoveIdentifiedDisease(requestBody.Id, user.Id)
}
