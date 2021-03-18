package main

import (
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	//db, err := sql.Open("sqlite3", "./data.db")
	db, err := storages.GetConnection("postgres", "localhost", 5432, "postgres", "postgres", "todo")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &storages.LiteDB{
			DB: db,
		},
	})
}
