package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvGoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv("MONGODB_URI")
}

func EnvGoDB() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv("MONGODB_NAME")
}
