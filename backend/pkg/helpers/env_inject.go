package helpers

import (
	"bufio"
	"github.com/hashicorp/go-envparse"
	"log"
	"os"
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

func InjectEnv() {
	envPath := os.Getenv("LOCAL_ENV_PATH")
	if envPath == "" {
		log.Println("No env file found. Make sure to set LOCAL_ENV_PATH env variable to apply.")
		return
	}
	varsMap, err := parseEnv(envPath)
	if err == nil {
		for k, v := range varsMap {
			os.Setenv(k, v)
		}
	}
}
