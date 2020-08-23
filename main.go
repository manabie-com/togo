package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/phuwn/togo/internal/services"
	"github.com/phuwn/togo/internal/storages/database"

	_ "github.com/lib/pq"
)

func init() {
	env := os.Getenv("RUN_MODE")
	if env == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("PG_DATASOURCE"))
	if err != nil {
		log.Fatal("error opening db", err)
	}

	port := ":5050"
	log.Println("transport HTTP, addr " + port)
	http.ListenAndServe(port, &services.ToDoService{
		JWTKey: os.Getenv("JWT_KEY"),
		Store:  &database.Storage{db},
	})
}
