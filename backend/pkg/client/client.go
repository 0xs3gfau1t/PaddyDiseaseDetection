package client

import (
	"log"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/pkg/config"
)

type Client struct {
	db   *ent.Client
	User UserClient
}

var Cli *Client

func init() {
	Cli = New()
}

func New() *Client {
	dbClient, err := config.NewDbClient()
	if err != nil {
		log.Fatal("Couldn't initialize a database client")
	}
	return &Client{
		db: dbClient,
		User: usercli{
			db: dbClient.User,
		},
	}
}
