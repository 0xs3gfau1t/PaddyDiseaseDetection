package handlers

import (
	"fmt"
	"segFault/PaddyDiseaseDetection/pkg/client"

	"github.com/labstack/echo/v4"
)

func SignUpHandler(cli *client.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		var formattedInput client.CreateUserValidInput
		c.Bind(&formattedInput)
		_, err := cli.User.CreateUser(&formattedInput)
		fmt.Printf("%v", err)
		return err
	}
}
