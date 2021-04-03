package main

import (
	"database/sql"

	"net/http"

	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/util"
	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger := logs.WithPrefix("main")
	_, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logger.Error("error opening db", zap.Any("Error", err.Error()))
	}

	err = util.LoadConfig("./configs")
	if err != nil {
		logger.Error("error loading config", zap.Any("Error", err.Error()))
	}

	logger.Info("Server is running")
	// serving and return error
	postgres := postgres.NewPostgres()

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: util.Conf.SecretKey,
		Store:  postgres,
	})
}
