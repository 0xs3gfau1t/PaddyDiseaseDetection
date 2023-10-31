package server

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"
	"segFault/PaddyDiseaseDetection/pkg/server/handlers"

	"github.com/labstack/echo/v4"
)

type HomeResponse struct {
	Message string `json:"message"`
}

func InitApiRoutes(e *echo.Echo) {
	cli := client.New()

	e.GET("/api", func(c echo.Context) error {
		res := &HomeResponse{
			Message: "Oniiii san, you hit a kawaiii endpoint.",
		}
		return c.JSONPretty(http.StatusOK, res, " ")
	})

	e.POST("/api/auth/signup", handlers.SignUpHandler(cli))
}
