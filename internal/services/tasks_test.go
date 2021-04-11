package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manabie-com/togo/internal/storages"
)

const testJWTKey = "wqGyEBBfPK9w3Lxw"

// mockDB implements DB interface for testing purpose
type mockDB struct {
	mockRetrieveTasks func(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error)
	mockAddTask       func(ctx context.Context, t *storages.Task) error
	mockValidateUser  func(ctx context.Context, userID, pwd sql.NullString) bool
}

// RetrieveTasks delegates the call to mockRetrieveTasks
func (db *mockDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	return db.mockRetrieveTasks(ctx, userID, createdDate)
}

// AddTask delegates the call to mockAddTask
func (db *mockDB) AddTask(ctx context.Context, t *storages.Task) error {
	return db.mockAddTask(ctx, t)
}

// ValidateUser delegates the call to mockValidateUser
func (db *mockDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	return db.mockValidateUser(ctx, userID, pwd)
}

// TestLoginOK tests /login with a valid user
func TestLoginOK(t *testing.T) {
	db := &mockDB{
		mockValidateUser: func(_ context.Context, _, _ sql.NullString) bool {
			return true
		},
	}

	svc := &ToDoService{
		JWTKey: testJWTKey,
		Store:  db,
	}

	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	svc.ServeHTTP(w, r)

	if resp := w.Result(); resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code (want %d  have %d)", http.StatusOK, resp.StatusCode)
	}
}

// TestLoginUnauthorized tests /login with an invalid user
func TestLoginUnauthorized(t *testing.T) {
	db := &mockDB{
		mockValidateUser: func(_ context.Context, _, _ sql.NullString) bool {
			return false
		},
	}

	svc := &ToDoService{
		JWTKey: testJWTKey,
		Store:  db,
	}

	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	svc.ServeHTTP(w, r)

	if resp := w.Result(); resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("unexpected status code (want %d  have %d)", http.StatusUnauthorized, resp.StatusCode)
	}
}

// TestListTasksInvalidToken tests /tasks with an invalid token
func TestListTasksInvalidToken(t *testing.T) {
	db := &mockDB{}

	svc := &ToDoService{
		JWTKey: testJWTKey,
		Store:  db,
	}

	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	svc.ServeHTTP(w, r)

	if resp := w.Result(); resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("unexpected status code (want %d  have %d)", http.StatusUnauthorized, resp.StatusCode)
	}
}

// TestListTasksOK tests /tasks with a valid token
func TestListTasksOK(t *testing.T) {
	var (
		user = "alpha"
		date = "2006-01-02"
	)

	db := &mockDB{
		mockRetrieveTasks: func(_ context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
			if userID.String != user {
				t.Errorf("unexpedted user ID (want %s have %s)", user, userID.String)
			}

			if createdDate.String != date {
				t.Errorf("unexpedted date (want %s have %s)", date, createdDate.String)
			}

			return nil, nil
		},
		mockValidateUser: func(_ context.Context, _, _ sql.NullString) bool {
			return true
		},
	}

	svc := &ToDoService{
		JWTKey: testJWTKey,
		Store:  db,
	}

	// Login to get token
	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := r.URL.Query()
	q.Add("user_id", user)
	r.URL.RawQuery = q.Encode()

	svc.ServeHTTP(w, r)

	var body map[string]string

	if err := json.NewDecoder(w.Result().Body).Decode(&body); err != nil {
		t.Fatal(err)
	}

	// List tasks with token
	w = httptest.NewRecorder()

	r, err = http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	q = r.URL.Query()
	q.Add("created_date", date)
	r.URL.RawQuery = q.Encode()
	r.Header.Add("Authorization", body["data"])

	svc.ServeHTTP(w, r)

	if resp := w.Result(); resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code (want %d  have %d)", http.StatusOK, resp.StatusCode)
	}
}

// TestAddTasksInvalidToken tests /tasks with an invalid token
func TestAddTasksInvalidToken(t *testing.T) {
	db := &mockDB{}

	svc := &ToDoService{
		JWTKey: testJWTKey,
		Store:  db,
	}

	w := httptest.NewRecorder()

	r, err := http.NewRequest("POST", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	svc.ServeHTTP(w, r)

	if resp := w.Result(); resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("unexpected status code (want %d  have %d)", http.StatusUnauthorized, resp.StatusCode)
	}
}
