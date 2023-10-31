package client

import (
	"log"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/pkg/config"
)

type Client struct {
	db   *ent.Client
	User User
}

func New() *Client {
	dbClient, err := config.NewDbClient()
	if err != nil {
		log.Fatal("Couldn't initialize a database client")
	}
	return &Client{
		db: dbClient,
		User: user{
			db: dbClient.User,
		},
	}
}
