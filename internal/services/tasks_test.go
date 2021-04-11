package services

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

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
