package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-envparse"
)

func parseEnv(env_path string) (map[string]string, error) {
	file, err := os.Open(env_path)
	if err != nil {
		log.Print("No environment file found. Not setting any variables.")
		return nil, err
	}

	reader := bufio.NewReader(file)
	varsMap, err := envparse.Parse(reader)
	if err != nil {
		log.Print("Error while parsing env file")
		return nil, err
	}
	return varsMap, nil

}

func InjectEnv(envPath string) {
	varsMap, err := parseEnv(envPath)
	if err == nil {
		for k, v := range varsMap {
			fmt.Println(v)
			os.Setenv(k, v)
		}
	}
}
