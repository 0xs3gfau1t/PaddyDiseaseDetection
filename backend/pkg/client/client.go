package client

import (
	"log"
	"os"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/pkg/config"
	"segFault/PaddyDiseaseDetection/pkg/storage"
)

type Client struct {
	db                 *ent.Client
	User               UserClient
	IdentifiedDiseases IdentifiedDiseasesClient
	MlConsumer         MlConsumerClient
}

var Cli *Client

func init() {
	Cli = New()
	go func() {
		log.Println("[+] Starting response consumer")
		Cli.MlConsumer.Run()
	}()
}

func New() *Client {
	dbClient, err := config.NewDbClient()
	storageAdapter := storage.NewSupaBaseStorage(os.Getenv("SUPABASE_CONN_STRING"), os.Getenv("SUPABASE_KEY"), os.Getenv("IMAGE_BUCKET"))
	if err != nil {
		log.Fatal("Couldn't initialize a database client")
	}
	rbtPublisher := config.Publisher
	if rbtPublisher == nil {
		log.Println("[!] Couldn't initialize rabbit queue")
	}

	return &Client{
		db: dbClient,
		User: usercli{
			db: dbClient.User,
		},
		IdentifiedDiseases: IdentifiedDiseases{
			dbDiseaseIdentified: dbClient.DiseaseIdentified,
			dbImage:             dbClient.Image,
			storage:             storageAdapter,
			rabbitPublisher:     rbtPublisher,
		},
		MlConsumer: MlConsumer{
			dbDisease:           dbClient.Disease,
			dbDiseaseIdentified: dbClient.DiseaseIdentified,
			channel:             config.QChannel,
		},
	}
}
