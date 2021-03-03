package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/manabie-com/togo/internal/services"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	//db, err := sql.Open("sqlite3", "./data.db")
	db, err := sql.Open("pgx", "postgresql://postgres:example@localhost/postgres")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()

	services.StartServer(db)
}
