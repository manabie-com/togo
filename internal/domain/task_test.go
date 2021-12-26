package domain

import (
	"context"
	"fmt"
	"testing"

	"github.com/perfectbuii/togo/common/constants"
	"github.com/perfectbuii/togo/common/errors"
	"github.com/perfectbuii/togo/internal/storages"
	"github.com/perfectbuii/togo/internal/storages/mocks"
	"github.com/perfectbuii/togo/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/* Test */

const (
	userID = "somethingID"
)

func Test_Task_Create_InvalidContext(t *testing.T) {
	userStore := &mocks.MockUserStore{}
	taskStore := &mocks.MockTaskStore{}
	taskCountStore := &mocks.MockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	task, err := taskDomain.Create(context.Background(), "something content")

	assert.Error(t, errors.ErrUserIdIsInvalid, err)
	assert.Nil(t, task)
}

func Test_Task_Create_UserNotExist(t *testing.T) {
	userStore := &mocks.MockUserStore{}
	taskStore := &mocks.MockTaskStore{}
	taskCountStore := &mocks.MockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	userStore.On("FindUser", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("something error"))
	ctx := utils.AddToContext(context.Background(), userID)
	task, err := taskDomain.Create(ctx, "something content")

	assert.Equal(t, errors.ErrUserDoesNotExist, err)
	assert.Nil(t, task)
}

func Test_Task_Create_TaskLimitExceeded(t *testing.T) {
	userStore := &mocks.MockUserStore{}
	taskStore := &mocks.MockTaskStore{}
	taskCountStore := &mocks.MockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)
	u := &storages.User{
		ID:      userID,
		MaxTodo: 5,
	}
	ctx := utils.AddToContext(context.Background(), userID)
	userStore.On("FindUser", mock.Anything, mock.Anything).Return(u, nil)
	taskStore.On("AddTask", mock.Anything).Return(nil)
	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(u.MaxTodo+1, nil).Once()

	task, err := taskDomain.Create(ctx, "something content")
	assert.Equal(t, errors.ErrTaskLimitExceeded, err)
	assert.Nil(t, task)
}

func Test_Task_Create_AddTaskFail(t *testing.T) {
	userStore := &mocks.MockUserStore{}
	taskStore := &mocks.MockTaskStore{}
	taskCountStore := &mocks.MockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	rErr := fmt.Errorf("can not add task")
	u := &storages.User{
		ID:      userID,
		MaxTodo: 5,
	}
	userStore.On("FindUser", mock.Anything, mock.Anything).Return(u, nil)

	taskCountStore.On("CreateIfNotExists", mock.Anything, mock.Anything).Return(nil)
	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(u.MaxTodo-1, nil)
	taskCountStore.On("Desc", mock.Anything, mock.Anything).Return(nil, nil)

	taskStore.On("AddTask", mock.Anything, mock.Anything).Return(rErr)

	ctx := utils.AddToContext(context.Background(), userID)
	task, err := taskDomain.Create(ctx, "something content")

	assert.Equal(t, rErr, err)
	assert.Nil(t, task)
}

func Test_Task_Create_Success(t *testing.T) {
	userStore := &mocks.MockUserStore{}
	taskStore := &mocks.MockTaskStore{}
	taskCountStore := &mocks.MockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	u := &storages.User{
		ID:      userID,
		MaxTodo: 5,
	}
	userStore.On("FindUser", mock.Anything, mock.Anything).Return(u, nil)

	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(u.MaxTodo-1, nil)

	taskStore.On("AddTask", mock.Anything, mock.Anything).Return(nil)

	ctx := utils.AddToContext(context.Background(), userID)

	task, err := taskDomain.Create(ctx, "something content")

	assert.NoError(t, err)
	assert.NotNil(t, task)
}

func Test_Task_GetList_InvalidContext(t *testing.T) {
	userStore := &mocks.MockUserStore{}
	taskStore := &mocks.MockTaskStore{}
	taskCountStore := &mocks.MockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	tasks, err := taskDomain.GetList(context.Background(), "something content")

	assert.Error(t, errors.ErrUserIdIsInvalid, err)
	assert.Nil(t, tasks)
}

func Test_Task_GetList_UserNotExist(t *testing.T) {
	userStore := &mocks.MockUserStore{}
	taskStore := &mocks.MockTaskStore{}
	taskCountStore := &mocks.MockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	userStore.On("FindUser", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("something error"))
	ctx := utils.AddToContext(context.Background(), userID)
	tasks, err := taskDomain.GetList(ctx, "2021-07-31")

	assert.Equal(t, errors.ErrUserDoesNotExist, err)
	assert.Nil(t, tasks)
}

func Test_Task_GetList_GetTasksFail(t *testing.T) {
	userStore := &mocks.MockUserStore{}
	taskStore := &mocks.MockTaskStore{}
	taskCountStore := &mocks.MockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)
	u := &storages.User{
		ID:      userID,
		MaxTodo: 5,
	}
	userStore.On("FindUser", mock.Anything, mock.Anything).Return(u, nil)

	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(u.MaxTodo-1, nil)
	date := "2021-12-26"
	rErr := fmt.Errorf("GetTasks error")
	taskStore.On("GetTasks", mock.Anything, mock.Anything).Return(nil, rErr)

	ctx := utils.AddToContext(context.Background(), userID)

	result, err := taskDomain.GetList(ctx, date)
	assert.Equal(t, rErr, err)
	assert.Nil(t, result)
}

func Test_Task_GetList_Success(t *testing.T) {
	userStore := &mocks.MockUserStore{}
	taskStore := &mocks.MockTaskStore{}
	taskCountStore := &mocks.MockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	userStore.On("FindUser", mock.Anything, mock.Anything).Return(&storages.User{
		ID: userID,
	}, nil)

	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(constants.TaskLimit-1, nil)
	date := "2021-07-31"
	tasks := []*storages.Task{
		{
			ID:          "id1",
			Content:     "something content",
			UserID:      userID,
			CreatedDate: date,
		},
		{
			ID:          "id2",
			Content:     "something content 2",
			UserID:      userID,
			CreatedDate: date,
		},

		{
			ID:          "id3",
			Content:     "something content 3",
			UserID:      userID,
			CreatedDate: date,
		},
	}
	taskStore.On("GetTasks", mock.Anything, mock.Anything).Return(tasks, nil)

	ctx := utils.AddToContext(context.Background(), userID)

	result, err := taskDomain.GetList(ctx, date)
	assert.NoError(t, err)
	assert.Equal(t, tasks, result)

}
