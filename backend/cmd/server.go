package cmd

import (
	"fmt"
	"os"
	"segFault/PaddyDiseaseDetection/pkg/server"

	"github.com/spf13/cobra"
)

var serverPort string

func init() {
	sP := os.Getenv("HOST_PORT")
	fmt.Println("Got port: ", sP)
	if sP == "" {
		sP = "3000"
	}
	serverCmd.Flags().StringVarP(&serverPort, "port", "p", sP, "server port")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the backend server",
	RunE:  func(cmd *cobra.Command, args []string) error { return server.Run(serverPort) },
}
