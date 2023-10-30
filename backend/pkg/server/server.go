package server

import (
	"segFault/PaddyDiseaseDetection/pkg/server/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(port string) error {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	// Register /api routes
        api.InitApiRoutes(e)

	return e.Start(":" + port)
}
