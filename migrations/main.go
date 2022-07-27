package main

import (
	"togo/configs"
	"togo/migrations/migrate"
	"togo/pkg/databases"
	"togo/pkg/logger"
)

func init() {
	logger.NewLogger()
	configs.ReadConfig()
}

func main() {
	db := databases.NewPostgres()

	err := migrate.Migrate(db)
	if err != nil {
		logger.L.Sugar().Error(err)
	} else {
		logger.L.Info("Success create table")
	}
}
