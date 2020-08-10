package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/services"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"
)

func main() {
	connStr := "postgres://postgres:root@db:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":5051", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.PgDB{
			DB: db,
		},
	})
}
