package main

import (
	"database/sql"

	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/manabie-com/togo/internal/util"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger := logs.WithPrefix("main")
	_, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logger.Error("error opening db", "process", err.Error())
	}

	err = util.LoadConfig("./configs")
	if err != nil {
		logger.Error("error loading config", "process", err.Error())
	}

	logger.Info("Server is running", "process", nil)
	// serving and return error
	postgres := postgres.NewPostgres()
	server := transport.NewServer(postgres)
	if err := server.Start("0.0.0.0:5050"); err != nil {
		logger.Error("Cannot start server", "process", err)
		return
	}
}
