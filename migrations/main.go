package main

import (
	"togo/configs"
	"togo/migrations/migrate"
	"togo/pkg/databases"
	"togo/pkg/logger"
)

func init() {
	configs.ReadConfig()
}

func main() {
	db := databases.NewPostgres()
	logger := logger.NewLogger()

	err := migrate.Migrate(db)
	if err != nil {
		logger.Sugar().Error(err)
	} else {
		logger.Info("Success create table")
	}
}
