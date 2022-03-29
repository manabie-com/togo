package main

import (
	"log"
	"os"

	"github.com/laghodessa/togo/config"
	"github.com/laghodessa/togo/infra/postgres"
)

func main() {
	config.Load()

	dbURL := os.Getenv("POSTGRES_URL")
	if err := postgres.Migrate(dbURL); err != nil {
		log.Fatalf("migrate db: %s", err)
	}
}
