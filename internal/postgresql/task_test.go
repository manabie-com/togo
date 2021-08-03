package postgresql

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateTask(t *testing.T) Task {
	user := CreateRandomUser(t)

	arg := InsertTaskParams{
		Content:     "Do some test",
		UserID:      user.ID,
		CreatedDate: time.Now(),
	}

	task, err := testQueries.InsertTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.Content, task.Content)
	require.Equal(t, arg.UserID, task.UserID)

	return task
}

func TestQueries_InsertTask(t *testing.T) {
	CreateTask(t)
}

func TestQueries_GetTask(t *testing.T) {
	task := CreateTask(t)

	user, err := testQueries.GetUser(context.Background(), task.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	arg := GetTaskParams{
		ID:     task.ID,
		UserID: user.ID,
	}

	task2, err := testQueries.GetTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task2.Content, task.Content)
	require.Equal(t, task2.UserID, task.UserID)
	require.Equal(t, task2.CreatedDate, task.CreatedDate)
}

func TestQueries_UpdateTask(t *testing.T) {
	task := CreateTask(t)

	user, err := testQueries.GetUser(context.Background(), task.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	arg := UpdateTaskParams{
		ID:     task.ID,
		IsDone: true,
	}

	task2, err := testQueries.UpdateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task2.Content, task.Content)
	require.Equal(t, task2.UserID, task.UserID)
	require.Equal(t, task2.CreatedDate, task.CreatedDate)
	require.Equal(t, task2.IsDone, true)
}

func TestQueries_DeleteTask(t *testing.T) {
	task := CreateTask(t)
	arg := DeleteTaskParams{
		ID:     task.ID,
		UserID: task.UserID,
	}

	err := testQueries.DeleteTask(context.Background(), arg)
	require.NoError(t, err)

	getArg := GetTaskParams{
		ID:     task.ID,
		UserID: task.UserID,
	}
	task2, err := testQueries.GetTask(context.Background(), getArg)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, task2)
}

func TestQueries_ListTasks(t *testing.T) {
	var lastTask Task
	for i := 0; i < 10; i++ {
		lastTask = CreateTask(t)
	}

	arg := ListTasksParams{
		UserID:      lastTask.UserID,
		CreatedDate: time.Now(),
		IsDone:      false,
	}

	tasks, err := testQueries.ListTasks(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tasks)

	for _, task := range tasks {
		require.NotEmpty(t, task)
		require.Equal(t, lastTask.UserID, task.UserID)
	}
}
