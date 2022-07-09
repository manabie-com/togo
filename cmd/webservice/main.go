package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lawtrann/togo/http"
	"github.com/lawtrann/togo/postgres"
	"github.com/lawtrann/togo/service"
)

func main() {
	// Get enviroment variables
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
			os.Exit(1)
		}
	}

	// Start Main
	m := NewMain()

	// Open database
	m.DB.DSN = GetDns()
	if err := m.DB.Open(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Instantiate repository.
	// UserRepository
	userRepo := postgres.NewUserRepo(m.DB)
	// UserRepository
	todoRepo := postgres.NewTodoRepo(m.DB)

	// Instantiate business services.
	// UserService
	userService := service.NewUserService(userRepo)
	// TodoService
	todoService := service.NewTodoService(todoRepo)
	todoService.UserService = userService

	// Attach underlying services to the HTTP server.
	m.HTTPServer.TodoService = todoService

	// Start the HTTP server.
	fmt.Println("Running on ... localhost:3000")
	m.HTTPServer.ListenAndServe(":3000")
}

// Main represents the program.
type Main struct {
	// Postgres database used by Postgres service implementations.
	DB *postgres.DB

	// HTTP server for handling HTTP communication.
	// Postgres services are attached to it before running.
	HTTPServer *http.Server
}

// NewMain returns a new instance of Main.
func NewMain() *Main {
	return &Main{
		DB:         postgres.NewDB(""),
		HTTPServer: http.NewServer(),
	}
}

func GetDns() string {
	host := os.Getenv("POSTGRESQL_HOST")
	port, err := strconv.Atoi(os.Getenv("POSTGRESQL_PORT"))
	if err != nil {
		log.Fatalf("Postgres port %s is not valid", os.Getenv("POSTGRESQL_PORT"))
		os.Exit(1)
	}
	user := os.Getenv("POSTGRESQL_USERNAME")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	dbname := os.Getenv("POSTGRESQL_DATABASE")

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
