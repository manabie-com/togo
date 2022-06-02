package service

import (
	"testing"
	"togo/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.Task), args.Error(1)
}

func TestValidateEmptyTask(t *testing.T) {
	testService := NewTaskService(nil)
	err := testService.Validate(nil)

	// Assert Nil
	assert.NotNil(t, err)

	// Assert Error message
	expected := "the task is empty"
	actual := err.Error()
	assert.Equal(t, expected, actual)
}

func TestValidateEmptyTaskTitle(t *testing.T) {
	task := models.Task{
		Title: "", Description: "description",
	}
	testService := NewTaskService(nil)
	err := testService.Validate(&task)

	// Assert Nil
	assert.NotNil(t, err)

	// Assert Error message
	expected := "the task title is empty"
	actual := err.Error()
	assert.Equal(t, expected, actual)
}

func TestCreateTask(t *testing.T) {
	mockRepository := new(MockTaskRepository)

	title := "title"
	description := "description"

	// Mock Task
	mocktask := models.Task{
		Title: title, Description: description,
	}

	// Setup expectation
	mockRepository.On("CreateTask").Return(&mocktask, nil)

	testService := NewTaskService(mockRepository)

	// Mock create task
	result, _ := testService.Create(&mocktask)

	// Mock Behavror Assertion
	mockRepository.AssertExpectations(t)

	// Data Assertion
	assert.NotNil(t, result.TaskID)
	assert.Equal(t, mocktask.Title, result.Title)
	assert.Equal(t, mocktask.Description, result.Description)
}
