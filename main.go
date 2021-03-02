package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	//db, err := sql.Open("sqlite3", "./data.db")
	db, err := sql.Open("pgx", "postgresql://postgres:example@localhost/postgres")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.PostgreDB{
			DB: db,
		},
	})
}
