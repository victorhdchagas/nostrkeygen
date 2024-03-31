package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func DotEnv() (DotEnvConfig, error) {
	config := DotEnvConfig{}

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "10101"
	}
	intPort, err := strconv.Atoi(port)
	if err != nil {
		return config, err
	}
	config.Port = intPort

	return config, nil
}

type DotEnvConfig struct {
	Port int
}
