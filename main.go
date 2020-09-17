package main

import (
	"database/sql"
	"fmt"
	"github.com/manabie-com/togo/constants"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db *sql.DB
	var err error
	switch constants.DB_TYPE {
	case constants.POSTGRES:
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable", constants.HOST, constants.PORT, constants.USER, constants.PASSWORD, constants.DB_NAME);
		db, err = sql.Open("postgres", psqlInfo)
		break
	case constants.SQLITE:
		db, err = sql.Open("sqlite3", "./data.db")
		break
	}
	if err != nil {
		log.Fatal("Error opening DB with type %s", err, constants.DB_TYPE)
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
