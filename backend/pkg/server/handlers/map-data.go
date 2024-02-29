package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"
	"segFault/PaddyDiseaseDetection/types"

	"github.com/labstack/echo/v4"
)

// {latitude: string, longitude: string, weight: number}[]
type MapDataResponse struct {
	Data []types.HeatMapEntry `json:"data"`
}

func GetMapDataHandler(c echo.Context) error {
	data, err := client.Cli.IdentifiedDiseases.GetMapEntries()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, MapDataResponse{Data: data})
}
