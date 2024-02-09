package handlers

import (
	"net/http"
	"segFault/PaddyDiseaseDetection/pkg/client"
	"segFault/PaddyDiseaseDetection/types"

	"github.com/labstack/echo/v4"
)

type UserProfileResponse struct {
	Data types.UserProfileData `json:"data"`
}

func GetProfileHandler(c echo.Context) error {
	sessionUser, ok := c.Get("user").(types.AuthenticatedUserRequestValues)
	if !ok {
		return c.JSON(http.StatusUnauthorized, &NoUserReturn{
			Error: "Couldn't find user info in request",
		})
	}

	user, err := client.Cli.User.UserDetails(sessionUser.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &UserProfileResponse{
		Data: *user,
	})

}

func EditProfileHandler(c echo.Context) error {
	sessionUser, ok := c.Get("user").(types.AuthenticatedUserRequestValues)
	if !ok {
		return c.JSON(http.StatusUnauthorized, &NoUserReturn{
			Error: "Couldn't find user info in request",
		})
	}

	var input types.UserProfileEditRequest
	c.Bind(&input)

	return client.Cli.User.UpdateUser(&sessionUser.Id, &input)
}

func ChangePassHandler(c echo.Context) error {
	sessionUser, ok := c.Get("user").(types.AuthenticatedUserRequestValues)
	if !ok {
		return c.JSON(http.StatusUnauthorized, &NoUserReturn{
			Error: "Couldn't find user info in request",
		})
	}

	var input types.ChangePassRequest
	c.Bind(&input)

	return client.Cli.User.ChangePassword(&sessionUser.Id, &input)
}

func DeleteProfileHandler(c echo.Context) error {
	sessionUser, ok := c.Get("user").(types.AuthenticatedUserRequestValues)
	if !ok {
		return c.JSON(http.StatusUnauthorized, &NoUserReturn{
			Error: "Couldn't find user info in request",
		})
	}
	return client.Cli.User.DeleteUser(&sessionUser.Id)
}
