package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/sqldb"
	"github.com/manabie-com/togo/internal/transport"

	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

type Config struct {
	StorageConfig storages.Config `json:"storages"`
	JwtKey        string          `json:"jwt_key"`
}

var appConfig Config

func main() {
	config.Load("./config.json", &appConfig)
	postgresConfig := appConfig.StorageConfig.Postgres
	postgresConnStr, err := postgresConfig.Build()
	if err != nil {
		log.Fatal("error building config postgres connStr", err)
	}
	dbPool, err := sql.Open("postgres", postgresConnStr)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	if postgresConfig.MaxIdleConns > 0 {
		dbPool.SetMaxIdleConns(postgresConfig.MaxIdleConns)
	}

	storeRepo := &sqldb.SqlDB{
		DB: dbPool,
	}
	todoService := services.NewToDoService(storeRepo)
	httpHandler := transport.NewHttpHandler(appConfig.JwtKey, todoService)
	http.ListenAndServe(":5050", httpHandler)
}
