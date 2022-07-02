package environment

import (
	"log"

	"github.com/joho/godotenv"
)

func Load(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error to load file at %s", path)
	}
}
