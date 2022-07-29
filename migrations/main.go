package main

import (
	"togo/configs"
	"togo/migrations/migrate"
	"togo/pkg/databases"
	"togo/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	config := configs.ReadConfig()
	db := databases.NewPostgres(config.PostgreSQL)

	err := migrate.Migrate(db)
	if err != nil {
		logger.Sugar().Error(err)
	} else {
		logger.Info("Success create table")
	}
}
