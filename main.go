package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/rest"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgresql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	host     = "localhost"
	port     = 8899
	user     = "phuonghau"
	password = "phuonghau"
	dbname   = "togo"
)

func newPSQLDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	return db, err
}

func main() {
	// Initializing postgres connection
	db, err := newPSQLDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected")
	}

	// Initialzing dependencies
	mux := http.NewServeMux()
	userRepo := postgresql.NewPSQLUserRepository(db)
	taskRepo := postgresql.NewPSQLTaskRespsitory(db)

	taskService := services.NewTaskService(services.TaskServiceConfiguration{
		TaskRepo: taskRepo,
	})
	authService := services.NewAuthService(services.AuthServiceConfiguration{
		UserRepo: userRepo,
		JWTKey:   services.DefaultJWTKey,
	})

	rest.NewTaskHandler(authService, taskService).Register(mux)
	rest.NewAuthHandler(authService).Register(mux)

	// HTTP Server configuration
	address := "0.0.0.0:8080"
	server := &http.Server{
		Handler: mux,
		Addr:    address,
	}

	// Binding http server to network interface
	log.Println("Server is listening on", address)
	log.Fatal(server.ListenAndServe())
}
