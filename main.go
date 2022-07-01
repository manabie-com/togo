package main

import (
	"log"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port, ok := syscall.Getenv("PORT")
	if !ok {
		port = "8000"
	}
	CONNECT_STR, ok := syscall.Getenv("CONNECT_STR")
	if !ok {
		log.Fatal("Please set CONNECT_STR environment")
	}
	app := &app.App{}
	app.Init(CONNECT_STR)
	app.Run(":" + port)
	defer app.DB.Close()
}
