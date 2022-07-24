package sql

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/trangmaiq/togo/internal/model"
)

func TestPersister_CreateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var (
		now  = time.Now()
		task = model.Task{
			ID:        uuid.NewString(),
			UserID:    uuid.NewString(),
			Title:     "1st task",
			Note:      "should store it properly",
			Status:    model.StatusInProgress,
			CreatedAt: now,
			UpdatedAt: now,
		}
	)

	mock.
		ExpectExec(`INSERT INTO tasks \(id, user_id, title, note, status, created_at, updated_at\)`).
		WithArgs(task.ID, task.UserID, task.Title, task.Note, task.Status, task.CreatedAt, task.UpdatedAt).
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))

	dbx := sqlx.NewDb(db, "postgres")
	persister := NewPersister(dbx)

	err = persister.CreateTask(context.TODO(), &task)
	require.NoError(t, err)

}
