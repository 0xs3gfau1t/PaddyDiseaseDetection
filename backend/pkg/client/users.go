package client

import (
	"context"
	"errors"
	"os"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/ent/user"
	"time"
	"unicode/utf8"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

type JwtType struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	IssuedAt  string `json:"iat"`
	ExpiresAt string `json:"exp"`
}

type UserClient interface {
	UserDetails(id uuid.UUID) (*ent.User, error)
	CreateUser(validatedUser *CreateUserValidInput) (*ent.User, error)
	HashPassword(unhashed string) ([]byte, error)
	Login(validatedInput *LoginUserValidInput) (string, error)
}

type usercli struct {
	db *ent.UserClient
}

func (u usercli) UserDetails(id uuid.UUID) (*ent.User, error) {
	return u.db.Get(context.Background(), id)
}

func (u usercli) CreateUser(validatedUser *CreateUserValidInput) (*ent.User, error) {
	toBeInsertedUser := u.db.Create()
	toBeInsertedUser.SetID(uuid.New()) // This line might cause server to crash however unlikely
	toBeInsertedUser.SetName(validatedUser.Name)
	toBeInsertedUser.SetLocation(validatedUser.Location)
	toBeInsertedUser.SetEmail(validatedUser.Email)

	hashed, err := u.HashPassword(validatedUser.Password)
	if err != nil || utf8.RuneCountInString(validatedUser.Password) < 5 {
		return nil, errors.New("Couldn't hash password. Make sure password has length >= 5")
	}
	toBeInsertedUser.SetPassword(string(hashed))

	return toBeInsertedUser.Save(context.Background())
}

func (u usercli) Login(validatedUser *LoginUserValidInput) (string, error) {
	// ALTERNATIVE: Instead of first fetching and comparing,
	// How about, fetching with hashed password?
	userEntity, err := u.db.Query().Unique(true).Where(user.Email(validatedUser.Email)).Select(user.FieldPassword, user.FieldID, user.FieldEmail).First(context.Background())
	if err != nil {
		return "", err
	}

	// Check creds
	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(validatedUser.Password))
	if err != nil {
		return "", err
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"id":    userEntity.ID.String(),
		"email": userEntity.Email,
		"exp":   30 * 24 * 60 * 60 * 1000,
		"iat":   time.Now().UnixMilli(),
	},
	).SignedString([]byte(os.Getenv("SIGNING_SECRET")))
}

func (u usercli) HashPassword(unhashed string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(unhashed), bcrypt.DefaultCost)
}
