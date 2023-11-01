package client

import (
	"context"
	"errors"
	"log"
	"segFault/PaddyDiseaseDetection/ent"

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

type User interface {
	UserDetails(id uuid.UUID) (*ent.User, error)
	CreateUser(validatedUser *CreateUserValidInput) (*ent.User, error)
	HashPassword(unhashed string) ([]byte, error)
}

type user struct {
	db *ent.UserClient
}

func (u user) UserDetails(id uuid.UUID) (*ent.User, error) {
	return u.db.Get(context.Background(), id)
}

func (u user) CreateUser(validatedUser *CreateUserValidInput) (*ent.User, error) {
	toBeInsertedUser := u.db.Create()
	toBeInsertedUser.SetID(uuid.New()) // This line might cause server to crash however unlikely
	toBeInsertedUser.SetName(validatedUser.Name)
	toBeInsertedUser.SetLocation(validatedUser.Location)
	toBeInsertedUser.SetEmail(validatedUser.Email)

	hashed, err := u.HashPassword(validatedUser.Password)
	if err != nil {
		return nil, errors.New("Couldn't hash password")
	}
	toBeInsertedUser.SetPassword(string(hashed))

	log.Printf("Inserting: %v", validatedUser.Name)

	return toBeInsertedUser.Save(context.Background())
}

func (u user) HashPassword(unhashed string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(unhashed), 4)
}
