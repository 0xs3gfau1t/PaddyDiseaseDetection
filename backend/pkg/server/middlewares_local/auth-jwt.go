package middlewareslocal

import (
	"log"
	"net/http"
	"os"
	"segFault/PaddyDiseaseDetection/types"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := c.Cookie("accessToken")
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "No authorization token found")
		}

		token, err := jwt.ParseWithClaims(tokenString.Value, &types.JwtType{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SIGNING_SECRET")), nil
		})
		claim, ok := token.Claims.(*types.JwtType)
		if err != nil || !token.Valid || !ok {
			log.Println(err, ok)
			return c.JSON(http.StatusOK, map[string]string{
				"error": "Access denied. Invalid token",
			})
		}
		c.Set("user", claim.AuthenticatedUserRequestValues)
		return next(c)
	}
}
