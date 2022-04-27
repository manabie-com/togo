package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"togo/internal/middleware"
	"togo/internal/route"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}

	port := os.Getenv("PORT")

	app := fiber.New(fiber.Config{
		IdleTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	})

	middleware.Logger(app)

	route.Setup(app)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Fiber was successful shutdown.")

}
