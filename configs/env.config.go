package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// EnvGoURI Get then URI of the mongodb database
func EnvGoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv("MONGODB_URI")
}

// EnvGoDB Get then name of the mongodb database
func EnvGoDB() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv("MONGODB_NAME")
}
