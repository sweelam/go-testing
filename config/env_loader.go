package envlaoder

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadAllEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func LoadEnv(envName string) string {
	env := os.Getenv(envName)
	if env == "" {
		log.Fatalf("environment variable %v is not set", envName)
	}
	return env
}
