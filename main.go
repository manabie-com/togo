package main

import (
	"context"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
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

	// Postgres config from env
	config := &postgres.Config{
		Host: util.GetEnv("POSTGRES_HOST", "localhost"),
		Port: util.GetEnv("POSTGRES_PORT", "5432"),
		Usr:  util.GetEnv("POSTGRES_USER", "togo"),
		Pwd:  util.GetEnv("POSTGRES_PASSWORD", "togo"),
		Db:   util.GetEnv("POSTGRES_DB", "togo"),
	}

	// New postgres db instance
	pg, err := postgres.NewPostgres(context.WithValue(context.Background(), "config", config))
	if err != nil {
		log.Println("error opening db", err)
		return
	}

	// New togo service instance
	s := services.NewToDoService("wqGyEBBfPK9w3Lxw", ":5050", pg)

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
		log.Println("|――http server was shut down")

		// Close db
		pg.Close()
		log.Println("|――db was shut down")

		log.Println("web app was shut down ")
	}()

	for {
		select {
		case <-interrupt:
			log.Println("app interrupt")
			return
		case err := <-s.HttpServerErr():
			log.Println("ERR:", err.Error())
			return
		}
	}
}
