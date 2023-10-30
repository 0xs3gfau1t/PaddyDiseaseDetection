package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
}

func InitApiRoutes(e *echo.Echo) {
	e.GET("/api", func(c echo.Context) error {
		res := &HomeResponse{
			Message: "Oniiii san, you hit a kawaiii endpoint.",
		}
		return c.JSONPretty(http.StatusOK, res, " ")
	})
}
