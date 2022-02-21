package repo_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/chi07/todo/internal/model"
	"github.com/chi07/todo/internal/repo"
	"github.com/chi07/todo/tests/pgtest"
)

var (
	util *dbutils.TestUtil
	ctx  = context.Background()
)

func init() {
	testutil := dbutils.New()
	if err := testutil.InitDB(); err != nil {
		testutil.Log.Panicf("testutil.initDB(): %v", err)
	}
	util = testutil
}

func TestTask_CountUserTasks(t *testing.T) {
	if err := util.SetupDB(); err != nil {
		util.Log.Panicf("util.SetupDB(): %v", err)
	}
	defer util.CleanAndClose()

	tests := []struct {
		name    string
		userID  uuid.UUID
		wantErr error
		want    int64
	}{
		{
			name:   "wrong userID",
			userID: uuid.New(),
			want:   int64(0),
		},
		{
			name:   "with valid userID",
			userID: uuid.MustParse("dc334e08-3842-4dac-9338-8f30ac5e2369"),
			want:   int64(0),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			taskRepo := repo.NewTask(util.PostgresDB)
			actual, err := taskRepo.CountUserTasks(ctx, tc.userID)
			assert.Equal(t, tc.want, actual)
			assert.NoError(t, err)
		})
	}
}

func TestTask_Create(t *testing.T) {
	if err := util.SetupDB(); err != nil {
		util.Log.Panicf("util.SetupDB(): %v", err)
	}
	defer util.CleanAndClose()
	taskID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	tests := []struct {
		name    string
		task    model.Task
		wantErr error
		want    uuid.UUID
	}{
		{
			name: "error: wrong priority",
			task: model.Task{
				ID:       uuid.New(),
				Title:    "Meeting",
				Status:   "TODO",
				Priority: "abc-xyz",
			},
			wantErr: errors.New(`pq: invalid input value for enum task_priorities: "abc-xyz"`),
			want:    uuid.Nil,
		},
		{
			name: "error: when duplicate task ID",
			task: model.Task{
				ID:       uuid.MustParse("dc334e08-3842-4dac-9338-8f30ac5e2369"),
				Title:    "Meeting",
				Status:   "TODO",
				Priority: "NORMAL",
			},
			wantErr: errors.New(`pq: duplicate key value violates unique constraint "tasks_pkey"`),
			want:    uuid.Nil,
		},
		{
			name: "success",
			task: model.Task{
				ID:       taskID,
				Title:    "Code",
				Status:   "TODO",
				Priority: "URGENT",
			},
			want: taskID,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			taskRepo := repo.NewTask(util.PostgresDB)
			actual, err := taskRepo.Create(ctx, &tc.task)
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("taskRepo.Create got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)
		})
	}
}
