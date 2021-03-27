package usecases

import (
	"context"
	"github.com/manabie-com/togo/domains"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTasksUseCaseSuccess(t *testing.T) {
	ctx := context.Background()
	expected := transformTasksToSliceTaskOutput([]*domains.Task{{}})
	mockRepo := new(DBMock)
	mockRepo.On("GetTasks", ctx, &domains.TaskRequest{}).Return([]*domains.Task{{}}, nil)

	uc := NewGetTasksUseCase(mockRepo)
	tsk, err := uc.Execute(ctx, GetTasksInput{})

	mockRepo.AssertExpectations(t)
	assert.Equal(t, expected, tsk)
	assert.Nil(t, err)
}

func TestGetTasksErrorNotFound(t *testing.T) {
	ctx := context.Background()
	expected := domains.ErrorNotFound
	mockRepo := new(DBMock)
	mockRepo.On("GetTasks", ctx, &domains.TaskRequest{}).Return(nil, expected)

	uc := NewGetTasksUseCase(mockRepo)
	tsk, err := uc.Execute(ctx, GetTasksInput{})

	mockRepo.AssertExpectations(t)
	assert.Equal(t, expected, err)
	assert.Nil(t, tsk)
}