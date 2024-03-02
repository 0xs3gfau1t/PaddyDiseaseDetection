package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(port string) error {
	e := echo.New()
	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "[${method}] (${status}): ${uri} ${error}\n",
		}),
		middleware.Recover(),
		middleware.CORS(),
	)

	// Register /api routes
	InitApiRoutes(e)
	return e.Start(":" + port)
}
