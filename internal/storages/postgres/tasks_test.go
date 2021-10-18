package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/jericogantuangco/togo/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomContentTask(user User) CreateTaskParams {
	stringLen := 5
	now := time.Now()
	date := now.Format("2006-01-02")
	return CreateTaskParams{
		Content:     util.RandomString(stringLen),
		UserID:      user.Username,
		CreatedDate: date,
	}
}

func createTestUser() User {
	arg := CreateUserParams {
		Username: "testUser",
		Password: "password",
	}
	user, err := testQueries.RetrieveUser(context.Background(),"testUser")
	if err != nil {
		user, _ := testQueries.CreateUser(context.Background(), arg)
		return user
	}
	return user
}

func TestCreateTask(t *testing.T) {
	user := createTestUser()
	arg := createRandomContentTask(user)
	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.Content, task.Content)
	require.Equal(t, arg.CreatedDate, task.CreatedDate)
}

func TestListTask(t *testing.T) {
	user := createTestUser()
	arg := createRandomContentTask(user)
	req := RetrieveTasksParams{
		UserID:      arg.UserID,
		CreatedDate: arg.CreatedDate,
	}
	tasks, err := testQueries.RetrieveTasks(context.Background(), req)
	require.NoError(t, err)

	for _, task := range tasks {
		require.NotEmpty(t, task)
		require.Equal(t, arg.UserID, task.UserID)
		require.Equal(t, arg.CreatedDate, task.CreatedDate)
	}
}
