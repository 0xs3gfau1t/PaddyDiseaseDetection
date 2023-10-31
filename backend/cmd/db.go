package cmd

import (
	"context"
	"segFault/PaddyDiseaseDetection/ent/migrate"
	"segFault/PaddyDiseaseDetection/pkg/config"
	"segFault/PaddyDiseaseDetection/pkg/helpers"

	"github.com/spf13/cobra"
)

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
