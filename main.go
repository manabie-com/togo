package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manabie-com/togo/internal/services"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func createDBConnection() (*sql.DB, error) {
	switch os.Getenv("DB_DRIVER") {
	case "postgres":
		dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"))
		return sql.Open("postgres", dns)
	case "sqlite3":
		return sql.Open("sqlite3", "./data.db")
	default:
		return nil, errors.New("cannot find db driver")
	}
}

func main() {
	db, err := createDBConnection()
	if err != nil {
		log.Fatal("error opening db ", err)
	}

	// seed user
	stmt := "INSERT INTO users(id, username, password) VALUES (1, 'nohattee', '1qaz@WSX');"
	db.ExecContext(context.Background(), stmt)

	log.Println("Start server")
	http.ListenAndServe(":5050", services.NewToDoService(db))
}
