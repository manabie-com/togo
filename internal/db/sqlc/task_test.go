package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateTask(t *testing.T) {
	user, err := createRandomUser(t)
	arg := CreateTaskParams{
		Title:  "Title 1",
		UserID: user.ID,
	}
	var task Task
	task, err = testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)
	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.UserID, task.UserID)
}
