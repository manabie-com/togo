package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/manabie-com/togo/internal/storages/redis"

	"github.com/manabie-com/togo/internal/storages/postgres"

	"github.com/manabie-com/togo/internal/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Load .env file failed. Use default config")
	}

	storeManager := postgres.GetStorageManager(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
	)

	maxRequestPeHour, err := strconv.Atoi(os.Getenv("MAX_REQ_PER_HOUR"))
	if err != nil || maxRequestPeHour <= 0 {
		log.Fatalln("Invalid rate limit config")
	}
	rateLimiter := redis.GetRateLimiter(
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		maxRequestPeHour,
	)

	server := http.Server{
		Addr:    os.Getenv("SERVER_ADDR"),
		Handler: services.GetTodoService(os.Getenv("JWT_KEY"), storeManager, rateLimiter),
	}

	// Grateful shutdown when terminate server
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-termChan
		ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer ctxCancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Failed to shutdown server: %v", err)
		}
	}()

	fmt.Printf("Starting server on %s ...\n", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("Start server failed:", err)
	}
}
