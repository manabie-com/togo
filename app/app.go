package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

type DB_Params struct {
	DB_NAME, DB_USERNAME, DB_PASSWORD string
}

func (app *App) Initialize(db *DB_Params) {
	app.Router = mux.NewRouter()
	app.initializeRoutes()
	app.initializeDatabase(db)
}

func (app *App) Run(addr string) error {
	return http.ListenAndServe(addr, app.Router)
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/todo", app.createTodo).Methods("POST")
}

func (app *App) initializeDatabase(db *DB_Params) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		db.DB_USERNAME,
		db.DB_PASSWORD,
		db.DB_NAME,
	)

	var err error
	app.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) createTodo(w http.ResponseWriter, r *http.Request) {
	var payload []int = []int{}
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}