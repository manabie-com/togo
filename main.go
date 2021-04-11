package main

import (
	"context"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"

	"github.com/jackc/pgx/v4"
)

func main() {
	const postgresURL = "postgres://test:test@localhost:5432/test"

	db, err := pgx.Connect(context.TODO(), postgresURL)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close(context.TODO())

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.PostgresDB{
			DB: db,
		},
	})
}
