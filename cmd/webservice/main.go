package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lawtrann/togo/http"
	"github.com/lawtrann/togo/postgres"
	"github.com/lawtrann/togo/services"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(fmt.Errorf("error - server failed to start. err: %v", err))
	}
}

func run() error {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
			os.Exit(1)
		}
	}

	db, err := postgres.NewTodoDB()
	if err != nil {
		panic(err)
	}
	svc := services.NewTodoService(db)
	h := http.NewHandler(svc)
	http.RegisterService(h)
	fmt.Println("Starting webservice... on localhost:3000")
	return http.ListenAndServe(":3000", nil)
}
