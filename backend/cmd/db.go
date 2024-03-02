package cmd

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"segFault/PaddyDiseaseDetection/ent/migrate"
	"segFault/PaddyDiseaseDetection/pkg/config"
	"segFault/PaddyDiseaseDetection/pkg/helpers"
	"segFault/PaddyDiseaseDetection/types"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func init() {
	migrateCmd.AddCommand(
		migrateReset,
	)
}

var migrateCmd = &cobra.Command{
	Use:    "migrate",
	Short:  "Apply database migrations veri hardlyy",
	PreRun: func(cmd *cobra.Command, args []string) { helpers.InjectEnv() },
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := config.NewDbClient()
		defer client.Close()

		if err != nil {
			return err
		}
		return client.Debug().Schema.Create(
			context.Background(),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
			migrate.WithForeignKeys(true))
	},
}

var migrateReset = &cobra.Command{
	Use:    "reset",
	Short:  "Migrates after dropping all values. Needed if schema changes with NULL constraint conflicts",
	PreRun: func(cmd *cobra.Command, args []string) { helpers.InjectEnv() },
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := config.NewDbClient()
		defer client.Close()

		if err != nil {
			log.Printf("Cannot create db client")
			return err
		}

		ctx := context.Background()
		_, err = client.User.Delete().Exec(ctx)
		// Upon creation of additional tables
		// Add delete queries here

		return client.Debug().Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
			migrate.WithForeignKeys(true))
	},
}

var seedCmd = &cobra.Command{
	Use:    "seed",
	Short:  "Seed disease and solutions to the database.",
	PreRun: func(cmd *cobra.Command, args []string) { helpers.InjectEnv() },
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := config.NewDbClient()
		defer client.Close()

		if err != nil {
			log.Printf("Cannot create db client")
			return err
		}

		jsonFile, err := os.Open(os.Getenv("SEED_FILE"))
		if err != nil {
			log.Printf("Cannot open json file")
			return err
		}
		defer jsonFile.Close()

		byteValue, err := io.ReadAll(jsonFile)
		if err != nil {
			log.Printf("Cannot read json file")
			return err
		}
		var seedData types.SeedJson
		json.Unmarshal(byteValue, &seedData)

		ctx := context.Background()

		for i := 0; i < len(seedData.Disease); i++ {
			d, err := client.Disease.Create().SetName(seedData.Disease[i].Name).SetID(uuid.New()).Save(ctx)
			if err != nil {
				log.Printf("Failed to create disease %s", seedData.Disease[i].Name)
				log.Println(err)
				continue
			}
			log.Printf("Created disease %s", seedData.Disease[i].Name)
			for j := 0; j < len(seedData.Disease[i].Solutions); j++ {
				qb := client.Solution.Create()
				qb.SetID(uuid.New())
				qb.SetName(seedData.Disease[i].Solutions[j].Name)
				qb.SetDescription(seedData.Disease[i].Solutions[j].Description)
				qb.SetIngredients(seedData.Disease[i].Solutions[j].Ingredients)
				qb.SetPhotos(seedData.Disease[i].Solutions[j].Photos)
				qb.AddDisease(d)
				err = qb.Exec(ctx)
				if err != nil {
					log.Printf("Failed to create solution %s", seedData.Disease[i].Solutions[j].Name)
					log.Println(err)
					continue
				}
				log.Printf("Created solution %s", seedData.Disease[i].Solutions[j].Name)
			}
			for k := 0; k < len(seedData.Disease[i].Causes); k++ {
				qb := client.Cause.Create()
				qb.SetName(seedData.Disease[i].Causes[k].Name)
				qb.SetImage(seedData.Disease[i].Causes[k].Image)
				qb.AddDisease(d)
				err := qb.Exec(ctx)
				if err != nil {
					log.Printf("Failed to create cause %s", seedData.Disease[i].Causes[k].Name)
					log.Println(err)
					continue
				}
				log.Printf("Created cause %s", seedData.Disease[i].Causes[k].Name)
			}
		}

		return nil
	},
}
