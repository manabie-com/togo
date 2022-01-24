package main

import (
	"github.com/joho/godotenv"
	"log"
	"togo/cmd"
)

func main()  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cmd.Execute()
}