package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	liteDB := sqlite.NewLiteDB(db)
	todoService := services.NewToDoService("wqGyEBBfPK9w3Lxw", liteDB)

	http.ListenAndServe(":5050", todoService)
}
