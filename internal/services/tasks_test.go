package services

import (
	"context"
	"database/sql"
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
