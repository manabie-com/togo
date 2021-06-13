package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages/ent"
	"github.com/manabie-com/togo/internal/storages/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestTaskService_CreateTask(t *testing.T) {
	mockUserRepos := &mocks.UserRepositoryMock{}
	mockTaskRepos := &mocks.TaskRepositoryMock{}

	taskService := NewTaskService(mockTaskRepos, mockUserRepos)

	ctx := context.WithValue(context.Background(), "userId", uuid.NewString())

	mockUserRepos.On("FindByUserId", ctx, mock.Anything).Return(&ent.User{Password: "$2a$10$DZjcIPh9cv.cWH62dYII0uaYsPjvSCR4hMfMBNl4GrSaktw7vaQ2O"}, nil)
	mockTaskRepos.On("CreateTask", ctx, mock.Anything, mock.Anything).Return(&ent.Task{}, nil)

	newTask, err := taskService.CreateTask(ctx, model.TaskCreationRequest{
		Content: "content",
	})

	assert.Nil(t, err)
	assert.NotNil(t, newTask)

	mockUserRepos.AssertExpectations(t)
}

func TestTaskService_GetTaskByDate(t *testing.T) {
	mockUserRepos := &mocks.UserRepositoryMock{}
	mockTaskRepos := &mocks.TaskRepositoryMock{}

	taskService := NewTaskService(mockTaskRepos, mockUserRepos)

	ctx := context.WithValue(context.Background(), "userId", uuid.NewString())

	mockTaskRepos.On("GetTaskByDate", ctx, mock.Anything, mock.Anything, mock.Anything).Return([]*ent.Task{}, nil)

	newTask, err := taskService.GetTaskByDate(ctx, time.Now())

	assert.Nil(t, err)
	assert.NotNil(t, newTask)

	mockUserRepos.AssertExpectations(t)
}
