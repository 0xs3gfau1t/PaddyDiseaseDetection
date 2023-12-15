package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"
	"segFault/PaddyDiseaseDetection/types"

	"github.com/labstack/echo/v4"
)

func SignUpHandler(c echo.Context) error {
	var formattedInput types.CreateUserValidInput
	c.Bind(&formattedInput)
	_, err := client.Cli.User.CreateUser(&formattedInput)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
