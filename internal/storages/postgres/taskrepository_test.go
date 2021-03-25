package postgres

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/require"
)

func TestAddTask(t *testing.T) {
	err := testQueries.AddTask(context.Background(), randomTask())
	require.Equal(t, nil, err)
}

func TestRetrieveTasks(t *testing.T) {
	taskList, err := testQueries.RetrieveTasks(context.Background(), createRTTaskParams())

	require.Empty(t, err)
	require.NotEmpty(t, taskList)
}

func createRTTaskParams() RetrieveTasksParams {
	param := RetrieveTasksParams{
		UserID:      sql.NullString{String: "firstUser", Valid: true},
		CreatedDate: sql.NullString{String: "2021-03-23", Valid: true},
	}

	return param
}

func randomTask() *storages.Task {
	task := storages.Task{
		ID:          uuid.New().String(),
		Content:     "unit test content",
		UserID:      "firstUser",
		CreatedDate: "2021-03-23",
	}

	return &task
}
