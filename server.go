package main

import (
	"log"
	"net/http"
	"os"

	"github.com/huynhhuuloc129/todo/models"
	"github.com/huynhhuuloc129/todo/routers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const defaultPort = "8000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	models.Connect() // connect to database

	http.HandleFunc("/", routers.Handle)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
