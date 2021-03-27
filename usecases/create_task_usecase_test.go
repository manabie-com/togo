package usecases

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/domains"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTaskUseCaseSuccess(t *testing.T) {
	ctx := context.Background()
	expected := transformTasksToSliceTaskOutput([]*domains.Task{{}})[0]
	mockRepo := new(DBMock)
	mockRepo.On("GetUserById", ctx, int64(0)).Return(&domains.User{MaxTodo: 5}, nil)
	mockRepo.On("GetCountCreatedTaskTodayByUser", ctx, int64(0)).Return(int64(4), nil)
	mockRepo.On("CreateTask", ctx, &domains.Task{}).Return(&domains.Task{}, nil)

	uc := NewCreateTaskUseCase(mockRepo, mockRepo)
	tsk, err := uc.Execute(ctx, TaskInput{})
	assert.Equal(t, expected, tsk)
	assert.Nil(t, err)
}

func TestCreateTaskErrorWithUserDoesNotExist(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(DBMock)
	mockRepo.On("GetUserById", ctx, int64(0)).Return(nil, domains.ErrorNotFound)

	uc := NewCreateTaskUseCase(mockRepo, mockRepo)
	tsk, err := uc.Execute(ctx, TaskInput{})

	mockRepo.AssertExpectations(t)
	assert.Equal(t, ErrorUserNotFound, err)
	assert.Nil(t, tsk)
}

func TestCreateTaskErrorWithReachedLimit(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(DBMock)
	mockRepo.On("GetUserById", ctx, int64(0)).Return(&domains.User{MaxTodo: 5}, nil)
	mockRepo.On("GetCountCreatedTaskTodayByUser", ctx, int64(0)).Return(int64(5), nil)

	uc := NewCreateTaskUseCase(mockRepo, mockRepo)
	tsk, err := uc.Execute(ctx, TaskInput{})

	mockRepo.AssertExpectations(t)
	assert.Nil(t, tsk)
	assert.Equal(t, ErrorReachedLimitCreateTaskPerDay, err)
}

func TestCreateTaskErrorInternalServerError(t *testing.T) {
	ctx := context.Background()
	expected := errors.New("internal error")

	mockRepo := new(DBMock)
	mockRepo.On("GetUserById", ctx, int64(0)).Return(nil, expected)

	uc := NewCreateTaskUseCase(mockRepo, mockRepo)
	tsk, err := uc.Execute(ctx, TaskInput{})

	mockRepo.AssertExpectations(t)
	assert.Nil(t, tsk)
	assert.Equal(t, expected, err)
}

func TestCreateTaskErrorInternalServerError2(t *testing.T) {
	ctx := context.Background()
	expected := errors.New("internal error")

	mockRepo := new(DBMock)
	mockRepo.On("GetUserById", ctx, int64(0)).Return(&domains.User{MaxTodo: 5}, nil)
	mockRepo.On("GetCountCreatedTaskTodayByUser", ctx, int64(0)).Return(int64(0), expected)

	uc := NewCreateTaskUseCase(mockRepo, mockRepo)
	tsk, err := uc.Execute(ctx, TaskInput{})

	mockRepo.AssertExpectations(t)
	assert.Nil(t, tsk)
	assert.Equal(t, expected, err)
}

func TestCreateTaskErrorInternalServerError3(t *testing.T) {
	ctx := context.Background()
	expected := errors.New("internal error")

	mockRepo := new(DBMock)
	mockRepo.On("GetUserById", ctx, int64(0)).Return(&domains.User{MaxTodo: 5}, nil)
	mockRepo.On("GetCountCreatedTaskTodayByUser", ctx, int64(0)).Return(int64(0), expected)

	uc := NewCreateTaskUseCase(mockRepo, mockRepo)
	tsk, err := uc.Execute(ctx, TaskInput{})

	mockRepo.AssertExpectations(t)
	assert.Nil(t, tsk)
	assert.Equal(t, expected, err)
}