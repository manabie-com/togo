package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/tylerb/graceful"

	"github.com/quochungphp/go-test-assignment/src/domain/api"
	"github.com/quochungphp/go-test-assignment/src/infrastructure/pg_driver"
	"github.com/quochungphp/go-test-assignment/src/infrastructure/redis_driver"
	"github.com/quochungphp/go-test-assignment/src/pkgs/settings"
)

func main() {
	// Init redis
	redisHost := os.Getenv(settings.RedisHost) + ":" + os.Getenv(settings.RedisPort)
	redisDriver := redis_driver.RedisDriver{}
	err := redisDriver.Setup(redis_driver.RedisConfiguration{
		Addr: redisHost,
	})
	if err != nil {
		panic(fmt.Sprint("Error while setup redis driver: ", err))
	}

	// Init postgresql
	pgSession, err := pg_driver.Setup(pg_driver.DBConfiguration{
		Driver:   os.Getenv(settings.DbDriver),
		Host:     os.Getenv(settings.PgHost),
		Port:     os.Getenv(settings.PgPort),
		Database: os.Getenv(settings.PgDB),
		User:     os.Getenv(settings.PgUser),
		Password: os.Getenv(settings.PgPass),
	})
	if err != nil {
		panic(fmt.Sprint("Error while setup postgres driver: ", err))
	}
	// Init Gorilla Router
	router := mux.NewRouter()
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
			Handler: router,
		},
	}

	log.Printf("Server is runing port: %v\n", os.Getenv(settings.Port))
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("While start server")
	}
}
