//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/trangmaiq/togo/internal/config"
	"github.com/trangmaiq/togo/internal/model"
	"github.com/trangmaiq/togo/internal/persistence/persister/sql"
)

func TestPersister_CreateTask(t *testing.T) {
	cfg := config.Load()
	db, err := sqlx.Connect(cfg.DBDriver, cfg.DSN)
	require.NoError(t, err)
	require.NoError(t, db.Ping())

	d := time.Date(2022, time.July, 30, 0, 0, 0, 0, time.FixedZone("", 0))
	task := model.Task{
		ID:        uuid.NewString(),
		UserID:    uuid.NewString(),
		Title:     "1st task",
		Note:      "togo should have a integration test",
		Status:    model.StatusInProgress,
		CreatedAt: d,
		UpdatedAt: d,
	}

	persister := sql.NewPersister(db)
	err = persister.CreateTask(context.TODO(), &task)
	require.NoError(t, err)

	recordedTask, err := persister.GetTask(context.TODO(), task.ID)
	require.NoError(t, err)
	require.Equal(t, task, *recordedTask)
}
