package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"

	store "github.com/manabie-com/togo/internal/storages"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"
)

var (
	db      = initDB()
	todoSRV = ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.PgDB{
			DB: db,
		},
	}
	token = ""
)

func initDB() *sql.DB {
	connStr := "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func TestLogin(t *testing.T) {
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	params := req.URL.Query()
	params.Add("user_id", "firstUser")
	params.Add("password", "example")
	req.URL.RawQuery = params.Encode()

	rr := httptest.NewRecorder()
	todoSRV.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	type Res struct {
		Data store.User
	}
	res := Res{}
	json.NewDecoder(rr.Body).Decode(&res)
	if len(res.Data.Token) == 0 {
		t.Error("Token is null")
	}
	token = res.Data.Token
}

func TestCreateTask(t *testing.T) {
	reqStr := []byte(`{"content": "test"}`)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(reqStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr := httptest.NewRecorder()
	todoSRV.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetListTask(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	params := req.URL.Query()
	params.Add("created_date", "2020-06-29")
	req.URL.RawQuery = params.Encode()

	rr := httptest.NewRecorder()
	todoSRV.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	type Res struct {
		Data []store.Task
	}
	res := Res{}
	json.NewDecoder(rr.Body).Decode(&res)
	if len(res.Data) == 0 {
		t.Error("Data is null")
	}
}
