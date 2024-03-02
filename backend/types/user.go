package types

import (
	"mime/multipart"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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
	Email string    `json:"email"`
	Id    uuid.UUID `json:"id"`
}

type JwtType struct {
	jwt.StandardClaims
	AuthenticatedUserRequestValues
}

type ImageUploadType struct {
	Images []*multipart.FileHeader `json:"images"`
}

type PublishMessage struct {
	Id   string `json:"id"`
	Link string `json:"link"`
}

type FrameData struct {
	Box        []float32 `json:"box"`
	Confidence float32   `json:"conf"`
	ClassId    int       `json:"classId"`
	Name       string    `json:"name"`
	Color      string    `json:"color"`
}

type ProcessedMessage struct {
	Frames []FrameData `json:"frames"`
}

type UserProfileData struct {
	Name     string     `json:"name"`
	Image    string     `json:"image"`
	Email    string     `json:"email"`
	Verified bool       `json:"verified"`
	Location string     `json:"location"`
	Coords   UserCoords `json:"coords"`
}
type UserCoords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type UserProfileEditRequest struct {
	Name      string   `form:"name"`
	Latitude  *float64 `form:"latitude"`
	Longitude *float64 `form:"longitude"`
	Location  string   `form:"location"`
}

type ChangePassRequest struct {
	OldPassword string `form:"oldPassword"`
	NewPassword string `form:"newPassword"`
}
