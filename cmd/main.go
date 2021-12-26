package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/perfectbuii/togo/configs"
	"github.com/perfectbuii/togo/internal/domain"
	"github.com/perfectbuii/togo/internal/storages"
	"github.com/perfectbuii/togo/internal/storages/inmem"
	"github.com/perfectbuii/togo/internal/storages/postgres"
	"github.com/perfectbuii/togo/internal/transport"
)

var (
	cfg *configs.Config

	dbClient *sql.DB

	// store
	userStore      storages.UserStore
	taskStore      storages.TaskStore
	taskCountStore storages.TaskCountStore

	// domain
	taskDomain domain.TaskDomain
	authDomain domain.AuthDomain

	// handler
	taskHandler transport.TaskHandler
	authHandler transport.AuthHandler
)

func main() {
	if err := loadConfig(); err != nil {
		panic(err)
	}
	if err := loadDatabase(); err != nil {
		panic(err)
	}

	loadStores()
	loadDomain()
	loadHandler()
	if err := loadHttpServer(); err != nil {
		panic(err)
	}
}

func loadConfig() error {
	// err := godotenv.Load(".env")
	// if err != nil {
	// log.Fatalf("Error loading .env file")
	// }

	pstr := os.Getenv("PORT")
	p, err := strconv.Atoi(pstr)
	if err != nil {
		return err
	}
	cfg = &configs.Config{
		DBAddress:    os.Getenv("DB_ADDRESS"),
		RedisAddress: os.Getenv("REDIS_ADDRESS"),
		Port:         p,
		JwtKey:       os.Getenv("JWT_KEY"),
	}
	return nil
}

func loadDatabase() error {
	var err error
	dbClient, err = postgres.NewPostgresClient(cfg.DBAddress)
	if err == nil {
		fmt.Println("connect database successful", cfg.DBAddress)
	}
	return err
}

func loadStores() {
	taskStore = postgres.NewTaskStore(dbClient)
	taskCountStore = inmem.NewTaskCountStore()
	userStore = postgres.NewUserStore(dbClient)
}

func loadDomain() {
	taskDomain = domain.NewTaskDomain(taskCountStore, taskStore, userStore)
	authDomain = domain.NewAuthDomain(userStore, cfg.JwtKey)
}

func loadHandler() {
	taskHandler = transport.NewTaskHandler(taskDomain)
	authHandler = transport.NewAuthHandler(authDomain)
}

func loadHttpServer() error {
	srv := transport.NewHttpServer(cfg.JwtKey, authHandler, taskHandler)
	log.Printf("http server listening port %v\n", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), srv)
}
