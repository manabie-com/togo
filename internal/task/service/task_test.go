package service

import (
	"testing"
	"time"
	"togo/internal/models"
	"togo/internal/task/dto"
	"togo/internal/task/mocks"
	userMocks "togo/internal/user/mocks"

	"github.com/test-go/testify/assert"
	"gorm.io/gorm"
)

func TestTaskService_CreateTaskSuccess(t *testing.T) {
	timeEndedAt := time.Date(2022, 11, 21, 0, 0, 0, 0, time.Local)
	createTaskDto := &dto.CreateTaskDto{
		Description: "description",
		EndedAt:     timeEndedAt,
	}
	user := &models.User{
		BaseModelID: models.BaseModelID{ID: 1},
		Name:        "name",
		LimitCount:  1,
	}
	userID := int(user.ID)
	task := &models.Task{
		UserID:      userID,
		Description: createTaskDto.Description,
		EndedAt:     createTaskDto.EndedAt,
	}

	userRepo := userMocks.NewUserRepository(t)
	userRepo.On("GetByID", userID).Return(user, nil)

	taskRepo := mocks.NewTaskRepository(t)
	taskRepo.On("GetNumberOfUserTaskOnToday", userID).Return(0, nil)
	taskRepo.On("Create", task).Return(task, nil)

	service := NewTaskService(taskRepo, userRepo)

	taskResponse, err := service.Create(createTaskDto, userID)
	assert.Nil(t, err)
	assert.Equal(t, createTaskDto.Description, taskResponse.Description)
	assert.Equal(t, createTaskDto.EndedAt, taskResponse.EndedAt)
	assert.Equal(t, userID, taskResponse.UserID)
}

func TestTaskService_CreateTaskNotFoundUser(t *testing.T) {
	userID := 1
	createTaskDto := &dto.CreateTaskDto{
		Description: "description",
		EndedAt:     time.Now(),
	}

	userRepo := userMocks.NewUserRepository(t)
	userRepo.On("GetByID", userID).Return(nil, gorm.ErrRecordNotFound)

	taskRepo := mocks.NewTaskRepository(t)

	service := NewTaskService(taskRepo, userRepo)

	_, err := service.Create(createTaskDto, userID)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "user_not_found")
}

func TestTaskService_CreateTaskMaxLimitCount(t *testing.T) {
	userID := 1
	createTaskDto := &dto.CreateTaskDto{
		Description: "description",
		EndedAt:     time.Now(),
	}
	user := &models.User{
		BaseModelID: models.BaseModelID{ID: uint64(userID)},
		Name:        "name",
		LimitCount:  1,
	}

	userRepo := userMocks.NewUserRepository(t)
	userRepo.On("GetByID", userID).Return(user, nil)

	taskRepo := mocks.NewTaskRepository(t)
	taskRepo.On("GetNumberOfUserTaskOnToday", userID).Return(1, nil)

	service := NewTaskService(taskRepo, userRepo)

	_, err := service.Create(createTaskDto, userID)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "limit_max")
}
