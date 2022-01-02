package task_test

import (
	"context"
	"testing"

	"github.com/manabie-com/togo/app/model"
	"github.com/manabie-com/togo/app/utils"

	taskRepo "github.com/manabie-com/togo/app/repo/mongo/task"

	"github.com/stretchr/testify/require"
)

func createRandomTask(t *testing.T) model.Task {

	arg := taskRepo.CreateReq{
		UserID:      int(utils.RandomInt(1, 10)),
		Name:        utils.RandomString(10),
		Description: utils.RandomString(10),
		// tracing
		CreatedIP: "127.0.0.1",
	}

	task, err := taskRepoInstance.Create(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.UserID, task.UserID)
	require.Equal(t, arg.Name, task.Name)
	require.Equal(t, arg.Description, task.Description)
	require.NotZero(t, task.CreatedDate)

	return task
}

func TestCreateUser(t *testing.T) {
	createRandomTask(t)
}
