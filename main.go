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
	app := &app.App{}
	defer app.DB.Close()
	app.Init()
	app.Run(":" + port)
}
