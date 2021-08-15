// tasks_test.go
package services

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func TestGetAuthToken(t *testing.T) {
	db, err := sql.Open("sqlite3", "./mock_data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	s := ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/login?user_id=firstUser&password=example", nil)
	s.getAuthToken(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestListTasks(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		db, err := sql.Open("sqlite3", "./mock_data.db")
		if err != nil {
			log.Fatal("error opening db", err)
		}

		s := ToDoService{
			JWTKey: "wqGyEBBfPK9w3Lxw",
			Store: &sqllite.LiteDB{
				DB: db,
			},
		}

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/tasks?created_date=2021-08-14", nil)
		s.listTasks(writer, request)

		if writer.Code != 200 {
			t.Errorf("Response code is %v", writer.Code)
		}
	})
}
