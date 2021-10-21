package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
)

//for the new version forked repository and changes
const version = "1.0.1"

// Initializing ToDoService struct which will be the main pointer of the app

//creating the main function responsible for running also the server
func main() {
	var cfg services.Config
	flag.IntVar(&cfg.Port, "port", 5050, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&cfg.Db.Dsn, "dsn", "postgres://postgres:root@localhost/todos?sslmode=disable", "Postgres connection string")
	flag.StringVar(&cfg.Jwt.Secret, "jwt-secret", "", "secret")
	flag.Parse()
	//date and time for the log
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	//reference to the application type we define and populate the field
	//this app will be the receiver for other parts of the application
	app := &services.ToDoService{
		Config: cfg,
		Logger: logger,
		Models: storages.NewModels(db),
		JWTKey: cfg.Jwt.Secret,
	}
	//running the server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Println("Starting server on port", cfg.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("test")
	}
}

/**
* openDB function for postgres connection
* @param cfg config
* @return DB connection or error
**/
func openDB(cfg services.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil

}
