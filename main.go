package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/psql"
)

//for the new version forked repository and changes
const version = "1.0.1"

// Initializing ToDoService struct which will be the main pointer of the app

//creating the main function responsible for running also the server
func main() {
	//load environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	//set config
	var cfg services.Config
	intPort, err := strconv.Atoi(os.Getenv("TOGO_PORT"))
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	flag.IntVar(&cfg.Port, "port", intPort, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development|production)")
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	flag.StringVar(&cfg.Db.Dsn, "dsn", url, "Postgres connection string")
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
		Task:   psql.NewModels(db),
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
