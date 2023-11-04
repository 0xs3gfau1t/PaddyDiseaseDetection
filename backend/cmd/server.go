package cmd

import (
	"segFault/PaddyDiseaseDetection/pkg/server"

	"github.com/spf13/cobra"
)

var serverPort string

func init() {
	serverCmd.Flags().StringVarP(&serverPort, "port", "p", "3000", "server port")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the backend server",
	RunE:  func(cmd *cobra.Command, args []string) error { return server.Run(serverPort) },
}
