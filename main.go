package main

import (
	"database/sql"
	"log"
	"net/http"

	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	todoTransportHttp "github.com/manabie-com/togo/internal/transport/http"
	todousecase "github.com/manabie-com/togo/internal/usecase"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	todoRepo := &sqllite.LiteDB{
		DB: db,
	}
	JWTKey := "wqGyEBBfPK9w3Lxw"
	userService := todousecase.NewTodoUsecase(todoRepo)
	log.Println("started")

	http.ListenAndServe(":5050", todoTransportHttp.NewTodoHandler(JWTKey, userService))
}
