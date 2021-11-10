package domain_test

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/domain/entity"
	"github.com/manabie-com/togo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var userEntity = &entity.User{
	Username:       "firstUser",
	HashedPassword: "123445",
	MaxTodo:        5,
}

func TestCreateTaskSuccess(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockTaskRepository.On("CountTaskInDayByUsername", mock.Anything, mock.AnythingOfType("string")).Return(4, nil)
	mockTaskRepository.On("Create", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
	mockUserRepository.On("GetUser", mock.Anything, mock.AnythingOfType("string")).Return(userEntity, nil)
	uc := domain.NewTaskUseCase(mockTaskRepository, mockUserRepository)
	err := uc.CreateTask(context.TODO(), "mtuan", "firstUser")
	require.NoError(t, err)
}

func TestCreateTaskFailedMaxTaskPerDay(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockTaskRepository.On("CountTaskInDayByUsername", mock.Anything, mock.AnythingOfType("string")).Return(5, nil)
	mockTaskRepository.On("Create", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
	mockUserRepository.On("GetUser", mock.Anything, mock.AnythingOfType("string")).Return(userEntity, nil)
	uc := domain.NewTaskUseCase(mockTaskRepository, mockUserRepository)
	err := uc.CreateTask(context.TODO(), "mtuan", "firstUser")
	require.EqualError(t, err, domain.ErrorMaximumTaskPerDay.Error())
}

func TestCreateTaskFailedCountTaskPerDay(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockTaskRepository.On("CountTaskInDayByUsername", mock.Anything, mock.AnythingOfType("string")).Return(0, errors.New("errors"))
	mockTaskRepository.On("Create", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
	mockUserRepository.On("GetUser", mock.Anything, mock.AnythingOfType("string")).Return(userEntity, nil)
	uc := domain.NewTaskUseCase(mockTaskRepository, mockUserRepository)
	err := uc.CreateTask(context.TODO(), "mtuan", "firstUser")
	require.Error(t, err)
}

func TestCreateTaskFailedCreateTaskRepo(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockTaskRepository.On("CountTaskInDayByUsername", mock.Anything, mock.AnythingOfType("string")).Return(1, nil)
	mockTaskRepository.On("Create", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(errors.New("errors"))
	mockUserRepository.On("GetUser", mock.Anything, mock.AnythingOfType("string")).Return(userEntity, nil)
	uc := domain.NewTaskUseCase(mockTaskRepository, mockUserRepository)
	err := uc.CreateTask(context.TODO(), "mtuan", "firstUser")
	require.Error(t, err)
}

func TestCreateTaskErrorGetUserError(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	mockUserRepository := new(mocks.UserRepository)
	mockTaskRepository.On("CountTaskInDayByUsername", mock.Anything, mock.AnythingOfType("string")).Return(1, nil)
	mockTaskRepository.On("Create", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(errors.New("errors"))
	mockUserRepository.On("GetUser", mock.Anything, mock.AnythingOfType("string")).Return(nil, domain.UserNotFound)
	uc := domain.NewTaskUseCase(mockTaskRepository, mockUserRepository)
	err := uc.CreateTask(context.TODO(), "mtuan", "firstUser")
	require.Error(t, err)
	require.EqualError(t, err, domain.UserNotFound.Error())
}
