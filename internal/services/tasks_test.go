package services

import (
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	postgresdb "github.com/cuongtop4598/togo-interview/togo/internal/storages/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var todoservice ToDoService

func NewToDoService() *ToDoService {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	dbmanager := postgresdb.DBmanager{gormDB}
	return &ToDoService{
		JWTKey: "",
		Store:  &dbmanager,
	}
}

func TestGetAuthToken(t *testing.T) {
	todoservice := NewToDoService()
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "cuongnm")
	q.Add("password", "123456")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoservice.getAuthToken)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v ", status, http.StatusUnauthorized)
	}
	// check the response body
	rule, _ := regexp.Compile(`"([^\"]+)"`)
	results := rule.FindAllString(rr.Body.String(), -1)
	if results[0] == "data" {
		t.Errorf("wrong json format output")
	}
	if results[1] == "" {
		t.Errorf("token has not been created")
	}
}

func TestListTasks(t *testing.T) {

}

func TestAddTask(t *testing.T) {

}
