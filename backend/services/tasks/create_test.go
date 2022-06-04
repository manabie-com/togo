package tasks_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jssoriao/todo-go/services/tasks"
	"github.com/jssoriao/todo-go/storage"
)

type mockTasksStore struct{}

func (m mockTasksStore) CreateTask(task storage.Task) (storage.Task, error) {
	timestamp := time.Now()
	task.ID = "taskId"
	task.Created = timestamp
	task.Updated = timestamp
	return task, nil
}

func (m mockTasksStore) CountTasksForTheDay(userID string, dueDate time.Time) (int, error) {
	return 10, nil
}

func (m mockTasksStore) GetUser(id string) (*storage.User, error) {
	timestamp := time.Now()
	if id == "valid" {
		return &storage.User{
			ID:         id,
			DailyLimit: 11,
			Created:    timestamp,
			Updated:    timestamp,
		}, nil
	}
	if id == "exceeded" {
		return &storage.User{
			ID:         id,
			DailyLimit: 10,
			Created:    timestamp,
			Updated:    timestamp,
		}, nil
	}
	return nil, errors.New("non-nil error")
}

func TestCreateTask(t *testing.T) {
	timestamp := time.Now()
	tests := []struct {
		name    string
		req     *tasks.Task
		want    *tasks.Task
		wantErr error
	}{
		{
			name: "Valid",
			req: &tasks.Task{
				UserID:  "valid",
				Title:   "Valid Title",
				DueDate: timestamp,
			},
			want: &tasks.Task{
				UserID:  "valid",
				ID:      "taskId",
				Title:   "Valid Title",
				DueDate: timestamp,
			},
			wantErr: nil,
		},
		{
			name: "Exceeded",
			req: &tasks.Task{
				UserID:  "exceeded",
				Title:   "Exceeded Title",
				DueDate: time.Now(),
			},
			want:    nil,
			wantErr: tasks.ErrExceededTasksLimit,
		},
	}

	errComparer := cmp.Comparer(func(x, y error) bool {
		if x == nil || y == nil {
			if x == nil && y == nil {
				return true
			}
			return false
		}
		return x.Error() == y.Error()
	})

	store := mockTasksStore{}
	h, _ := tasks.NewHandler(store)
	for _, test := range tests {
		task, err := h.CreateTask(context.Background(), test.req)
		if diff := cmp.Diff(err, test.wantErr, errComparer); diff != "" {
			t.Errorf("error mismatch %s", diff)
		}
		if diff := cmp.Diff(task, test.want); diff != "" {
			t.Errorf("result mismatch %s", diff)
		}
	}
}
