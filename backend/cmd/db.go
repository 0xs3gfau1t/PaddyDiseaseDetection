package cmd

import (
	"context"
	"log"
	"segFault/PaddyDiseaseDetection/ent/migrate"
	"segFault/PaddyDiseaseDetection/pkg/config"
	"segFault/PaddyDiseaseDetection/pkg/helpers"

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
