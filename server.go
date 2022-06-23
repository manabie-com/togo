package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	port := os.Getenv("PORT") // Load PORT from env file (if it had)
	if port == "" {
		port = defaultPort
	}
	models.Connect() // connect to database

	r := mux.NewRouter().StrictSlash(true)
	routers.Routing(r)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
