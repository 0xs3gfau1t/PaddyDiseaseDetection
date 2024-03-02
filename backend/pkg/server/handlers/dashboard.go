package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"
	"segFault/PaddyDiseaseDetection/types"

	"github.com/labstack/echo/v4"
)

type DashboardDataResponse struct {
	Data *types.DashboardData `json:"data"`
}

func GetDashboardHandler(c echo.Context) error {
	user, ok := c.Get("user").(types.AuthenticatedUserRequestValues)
	if !ok {
		return c.JSON(http.StatusUnauthorized, &NoUserReturn{
			Error: "Couldn't find user info in request",
		})
	}

	dashboardData, err := client.Cli.User.GetDashboardData(&user.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, DashboardDataResponse{
		Data: dashboardData,
	})
}
