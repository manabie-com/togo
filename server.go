package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"togo/common/environment"
	"togo/db"
	"togo/rest"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	environment.Load(".env")

	db, err := db.Connect()

	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
	sqlDB.SetMaxIdleConns(125)
	sqlDB.SetMaxOpenConns(250)

	defer sqlDB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	restService := rest.Handler(db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/user", restService.CreateUser).Methods("POST")
	router.HandleFunc("/api/user", restService.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/user", restService.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/task", restService.CreateTask).Methods("POST")

	log.Printf("Server running on http://localhost:%s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
