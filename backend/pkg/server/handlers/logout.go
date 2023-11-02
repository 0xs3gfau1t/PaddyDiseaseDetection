package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"

	"github.com/labstack/echo/v4"
)

func LogoutHandler(cli *client.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// user, ok := c.Get("user").(client.AuthenticatedUserRequestValues) // This is how you access authenticated token values

		c.SetCookie(&http.Cookie{
			Name:   "accessToken",
			MaxAge: -1,
			Value:  "",
		})
		return nil
	}
}
