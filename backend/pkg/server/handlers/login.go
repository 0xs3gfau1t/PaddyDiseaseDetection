package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"
	"time"

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

		c.SetCookie(&http.Cookie{
			Name:     "accessToken",
			Value:    jwt,
			MaxAge:   int(time.Now().Unix()) + 30*24*60*60,
			SameSite: http.SameSiteLaxMode,
		})
		return c.JSON(http.StatusOK, &LoginReturn{
			AccessToken: jwt,
		})
	}
}
