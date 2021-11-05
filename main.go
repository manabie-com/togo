package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tylerb/graceful"

	"github.com/quochungphp/go-test-assignment/src/domain/api"
	"github.com/quochungphp/go-test-assignment/src/pkgs/db"
	"github.com/quochungphp/go-test-assignment/src/pkgs/redis"
	"github.com/quochungphp/go-test-assignment/src/pkgs/settings"
)

func main() {
	// Init redis
	redis.Init()

	// Init postgresql
	pgSession := db.Init()

	// Init Gorilla Router
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	})
	api.APIs{router, pgSession}.Init()

	srv := &graceful.Server{
		Timeout: 5 * time.Second,
		BeforeShutdown: func() bool {
			pgSession.Close()
			log.Println("Server is shutting down database connection")

			return true
		},
		Server: &http.Server{
			Addr:    ":" + os.Getenv(settings.Port),
			Handler: c.Handler(router),
		},
	}

	log.Printf("Server is runing port: %v\n", os.Getenv(settings.Port))
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("While start server")
	}
}
