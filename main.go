package main

import (
	"database/sql"

	"net/http"

	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
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

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: util.Conf.SecretKey,
		Store:  postgres,
	})
}
