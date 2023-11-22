package server

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/server/handlers"
	middlewareslocal "segFault/PaddyDiseaseDetection/pkg/server/middlewares_local"

	"github.com/labstack/echo/v4"
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

	e.POST("/api/auth/signup", handlers.SignUpHandler)
	e.POST("/api/auth/login", handlers.LoginHandler)
	e.POST("/api/auth/logout", handlers.LogoutHander)
	e.POST("/api/upload", handlers.UploadHandler, middlewareslocal.JwtMiddleware)
	// logout shouldn't be protected, remove this middleware later
}
