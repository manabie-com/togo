package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerLoop(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	StartServer(a)

	<-idleConnsClosed
}

func StartServer(a *fiber.App) {
	// Run server.
	port := ":" + os.Getenv("PORT")
	if err := a.Listen(port); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
