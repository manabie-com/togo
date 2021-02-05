// +build integration

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"togo/internal/services"
	"togo/internal/storages/postgres"
	sqllite "togo/internal/storages/sqlite"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setUp() services.ToDoService {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	pgHost := "localhost"
	pgPort := 5432
	pgUser := "postgres"
	pgPassword := "changeme"
	pgDbname := "pg_db"
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPassword, pgDbname)
	dbPg, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalf("connect postgres: %s\n", err)
	}

	return services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
		StorePg: &postgres.ProstgresDB{
			DB: dbPg,
		},
	}
}

func TestWrongUserIDPassword_getAuthToken(t *testing.T) {
	t.Run("TestWrongUserIDPassword", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/login?user_id=firstUser&password=", nil)
		responseRecorder := httptest.NewRecorder()

		app := setUp()
		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusUnauthorized)
		assert.Equal(t, strings.TrimSpace(responseRecorder.Body.String()), `{"error":"incorrect user_id/pwd"}`)
	})
}

func TestInvalid_getAuthToken(t *testing.T) {
	t.Run("TestInvalid_getAuthToken", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/login?user_id=firstUser&password=example", nil)
		responseRecorder := httptest.NewRecorder()

		app := setUp()
		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
	})
}

func TestUnauthorized_listTasks(t *testing.T) {
	t.Run("TestUnauthorized_listTasks", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/tasks?created_date=2020-06-29", nil)
		responseRecorder := httptest.NewRecorder()
		request.Header.Set("Authorization", "")

		app := setUp()
		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusUnauthorized)
		assert.Equal(t, strings.TrimSpace(responseRecorder.Body.String()), "")
	})
}

func TestInvalid_listTasks(t *testing.T) {
	t.Run("TestInvalid_listTasks", func(t *testing.T) {

		requestAuth := httptest.NewRequest(http.MethodGet, "/login?user_id=firstUser&password=example", nil)
		responseRecorderAuth := httptest.NewRecorder()

		app := setUp()
		app.ServeHTTP(responseRecorderAuth, requestAuth)
		assert.Equal(t, responseRecorderAuth.Code, http.StatusOK)

		var auth struct {
			Data string `json:"data"`
		}
		err := json.Unmarshal(responseRecorderAuth.Body.Bytes(), &auth)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodGet, "/tasks?created_date=2020-06-29", nil)
		responseRecorder := httptest.NewRecorder()
		request.Header.Set("Authorization", auth.Data)

		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.Equal(t, strings.TrimSpace(responseRecorder.Body.String()), `{"data":[{"id":"e1da0b9b-7ecc-44f9-82ff-4623cc50446a","content":"first content","user_id":"firstUser","created_date":"2020-06-29"},{"id":"055261ab-8ba8-49e1-a9e8-e9f725ba9104","content":"second content","user_id":"firstUser","created_date":"2020-06-29"},{"id":"2bf3d510-c0fb-41e9-ad12-4b9a60b37e7a","content":"another content","user_id":"firstUser","created_date":"2020-06-29"}]}`)
	})
}

func TestUnauthorized_addTask(t *testing.T) {
	t.Run("TestUnauthorized_addTask", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		responseRecorder := httptest.NewRecorder()
		request.Header.Set("Authorization", "")

		app := setUp()
		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusUnauthorized)
		assert.Equal(t, strings.TrimSpace(responseRecorder.Body.String()), "")
	})
}

func TestInvalid_addTask(t *testing.T) {
	t.Run("TestInvalid_addTask", func(t *testing.T) {

		requestAuth := httptest.NewRequest(http.MethodGet, "/login?user_id=firstUser&password=example", nil)
		responseRecorderAuth := httptest.NewRecorder()

		app := setUp()
		app.ServeHTTP(responseRecorderAuth, requestAuth)
		assert.Equal(t, responseRecorderAuth.Code, http.StatusOK)

		var auth struct {
			Data string `json:"data"`
		}
		err := json.Unmarshal(responseRecorderAuth.Body.Bytes(), &auth)
		assert.NoError(t, err)

		// ----------------------
		request := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		responseRecorder := httptest.NewRecorder()
		request.Header.Set("Authorization", auth.Data)

		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
	})
}

func TestPgWrongUserIDPassword_getAuthToken(t *testing.T) {
	t.Run("TestPgWrongUserIDPassword_getAuthToken", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/login-pg?user_id=firstUser&password=", nil)
		responseRecorder := httptest.NewRecorder()

		app := setUp()
		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusUnauthorized)
		assert.Equal(t, strings.TrimSpace(responseRecorder.Body.String()), `{"error":"incorrect user_id/pwd"}`)
	})
}

func TestPgInvalid_getAuthToken(t *testing.T) {
	t.Run("TestPgInvalid_getAuthToken", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/login-pg?user_id=firstUser&password=example", nil)
		responseRecorder := httptest.NewRecorder()

		app := setUp()
		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
	})
}

func TestPgUnauthorized_listTasks(t *testing.T) {
	t.Run("TestPgUnauthorized_listTasks", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/tasks-pg?created_date=2020-06-29", nil)
		responseRecorder := httptest.NewRecorder()
		request.Header.Set("Authorization", "")

		app := setUp()
		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusUnauthorized)
		assert.Equal(t, strings.TrimSpace(responseRecorder.Body.String()), "")
	})
}

func TestPgInvalid_listTasks(t *testing.T) {
	t.Run("TestPgInvalid_listTasks", func(t *testing.T) {

		requestAuth := httptest.NewRequest(http.MethodGet, "/login-pg?user_id=firstUser&password=example", nil)
		responseRecorderAuth := httptest.NewRecorder()

		app := setUp()
		app.ServeHTTP(responseRecorderAuth, requestAuth)
		assert.Equal(t, responseRecorderAuth.Code, http.StatusOK)

		var auth struct {
			Data string `json:"data"`
		}
		err := json.Unmarshal(responseRecorderAuth.Body.Bytes(), &auth)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodGet, "/tasks?created_date=2020-06-29", nil)
		responseRecorder := httptest.NewRecorder()
		request.Header.Set("Authorization", auth.Data)

		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.Equal(t, strings.TrimSpace(responseRecorder.Body.String()), `{"data":[{"id":"e1da0b9b-7ecc-44f9-82ff-4623cc50446a","content":"first content","user_id":"firstUser","created_date":"2020-06-29"},{"id":"055261ab-8ba8-49e1-a9e8-e9f725ba9104","content":"second content","user_id":"firstUser","created_date":"2020-06-29"},{"id":"2bf3d510-c0fb-41e9-ad12-4b9a60b37e7a","content":"another content","user_id":"firstUser","created_date":"2020-06-29"}]}`)
	})
}

func TestPgUnauthorized_addTask(t *testing.T) {
	t.Run("TestPgUnauthorized_addTask", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/tasks-pg", nil)
		responseRecorder := httptest.NewRecorder()
		request.Header.Set("Authorization", "")

		app := setUp()
		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusUnauthorized)
		assert.Equal(t, strings.TrimSpace(responseRecorder.Body.String()), "")
	})
}

func TestPgInvalid_addTask(t *testing.T) {
	t.Run("TestPgInvalid_addTask", func(t *testing.T) {

		requestAuth := httptest.NewRequest(http.MethodGet, "/login-pg?user_id=firstUser&password=example", nil)
		responseRecorderAuth := httptest.NewRecorder()

		app := setUp()
		app.ServeHTTP(responseRecorderAuth, requestAuth)
		assert.Equal(t, responseRecorderAuth.Code, http.StatusOK)

		var auth struct {
			Data string `json:"data"`
		}
		err := json.Unmarshal(responseRecorderAuth.Body.Bytes(), &auth)
		assert.NoError(t, err)

		// ----------------------
		request := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		responseRecorder := httptest.NewRecorder()
		request.Header.Set("Authorization", auth.Data)

		app.ServeHTTP(responseRecorder, request)
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
	})
}
