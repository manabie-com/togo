package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/ratelimiters"
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	// Init rate limiter
	ratelimiters.InitLocalRatelimiter(&sqllite.LiteDB{db})

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	})
}
