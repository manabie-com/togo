package main

import (
	"togo/configs"
	"togo/pkg/databases"
	"togo/pkg/logger"
	"togo/server"
)

func main() {
	logger := logger.NewLogger()
	config := configs.ReadConfig()
	db := databases.NewPostgres(config.PostgreSQL)

	server := server.Server{
		Config: config,
		Logger: logger,
		Db:     db,
	}
	server.Start()
}
