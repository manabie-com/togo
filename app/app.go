package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Init() {
	if err := godotenv.Load("C:\\Users\\quang\\Desktop\\Project\\Go\\internship\\togo\\.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	connectURL := os.Getenv("CONNECT_STR")
	db, err := sql.Open("postgres", connectURL)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database is connected!")
	a.DB = db
	a.Router = mux.NewRouter()
	a.Routes()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
