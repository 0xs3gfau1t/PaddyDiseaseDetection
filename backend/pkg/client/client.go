package client

import (
	"log"
	"os"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/pkg/config"
	"segFault/PaddyDiseaseDetection/pkg/storage"
)

type Client struct {
	db      *ent.Client
	User    UserClient
	Storage storage.Storage
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
		Storage: &storage.SupaBaseStorage{
			Client: config.NewSupabaseClient(),
			Bucket: os.Getenv("IMAGE_BUCKET"),
		},
	}
}
