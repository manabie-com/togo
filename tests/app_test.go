package main_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	togo "togo/app"

	"github.com/joho/godotenv"
)

var app = togo.App{}

func TestMain(m *testing.M) {
	if err := godotenv.Load("../app.env"); err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	db_params := togo.DB_Params{
		DB_NAME:     os.Getenv("APP_DB_NAME"),
		DB_USERNAME: os.Getenv("APP_DB_USERNAME"),
		DB_PASSWORD: os.Getenv("APP_DB_PASSWORD"),
	}

	app.Initialize(&db_params)

	ensureTablesExist()
	code := m.Run()
	clearTables()
	os.Exit(code)
}

func ensureTablesExist() {
	if _, err := app.DB.Exec(usersTableCreationQuery); err != nil {
		log.Fatal(err)
	}

	if _, err := app.DB.Exec(todosTableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTables() {
	app.DB.Exec("DELETE * FROM todos")
	app.DB.Exec("ALTER SEQUENCE todos_id_seq RESTART WITH 1")
	app.DB.Exec("DELETE * FROM users")
	app.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

const todosTableCreationQuery = `CREATE TABLE IF NOT EXISTS todos
(
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
	created_date TEXT NOT NULL,
	CONSTRAINT todos_PK PRIMARY KEY (id),
	CONSTRAINT todos_FK FOREIGN KEY (user_id) REFERENCES users(id)
)`

const usersTableCreationQuery = `CREATE TABLE IF NOT EXISTS users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
)`

func makeRequestTo(endpoint, method string) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest(method, endpoint, nil)
	app.Router.ServeHTTP(responseRecorder, request)

	return responseRecorder
}

func TestApp_ShouldHandleNonExistingRoutes(t *testing.T) {
	routes := map[string][]string{
		"POST":   {"/offices", "/are", "/outdated"},
		"GET":    {"/me", "/a", "/pet", "/gopher"},
		"PUT":    {"/some", "/spice", "/in", "/your", "/life"},
		"DELETE": {"/unwated", "/memories", "/from", "/your", "/mind"},
	}

	for method, endpoints := range routes {
		for _, endpoint := range endpoints {
			response := makeRequestTo(endpoint, method)

			if http.StatusNotFound != response.Code {
				t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, response.Code)
			}
		}
	}
}

func TestApp_ShouldHaveARouteToTodo(t *testing.T) {
	response := makeRequestTo("/todo", "POST")

	if http.StatusOK != response.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
	}
}
