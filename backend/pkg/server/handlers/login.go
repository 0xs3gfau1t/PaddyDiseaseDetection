package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"

	"github.com/labstack/echo/v4"
)

func LoginHandler(cli *client.Client) echo.HandlerFunc {
	type LoginReturn struct {
		AccessToken string `json:"accessToken"`
	}
	return func(c echo.Context) error {
		var formattedInput client.LoginUserValidInput
		c.Bind(&formattedInput)
		jwt, err := cli.User.Login(&formattedInput)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, &LoginReturn{
			AccessToken: jwt,
		})
	}
}
