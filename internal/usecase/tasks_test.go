package usecase_test

import (
	"context"
	"testing"
	"time"

	mocks "github.com/manabie-com/togo/internal/mock"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListTasks(t *testing.T) {
	mockTaskDB := new(mocks.MockedTaskDB)
	mockTask := storages.Task{
		Content: "Learning Go",
		UserID:  "first user",
	}
	mockListTask := make([]storages.Task, 0)
	mockListTask = append(mockListTask, mockTask)
	t.Run("success", func(t *testing.T) {
		mockTaskDB.On("RetrieveTasks", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(mockListTask, nil).Once()

		taskUs := usecase.NewTaskUsecase(mockTaskDB, time.Second*2)
		listTask, err := taskUs.ListTasks(context.TODO(), "first user", time.Now())
		assert.Len(t, listTask, len(mockListTask))
		assert.NoError(t, err)
		mockTaskDB.AssertExpectations(t)
	})
}

func TestAddTask(t *testing.T) {

	mockTaskDB := new(mocks.MockedTaskDB)
	mockTask := storages.Task{
		ID:          1,
		Content:     "Learning Go",
		UserID:      "first user",
		CreatedDate: time.Now(),
	}
	t.Run("success", func(t *testing.T) {
		tempMockTask := mockTask
		mockTaskDB.On("AddTask", mock.Anything, mock.AnythingOfType("*storages.Task")).Return(nil).Once()
		taskUs := usecase.NewTaskUsecase(mockTaskDB, time.Second*2)
		err := taskUs.AddTask(context.TODO(), &tempMockTask)
		assert.NoError(t, err)
		assert.Equal(t, mockTask.ID, tempMockTask.ID)
		mockTaskDB.AssertExpectations(t)
	})
}

func TestValidateUser(t *testing.T) {
	mockTaskDB := new(mocks.MockedTaskDB)
	userID := "user"
	pass := "password"
	t.Run("success", func(t *testing.T) {
		mockTaskDB.On("ValidateUser", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(true, nil).Once()
		taskUs := usecase.NewTaskUsecase(mockTaskDB, time.Second*2)
		val, err := taskUs.ValidateUser(context.TODO(), userID, pass)
		assert.NoError(t, err)
		assert.NotNil(t, val)
		mockTaskDB.AssertExpectations(t)
	})
}
