package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/huynhhuuloc129/todo/routers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const defaultPort = "8000"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT") // Load PORT from env file (if it had)
	if port == "" {
		port = defaultPort
	}
	DB_URI := os.Getenv("DB_URI")
	
	db := models.Connect(DB_URI) // connect to database
	Repo := controllers.NewBaseHandler(db)

	r := mux.NewRouter().StrictSlash(true)
	routers.Routing(r, Repo)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
