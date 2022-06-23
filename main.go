package main

import (
	"fmt"
	"lntvan166/togo/db"
	"lntvan166/togo/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()

	route := mux.NewRouter()
	routes.HandleRequest(route)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started!")

}
