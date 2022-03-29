package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/laghodessa/togo/config"
	"github.com/laghodessa/togo/infra/rest"
)

func main() {
	config.Load()

	db, err := sql.Open("pgx", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("connect db: %s", err)
	}

	server := rest.NewFiber(db)
	server.Listen(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
