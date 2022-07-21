package main

import (
	"fmt"
	"log"
	"os"
	togo "togo/app"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("app.env"); err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	togoApp := togo.App{}
	db_params := togo.DB_Params{
		DB_NAME:     os.Getenv("APP_DB_NAME"),
		DB_USERNAME: os.Getenv("APP_DB_USERNAME"),
		DB_PASSWORD: os.Getenv("APP_DB_PASSWORD"),
	}
	togoApp.Initialize((&db_params))

	err := togoApp.Run(":5000")
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	}
}
