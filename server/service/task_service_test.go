package service

import (
	"testing"
	"time"
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

func (m *MockTaskRepository) CountTask(userid string, now time.Time) (int, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(int), args.Error(1)
}

func TestValidateEmptyTask(t *testing.T) {
	testService := NewTaskService(nil, nil)
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
	testService := NewTaskService(nil, nil)
	err := testService.Validate(&task)

	// Assert Nil
	assert.NotNil(t, err)

	// Assert Error message
	expected := "the task title is empty"
	actual := err.Error()
	assert.Equal(t, expected, actual)
}

func TestCreateTask(t *testing.T) {
	mockTaskRepository := new(MockTaskRepository)
	mockUserRepository := new(MockUserRepository)

	title := "title"
	description := "description"

	// Mock Task
	mocktask := models.Task{
		Title: title, Description: description,
	}

	// Setup expectation
	mockTaskRepository.On("CreateTask").Return(&mocktask, nil)

	testService := NewTaskService(mockTaskRepository, mockUserRepository)

	// Mock create task
	result, _ := testService.Create(&mocktask)

	// Mock Behavror Assertion
	mockTaskRepository.AssertExpectations(t)

	// Data Assertion
	assert.NotNil(t, result.CreatedAt)
	assert.NotNil(t, result.TaskID)
	assert.Equal(t, mocktask.Title, result.Title)
	assert.Equal(t, mocktask.Description, result.Description)
}

func TestGetLimitNoToken(t *testing.T) {
	testService := NewTaskService(nil, nil)

	token := ""
	err := testService.GetLimit(token)

	// Assert Nil
	assert.NotNil(t, err)

	// Assert Error message
	expected := "token is empty"
	actual := err.Error()
	assert.Equal(t, expected, actual)
}
