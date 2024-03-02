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
	e.GET("/api/upload", handlers.GetUploadHandler, middlewareslocal.JwtMiddleware)
	e.GET("/api/uploadStat", handlers.GetUploadStatHandler, middlewareslocal.JwtMiddleware)
	e.GET("/api/uploads", handlers.GetUploadsHandler, middlewareslocal.JwtMiddleware)
	e.GET("/api/profile", handlers.GetProfileHandler, middlewareslocal.JwtMiddleware)
	e.PATCH("/api/profile", handlers.EditProfileHandler, middlewareslocal.JwtMiddleware)
	e.DELETE("/api/profile", handlers.DeleteProfileHandler, middlewareslocal.JwtMiddleware)
	e.POST("/api/profile/change_password", handlers.ChangePassHandler, middlewareslocal.JwtMiddleware)
	e.DELETE("/api/disease", handlers.RemoveIdentifiedDiseaseHandler, middlewareslocal.JwtMiddleware)
	e.GET("/api/heatmap", handlers.GetMapDataHandler)
	e.GET("/api/dashboard", handlers.GetDashboardHandler, middlewareslocal.JwtMiddleware)
}
