package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/configor"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/configs"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/transports"
	"github.com/rs/cors"
)

func main() {
	// Load configs
	configs := configs.Config{}
	err := configor.Load(&configs, "config.yml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", configs.DB.Host, configs.DB.Port, configs.DB.User, configs.DB.Password, configs.DB.Name))
	if err != nil {
		log.Fatal("Occurred error while opening the database with the configuration information", err)
	}

	// Handle routing
	router := mux.NewRouter()
	router.Handle("/login", http.HandlerFunc(transports.NewToDoLoginController(&repositories.DB{DB: db}, configs.JWT.Key).GetAuthToken))
	router.Handle("/tasks", transports.Middleware(http.HandlerFunc(transports.NewToDoTaskController(&repositories.DB{DB: db}, configs.JWT.Key).ListTasks), configs.JWT.Key)).Methods("GET")
	router.Handle("/tasks", transports.Middleware(http.HandlerFunc(transports.NewToDoTaskController(&repositories.DB{DB: db}, configs.JWT.Key).AddTask), configs.JWT.Key)).Methods("POST")

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	log.Fatal(http.ListenAndServe(":5050", corsWrapper.Handler(router)))
}
