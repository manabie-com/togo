package controllers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"manabie.com/internal/common"
	"manabie.com/internal/repositories"
	"context"
)

func TestTaskControllerCreateTaskForUserId(t *testing.T) {
	clockMock := common.MakeClockMock()
	clockMock.AddTimestamps(0)
	factory := repositories.MakeRepositoryFactoryMock()
	factory.InitUsers(10000)
	controller := MakeTaskController(factory, clockMock)
	
	ctx := context.Background()
	task, err := controller.CreateTaskForUserId(ctx, 1, "task-title", "task-content")
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(factory.TaskRepository.Tasks[1]), "task was not created")
	assert.Equal(t, "task-title", factory.TaskRepository.Tasks[1][0].Title)
	assert.Equal(t, "task-content", factory.TaskRepository.Tasks[1][0].Content)
	assert.Equal(t, 1, factory.TaskRepository.Tasks[1][0].Owner.Id)
	assert.Equal(t, factory.TaskRepository.Tasks[1][0], task)
	assert.Equal(t, repositories.TransactionLevel(repositories.Serializable), factory.TransactionLevelsHistory[factory.Count])
}

func TestTaskControllerCreateTaskForExceedLimit(t *testing.T) {
	clockMock := common.MakeClockMock()
	clockMock.AddTimestamps(0)
	factory := repositories.MakeRepositoryFactoryMock()
	factory.InitUsers(10000)
	controller := MakeTaskController(factory, clockMock)
	
	ctx := context.Background()
	for i := 0; i < 13; i++ {
		_, err := controller.CreateTaskForUserId(ctx, 10, "task-title", "task-content")
		if (i < 11) {
			assert.Equal(t, nil, err)
		} else {
			assert.Equal(t, TaskLimitExceeds, err)
		}
	}

	/// can still create for other users
	_, err := controller.CreateTaskForUserId(ctx, 11, "task-title", "task-content")
	assert.Equal(t, err, nil)
	assert.Equal(t, repositories.TransactionLevel(repositories.Serializable), factory.TransactionLevelsHistory[factory.Count])

	/// should reset after 1 day
	clockMock.AddTimestamps(24 * 3600 * 1000)
	for i := 0; i < 13; i++ {
		_, err := controller.CreateTaskForUserId(ctx, 10, "task-title", "task-content")
		if (i < 11) {
			assert.Equal(t, nil, err)
		} else {
			assert.Equal(t, TaskLimitExceeds, err)
		}
	}
}