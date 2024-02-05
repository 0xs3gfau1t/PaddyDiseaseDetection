package client

import (
	"context"
	"errors"
	"log"
	"os"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/ent/user"
	"segFault/PaddyDiseaseDetection/pkg/location"
	"segFault/PaddyDiseaseDetection/types"
	"time"
	"unicode/utf8"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserClient interface {
	UserDetails(uuid.UUID) (*types.UserProfileData, error)
	CreateUser(*types.CreateUserValidInput) (*ent.User, error)
	HashPassword(string) ([]byte, error)
	CompareHashedPassword(string, string) error
	Login(*types.LoginUserValidInput) (string, error)
}

type usercli struct {
	db *ent.UserClient
}

func (u usercli) UserDetails(id uuid.UUID) (*types.UserProfileData, error) {
	user, err := u.db.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}

	userLE := location.LocationExtractorFromUser{
		Userid: user.ID,
		Db:     Cli.db,
	}
	latitude, longitude, err := userLE.GetLocation()
	if err != nil {
		log.Printf("Couldn't get user location: %s", err)
	}

	return &types.UserProfileData{
		Name:     user.Name,
		Image:    "", // TODO: add user image to db
		Email:    user.Email,
		Verified: user.Password != "",
		Location: types.UserCoords{
			Latitude:  latitude.ToFloat(),
			Longitude: longitude.ToFloat(),
		},
	}, nil
}

func (u usercli) CreateUser(validatedUser *types.CreateUserValidInput) (*ent.User, error) {
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

func (u usercli) Login(validatedUser *types.LoginUserValidInput) (string, error) {
	userEntity, err := u.db.Query().Unique(true).Where(user.Email(validatedUser.Email)).Select(user.FieldPassword, user.FieldID, user.FieldEmail).First(context.Background())
	if err != nil {
		return "", err
	}

	// Check creds
	err = u.CompareHashedPassword(validatedUser.Password, userEntity.Password)
	if err != nil {
		return "", err
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, &types.JwtType{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 30*24*60*60,
			IssuedAt:  time.Now().Unix(),
		},
		AuthenticatedUserRequestValues: types.AuthenticatedUserRequestValues{
			Id:    userEntity.ID,
			Email: userEntity.Email,
		},
	},
	).SignedString([]byte(os.Getenv("SIGNING_SECRET")))
}

func (u usercli) HashPassword(unhashed string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(unhashed), bcrypt.DefaultCost)
}

func (u usercli) CompareHashedPassword(unhashed string, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(unhashed))
}
