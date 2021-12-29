package services_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/entities"
	"github.com/manabie-com/togo/helpers"
	repositories_mocks "github.com/manabie-com/togo/mocks/repositories"
	"github.com/manabie-com/togo/services"
	"github.com/stretchr/testify/assert"
)

func TestTaskService_GetAllTask(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	expectedVal := []entities.Task{{
		ID:          1,
		UserId:      1,
		Title:       "Test",
		Description: "Test",
		IsCompleted: false,
		CreatedAt:   helpers.GetDateNow(),
	},
		{
			ID:          2,
			UserId:      1,
			Title:       "Test",
			Description: "Test",
			IsCompleted: false,
			CreatedAt:   helpers.GetDateNow(),
		}}

	mockTaskRepo := repositories_mocks.NewMockITaskRepository(ctl)
	mockUserRepo := repositories_mocks.NewMockIUserRepository(ctl)

	t.Run("Test: Service's GetAllTask success", func(t *testing.T) {
		mockTaskRepo.EXPECT().GetAllTask().Return(expectedVal, nil)
		taskService := services.TaskService{TaskRepo: mockTaskRepo, UserRepo: mockUserRepo}

		result, err := taskService.GetAllTask()
		assert.Nil(t, nil, err)
		assert.Equal(t, result, expectedVal)
	})

	t.Run("Test: Service's GetAllTask error", func(t *testing.T) {
		mockTaskRepo.EXPECT().GetAllTask().Return(nil, errors.New("an error"))
		taskService := services.TaskService{TaskRepo: mockTaskRepo, UserRepo: mockUserRepo}

		result, err := taskService.GetAllTask()
		assert.NotNil(t, err)
		assert.Nil(t, nil, result)
	})
}

func TestTaskService_CreateTask(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	inputVal := entities.Task{
		ID:          1,
		UserId:      1,
		Title:       "Test",
		Description: "Test",
		IsCompleted: false,
		CreatedAt:   helpers.GetDateNow(),
	}

	mockTaskRepo := repositories_mocks.NewMockITaskRepository(ctl)
	mockUserRepo := repositories_mocks.NewMockIUserRepository(ctl)

	t.Run("Test: Service's CreateTask success", func(t *testing.T) {
		mockUserRepo.EXPECT().GetLimitTaskPerDay(inputVal.UserId).Return(2, nil)                                       //LIMIT TASKS PER DAY: 2
		mockTaskRepo.EXPECT().CountTaskByUserIdAndCreatedAt(inputVal.UserId, inputVal.CreatedAt).Return(int64(1), nil) //TODAY TASKS: 1
		mockTaskRepo.EXPECT().CreateTask(&inputVal).Return(nil)

		taskService := services.TaskService{TaskRepo: mockTaskRepo, UserRepo: mockUserRepo}

		internalErr, userErr := taskService.CreateTask(&inputVal)
		assert.Nil(t, nil, internalErr)
		assert.Nil(t, nil, userErr)
	})

	t.Run("Test: Service's GetAllTask internal error", func(t *testing.T) {
		mockUserRepo.EXPECT().GetLimitTaskPerDay(inputVal.UserId).Return(2, nil)                                       //LIMIT TASKS PER DAY: 2
		mockTaskRepo.EXPECT().CountTaskByUserIdAndCreatedAt(inputVal.UserId, inputVal.CreatedAt).Return(int64(1), nil) //TODAY TASKS: 1
		mockTaskRepo.EXPECT().CreateTask(&inputVal).Return(errors.New("an internal error"))

		taskService := services.TaskService{TaskRepo: mockTaskRepo, UserRepo: mockUserRepo}

		internalErr, userErr := taskService.CreateTask(&inputVal)
		assert.NotNil(t, internalErr)
		assert.Nil(t, nil, userErr)
	})

	t.Run("Test: Service's GetAllTask user error (the task created today exceeded the allowed limit)", func(t *testing.T) {
		mockUserRepo.EXPECT().GetLimitTaskPerDay(inputVal.UserId).Return(2, nil)                                       //LIMIT TASKS PER DAY: 2
		mockTaskRepo.EXPECT().CountTaskByUserIdAndCreatedAt(inputVal.UserId, inputVal.CreatedAt).Return(int64(2), nil) //TODAY TASKS: 2

		taskService := services.TaskService{TaskRepo: mockTaskRepo, UserRepo: mockUserRepo}

		internalErr, userErr := taskService.CreateTask(&inputVal)
		assert.Nil(t, nil, internalErr)
		assert.NotNil(t, userErr)
	})
}
