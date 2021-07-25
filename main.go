package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/surw/togo/internal/config"
	"github.com/surw/togo/internal/services"
	sqllite "github.com/surw/togo/internal/storages/sqlite"
	"log"
)

func main() {
	db, err := sql.Open("postgres", config.PsqlInfo)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	service := services.NewToDoService(&sqllite.LiteDB{
		DB: db,
	})
	router := services.NewRouter()
	service.Register(router)

	router.Start(5050)
}
