package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/jericogantuangco/togo/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomContentTask() CreateTaskParams {
	stringLen := 5
	now := time.Now()
	date := now.Format("2006-01-02")
	return CreateTaskParams{
		Content: util.RandomString(stringLen),
		UserID: 12323,
		CreatedDate: date,
	}
}

func TestCreateTask(t *testing.T) {
	arg := createRandomContentTask()
	task , err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.Content, task.Content)
	require.Equal(t, arg.CreatedDate, task.CreatedDate)
}

func TestListTask(t *testing.T) {
	arg := createRandomContentTask()
	req := ListTasksParams{
		UserID: arg.UserID,
		CreatedDate: arg.CreatedDate,
	}
	tasks, err := testQueries.ListTasks(context.Background(), req)
	require.NoError(t, err)
	
	for _, task := range tasks {
		require.NotEmpty(t, task)
		require.Equal(t, arg.UserID, task.UserID)
		require.Equal(t, arg.CreatedDate, task.CreatedDate)
	}
}
