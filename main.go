package main

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/services"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
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
			log.Println(err.Error())
			return
		}
		log.Println("|--http server was shut down")

		// Close db
		if err := db.Close(); err != nil {
			log.Println(err.Error())
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

	/*http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	})*/
}
