package types

import (
	"mime/multipart"

	"github.com/golang-jwt/jwt"
)

type CreateUserValidInput struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Location string `form:"location"`
	Coord    string `form:"coord"`
	Password string `form:"password"`
}

type LoginUserValidInput struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type AuthenticatedUserRequestValues struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}

type JwtType struct {
	jwt.StandardClaims
	AuthenticatedUserRequestValues
}

type ImageUploadType struct {
	Images []*multipart.FileHeader `json:"images"`
}
