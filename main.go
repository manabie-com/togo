package main

import (
	"context"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/internal/util"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	db, err := sqllite.NewSqliteDb()
	if err != nil {
		log.Println("error opening db", err)
		return
	}

	config := &postgres.Config{
		Host: util.GetEnv("POSTGRES_HOST", "localhost"),
		Port: util.GetEnv("POSTGRES_PORT", "5432"),
		Usr:  util.GetEnv("POSTGRES_USER", "togo"),
		Pwd:  util.GetEnv("POSTGRES_PASSWORD", "togo"),
		Db:   util.GetEnv("POSTGRES_DB", "togo"),
	}

	_, err = postgres.NewPostgres(context.WithValue(context.Background(), "config", config))
	if err != nil {
		log.Println("error opening db", err)
		return
	}

	s := services.NewToDoService(db)

	// Release resources
	defer func() {
		log.Println("shutting down web app")
		// Close http server
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		err := s.Shutdown(ctx)
		cancel()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("|--http server was shut down")

		// Close db
		if err := db.Close(); err != nil {
			log.Println(err)
		}
		log.Println("|--db was shut down")

		log.Println("web app was shut down ")
	}()

	for {
		select {
		case <-interrupt:
			return
		case <-s.HttpServerErr():
			return
		}
	}
}


