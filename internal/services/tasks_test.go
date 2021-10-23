package services

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/assert"
)

var date, _ = time.Parse("2006-01-02", "2020-01-29")

type mockStore struct{}

func (m *mockStore) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]storages.Task, error) {
	return []storages.Task{
		{
			ID:          "hi",
			Content:     "hihi",
			UserID:      "ex",
			CreatedDate: date.String(),
		},
	}, nil
}

func (m *mockStore) AddTask(ctx context.Context, t *storages.Task) error {
	return nil
}

func (m *mockStore) CountUserTasks(ctx context.Context, userID, createdDate sql.NullString) (int, error) {
	return 0, nil
}

func (m *mockStore) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	return false
}

func TestListTasks(t *testing.T) {
	tcs := []struct {
		name          string
		userID        string
		expectedTasks []Task
	}{
		{"list user's tasks successfully",
			"ex",
			[]Task{
				{
					ID:          "hi",
					Content:     "hihi",
					UserID:      "ex",
					CreatedDate: date.String(),
				},
			}},
	}

	svc := ToDoService{
		Store: &mockStore{},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			tasks, err := svc.ListTasks(context.Background(), tc.userID, date.String())
			assert.Nil(err)
			assert.Equal(tc.expectedTasks, tasks)
		})
	}
}
