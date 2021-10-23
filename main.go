package main

import (
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/transports"
	"github.com/manabie-com/togo/internal/services"
	pg "github.com/manabie-com/togo/internal/storages/pg"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	db, err := pgxpool.Connect(context.Background(),"postgresql://admin:123456@localhost:5432/togodb?sslmode=disable")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":5050", &transports.ToDoTrans{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		TodoSvc: services.ToDoService{
		Store: &pg.PgDB{
			DB: db,
		},
	},
	})
}
