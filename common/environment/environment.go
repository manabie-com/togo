package environment

import (
	"log"

	"github.com/joho/godotenv"
)

// Load for loading the .env file from dynamic path
func Load(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error to load file at %s", path)
	}
}
