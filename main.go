package main

import (
	"database/sql"
	"github.com/manabie-com/togo/internal/api"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	todoApi := api.NewToDoApi("wqGyEBBfPK9w3Lxw", db)
	err = http.ListenAndServe(":5050", &todoApi)
	if err != nil {
		log.Fatal("error listen and serve api", err)
	}
}
