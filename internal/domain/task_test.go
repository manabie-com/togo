package domain

import (
	"context"
	"fmt"
	"testing"

	"manabie/togo/common/constants"
	"manabie/togo/common/errors"
	"manabie/togo/internal/model"
	"manabie/togo/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/* Test */

const (
	userID = "somethingID"
)

func Test_Task_Create_InvalidContext(t *testing.T) {
	userStore := &mockUserStore{}
	taskStore := &mockTaskStore{}
	taskCountStore := &mockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	task, err := taskDomain.Create(context.Background(), "something content")

	assert.Error(t, errors.ErrUserIDIsInvalid, err)
	assert.Nil(t, task)
}

func Test_Task_Create_UserNotExist(t *testing.T) {
	userStore := &mockUserStore{}
	taskStore := &mockTaskStore{}
	taskCountStore := &mockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	userStore.On("FindUser", mock.Anything).Return(nil, fmt.Errorf("something error"))
	ctx := utils.AddToContext(context.Background(), userID)
	task, err := taskDomain.Create(ctx, "something content")

	assert.Equal(t, errors.ErrUserDoesNotExist, err)
	assert.Nil(t, task)
}

func Test_Task_Create_CreateIfNotExistsFail(t *testing.T) {
	userStore := &mockUserStore{}
	taskStore := &mockTaskStore{}
	taskCountStore := &mockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	rErr := fmt.Errorf("redis disconnected or redis down")

	userStore.On("FindUser", mock.Anything).Return(&model.User{
		ID: userID,
	}, nil)
	taskCountStore.On("CreateIfNotExists", mock.Anything, mock.Anything).Return(rErr)

	ctx := utils.AddToContext(context.Background(), userID)

	task, err := taskDomain.Create(ctx, "something content")

	assert.Equal(t, err, rErr)
	assert.Nil(t, task)
}

func Test_Task_Create_TaskLimitExceeded(t *testing.T) {
	userStore := &mockUserStore{}
	taskStore := &mockTaskStore{}
	taskCountStore := &mockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)
	u := &model.User{
		ID:      userID,
		MaxTodo: 5,
	}
	ctx := utils.AddToContext(context.Background(), userID)
	userStore.On("FindUser", mock.Anything).Return(u, nil)

	taskCountStore.On("CreateIfNotExists", mock.Anything, mock.Anything).Return(nil)
	taskStore.On("AddTask", mock.Anything).Return(nil)

	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(u.MaxTodo+1, nil).Once()

	task, err := taskDomain.Create(ctx, "something content")
	assert.Equal(t, errors.ErrTaskLimitExceeded, err)
	assert.Nil(t, task)

	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(u.MaxTodo, nil).Once()

	task, err = taskDomain.Create(ctx, "something content")
	assert.NoError(t, err)
	assert.NotNil(t, task)

}

// func Test_Task_Create_AddTaskFail(t *testing.T) {
// 	userStore := &mockUserStore{}
// 	taskStore := &mockTaskStore{}
// 	taskCountStore := &mockTaskCountStore{}
// 	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

// 	rErr := fmt.Errorf("can not add task")
// 	u := &model.User{
// 		ID:      userID,
// 		MaxTodo: 5,
// 	}
// 	userStore.On("FindUser", mock.Anything).Return(u, nil)

// 	taskCountStore.On("CreateIfNotExists", mock.Anything, mock.Anything).Return(nil)
// 	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(u.MaxTodo-1, nil)

// 	taskStore.On("AddTask", mock.Anything).Return(rErr)

// 	ctx := utils.AddToContext(context.Background(), userID)

// 	task, err := taskDomain.Create(ctx, "something content")

// 	assert.Equal(t, rErr, err)
// 	assert.Nil(t, task)
// }

func Test_Task_Create_Success(t *testing.T) {
	userStore := &mockUserStore{}
	taskStore := &mockTaskStore{}
	taskCountStore := &mockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	u := &model.User{
		ID:      userID,
		MaxTodo: 5,
	}
	userStore.On("FindUser", mock.Anything).Return(u, nil)

	taskCountStore.On("CreateIfNotExists", mock.Anything, mock.Anything).Return(nil)
	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(u.MaxTodo-1, nil)

	taskStore.On("AddTask", mock.Anything).Return(nil)

	ctx := utils.AddToContext(context.Background(), userID)

	task, err := taskDomain.Create(ctx, "something content")

	assert.NoError(t, err)
	assert.NotNil(t, task)
}

func Test_Task_GetList_InvalidContext(t *testing.T) {
	userStore := &mockUserStore{}
	taskStore := &mockTaskStore{}
	taskCountStore := &mockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	tasks, err := taskDomain.GetList(context.Background(), "something content")

	assert.Error(t, errors.ErrUserIDIsInvalid, err)
	assert.Nil(t, tasks)
}

func Test_Task_GetList_UserNotExist(t *testing.T) {
	userStore := &mockUserStore{}
	taskStore := &mockTaskStore{}
	taskCountStore := &mockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	userStore.On("FindUser", mock.Anything).Return(nil, fmt.Errorf("something error"))
	ctx := utils.AddToContext(context.Background(), userID)
	tasks, err := taskDomain.GetList(ctx, "2021-07-31")

	assert.Equal(t, errors.ErrUserDoesNotExist, err)
	assert.Nil(t, tasks)
}

func Test_Task_GetList_RetrieveTasksFail(t *testing.T) {
	userStore := &mockUserStore{}
	taskStore := &mockTaskStore{}
	taskCountStore := &mockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)
	u := &model.User{
		ID:      userID,
		MaxTodo: 5,
	}
	userStore.On("FindUser", mock.Anything).Return(u, nil)

	taskCountStore.On("CreateIfNotExists", mock.Anything, mock.Anything).Return(nil)
	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(u.MaxTodo-1, nil)
	date := "2021-07-31"
	rErr := fmt.Errorf("RetrieveTasks error")
	taskStore.On("RetrieveTasks", mock.Anything).Return(nil, rErr)

	ctx := utils.AddToContext(context.Background(), userID)

	result, err := taskDomain.GetList(ctx, date)
	assert.Equal(t, rErr, err)
	assert.Nil(t, result)
}

func Test_Task_GetList_Success(t *testing.T) {
	userStore := &mockUserStore{}
	taskStore := &mockTaskStore{}
	taskCountStore := &mockTaskCountStore{}
	taskDomain := NewTaskDomain(taskCountStore, taskStore, userStore)

	userStore.On("FindUser", mock.Anything).Return(&model.User{
		ID: userID,
	}, nil)

	taskCountStore.On("CreateIfNotExists", mock.Anything, mock.Anything).Return(nil)
	taskCountStore.On("Inc", mock.Anything, mock.Anything).Return(constants.TaskLimit-1, nil)
	date := "2021-07-31"
	tasks := []*model.Task{
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
	taskStore.On("RetrieveTasks", mock.Anything).Return(tasks, nil)

	ctx := utils.AddToContext(context.Background(), userID)

	result, err := taskDomain.GetList(ctx, date)
	assert.NoError(t, err)
	assert.Equal(t, tasks, result)

}
