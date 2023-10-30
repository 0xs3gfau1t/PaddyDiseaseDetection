package cmd

import (
	"github.com/spf13/cobra"
	"segFault/PaddyDiseaseDetection/pkg/server"
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
