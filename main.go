package main

import (
	"log"
	"net/http"
	"os"

	"context"

	"github.com/joho/godotenv"
	"github.com/manabie-com/togo/internal/services"
	pg "github.com/manabie-com/togo/internal/storages/pg"
	"github.com/manabie-com/togo/internal/transports"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	db, err := pgxpool.Connect(context.Background(), os.Getenv("PG_CONN_URI"))
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":5050", &transports.ToDoTrans{
		JWTKey: os.Getenv("JWK_KEY"),
		TodoSvc: services.ToDoService{
			Store: &pg.PgDB{
				DB: db,
			},
		},
	})
}
