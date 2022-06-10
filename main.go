package main

import (
	"github.com/joho/godotenv"
	"log"
	"time"
	"togo/cmd"
)

func run(stop chan bool) {
	time.Sleep(3 * time.Second)
	stop <- true
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cmd.Execute()
}
