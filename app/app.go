package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"togo/helpers"

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
	app.seedData()
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

func (app *App) seedData() {
	helpers.ClearTables(app.DB)
	helpers.EnsureTablesExist(app.DB)
	helpers.CreateInitialUser(app.DB)
}

func (app *App) createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&todo); err != nil {
		sendJSONResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := todo.createTodo(app.DB); err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, err.Error())
	}

	sendJSONResponse(w, http.StatusCreated, todo)
}

func sendJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
