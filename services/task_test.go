package services

import (
	"github.com/manabie-com/togo/mocks"
	"github.com/manabie-com/togo/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetTasksByUserName(t *testing.T) {
	taskRepo := new(mocks.ITaskRepository)

	task := models.Task{
		Username:  "huyha",
		Content:   "Hello World",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	expectedTasks := &[]models.Task{task}

	c := taskRepo.On("GetTasksByUserName", task.Username, "").Return(expectedTasks, nil)

	require.Equal(t, []*mock.Call{c}, taskRepo.ExpectedCalls)

	taskService := TaskService{taskRepo}

	actualTasks, err := taskService.GetTasksByUserName(task.Username, "")

	if err != nil {
		t.Errorf("expected %q but got %q", expectedTasks, actualTasks)
	}

	assert.Equal(t, expectedTasks, actualTasks)
}

func TestCount(t *testing.T) {
	taskRepo := new(mocks.ITaskRepository)

	task := models.Task{
		Username:  "huyha",
		Content:   "Hello World",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var expectedCount int64 = 2

	c := taskRepo.On("Count", task.Username).Return(expectedCount, nil)

	require.Equal(t, []*mock.Call{c}, taskRepo.ExpectedCalls)

	taskService := TaskService{taskRepo}

	actualCount, err := taskService.Count(task.Username)

	if err != nil {
		t.Errorf("expected %q but got %q", expectedCount, actualCount)
	}

	assert.Equal(t, expectedCount, actualCount)
}

func TestCreateTask(t *testing.T) {
	taskRepo := new(mocks.ITaskRepository)

	task := models.Task{
		Username:  "huyha",
		Content:   "Hello World",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	expectedTask := &task

	c := taskRepo.On("CreateTask", expectedTask).Return(expectedTask, nil)

	require.Equal(t, []*mock.Call{c}, taskRepo.ExpectedCalls)

	taskService := TaskService{taskRepo}

	actualTask, err := taskService.CreateTask(expectedTask)

	if err != nil {
		t.Errorf("expected %q but got %q", expectedTask, actualTask)
	}

	assert.Equal(t, expectedTask, actualTask)
}
