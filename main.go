package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/handlers"
)

func main() {
	// DB URL was supposed to be configured as environment variable
	// or console param/flag
	database, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5433/manabie?sslmode=disable")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	serveHandlers := &handlers.Handlers{
		JWTSecret: "test",
		DB:        database,
	}
	serveHandlers.LoadHandlers()
	if err := http.ListenAndServe(":5050", serveHandlers); err != nil {
		log.Println(err)
	}
}
