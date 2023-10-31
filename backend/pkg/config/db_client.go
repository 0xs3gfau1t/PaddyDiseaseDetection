package config

import (
	"fmt"
	"os"
	"segFault/PaddyDiseaseDetection/ent"

	_ "github.com/lib/pq"
)

func NewDbClient() (*ent.Client, error) {
	formattedUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	return ent.Open("postgres", formattedUrl)
}
