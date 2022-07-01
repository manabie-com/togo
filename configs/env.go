package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load("../.env")
	if err != nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return os.Getenv("MONGOURI")
}

