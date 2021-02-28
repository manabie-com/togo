package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/db"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/handlers"
)

func main() {
	err := config.Setup()
	if err != nil {
		log.Fatalf(err.Error())
	}

	database, _ := db.SetupPostgres(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Values.PostgresUser,
		config.Values.PostgresPassword,
		config.Values.PostgresHost,
		config.Values.PostgresPort,
		config.Values.PostgresDb,
	))
	if err != nil {
		log.Fatalf("error when connecting to postgres: %s\n", err)
	}

	redis, err := db.SetupRedis(fmt.Sprintf("%s:%d", config.Values.RedisHost, config.Values.RedisPort),
		config.Values.RedisPassword, config.Values.RedisDb)
	if err != nil {
		log.Fatalf("error when connecting to redis: %s\n", err)
	}

	addr := fmt.Sprintf("%s:%d", config.Values.HttpInterface, config.Values.HttpPort)
	log.Fatal(http.ListenAndServe(addr, &handlers.HttpHandler{
		UserService: &services.UserService{
			JWTKey: config.Values.JWTKey,
			Storage: &postgres.PostgresDB{
				DB: database,
			},
		},
		TaskService: &services.TaskService{
			Redis: redis,
			Storage: &postgres.PostgresDB{
				DB: database,
			},
		},
	}))
}
