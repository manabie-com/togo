package main

import (
	"database/sql"
	"log"

	"github.com/jericogantuangco/togo/internal/services"
	"github.com/jericogantuangco/togo/internal/storages/postgres"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/todo?sslmode=disable"
	serverAddress = "0.0.0.0:5050"
)

func main() {
	connection, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can't connect to the database:", err)
	}

	store := postgres.NewStore(connection)
	server, err := services.NewServer(store)
	if err != nil {
		log.Fatal("Cannot create server")
	}
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Can't Start server:", err)
	}

}
