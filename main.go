package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/utils"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func StartServer(db *sql.DB) {
	appPort := os.Getenv("APP_PORT")
	jwtKey := os.Getenv("APP_JWTKEY")

	http.ListenAndServe(":"+appPort, &services.TransportService{
		JWTKey: jwtKey,
		DB:     db,
	})
}

func main() {
	/*db, err := sql.Open("sqlite3", "./db/sqlite/data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}*/

	if err := utils.InitEnv(); err != nil {
		log.Fatal("error loading env", err)
	}

	db, err := utils.InitDB()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	StartServer(db)
}
