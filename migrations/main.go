package main

import (
	"log"
	"togo/migrations/migrate"
	"togo/pkg/databases"
	"togo/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
	db := databases.NewPostgres()
	logger := logger.NewLogger()

	err = migrate.Migrate(db)
	if err != nil {
		logger.Sugar().Error(err)
	} else {
		logger.Info("Success create table")
	}
}
