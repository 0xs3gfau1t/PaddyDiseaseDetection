package client

import (
	"context"
	"log"
	"segFault/PaddyDiseaseDetection/ent"

	"github.com/google/uuid"
)

type CreateUserValidInput struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Location string `form:"location"`
	Coord    string `form:"coord"`
}

type User interface {
	UserDetails(id uuid.UUID) (*ent.User, error)
	CreateUser(validatedUser *CreateUserValidInput) (*ent.User, error)
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

	log.Printf("Inserting: %v", validatedUser.Name)

	return toBeInsertedUser.Save(context.Background())
}
