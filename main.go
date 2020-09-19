package main

import (
	"database/sql"
	"github.com/manabie-com/togo/constants"
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/database"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db *sql.DB
	var err error
	switch constants.DB_TYPE {
	case constants.POSTGRES:
		db, err = sql.Open("postgres", constants.GetPostgreConnectionString())
		break
	case constants.SQLITE:
		db, err = sql.Open("sqlite3", "./data.db")
		break
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Connection string invalid ", err)
	}
	err = http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: constants.JWT_KEY,
		Store: &sqllite.Vendor{
			DB: db,
		},
	})
	if err != nil {
		log.Fatal("Server start failed", err)
	}
}
