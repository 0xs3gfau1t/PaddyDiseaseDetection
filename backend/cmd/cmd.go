package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(
		//
		// Add more commands here
		// Eg: Migrations, Seeding, Testing commands
		//
		serverCmd,
	)
}

var cmd = &cobra.Command{
	Use:   "pd",
	Short: "Paddy Disease Detection Backend Application",
}

func Execute() error {
	return cmd.Execute()
}
