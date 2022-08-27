package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DBConnect() *sql.DB {

	connStr := "postgres://testuser:P@55w0rd@localhost:5432/testdb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if db == nil {
		log.Fatal("Failed to connect to database")
		os.Exit(100)
	}
	log.Printf("Database connection successful")
	return db
}
