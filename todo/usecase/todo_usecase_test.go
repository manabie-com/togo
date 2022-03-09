package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/triet-truong/todo/todo/dto"
	"github.com/triet-truong/todo/todo/model"
	"github.com/triet-truong/todo/todo/repository/mocks"
	"github.com/triet-truong/todo/todo/usecase"
)

type TodoUsecaseTestSuite struct {
	suite.Suite
}

func (s *TodoUsecaseTestSuite) TestGivenUserFoundInCache_WhenAddingItem_ShouldReturnNoError() {
	userId := uint(1)
	t := s.T()
	mockCacheRepo := new(mocks.TodoCacheRepository)
	mockCacheRepo.On("GetCachedUser", uint(1)).Return(model.UserRedisModel{
		ID:               userId,
		CurrentUsage:     1,
		DailyRecordLimit: 5,
	}, nil)
	mockCacheRepo.On("SetUser", model.UserRedisModel{
		ID:               userId,
		CurrentUsage:     2,
		DailyRecordLimit: 5,
	}).Return(nil)

	mockRepo := new(mocks.TodoRepository)
	mockRepo.On("InsertItem", model.TodoItemModel{
		Content: "Hello",
		UserID:  userId,
		IsDone:  false,
	}).Return(nil)

	u := usecase.NewTodoUseCase(mockRepo, mockCacheRepo)
	err := u.AddTodo(dto.TodoDto{
		ID:      1,
		Content: "Hello",
		UserId:  userId,
	})
	assert.NoError(t, err)
	mockCacheRepo.AssertCalled(t, "GetCachedUser", userId)
	mockRepo.AssertNumberOfCalls(t, "SelectUser", 0)

	mockCacheRepo.AssertNumberOfCalls(t, "SetUser", 1)
	mockRepo.AssertNumberOfCalls(t, "InsertItem", 1)
}

func (s *TodoUsecaseTestSuite) TestGivenUserNotFoundInDB_WhenAddingItem_ShouldReturnError() {
	userId := uint(1)
	t := s.T()
	mockCacheRepo := new(mocks.TodoCacheRepository)
	mockCacheRepo.On("GetCachedUser", uint(1)).Return(model.UserRedisModel{
		ID:               userId,
		CurrentUsage:     1,
		DailyRecordLimit: 5,
	}, errors.New("not found"))
	mockCacheRepo.On("SetUser", model.UserRedisModel{
		ID:               userId,
		CurrentUsage:     1,
		DailyRecordLimit: 5,
	}).Return(nil)

	mockRepo := new(mocks.TodoRepository)
	mockRepo.On("InsertItem", model.TodoItemModel{
		Content: "Hello",
		UserID:  userId,
		IsDone:  false,
	}).Return(nil)
	mockRepo.On("SelectUser", userId).Return(model.UserModel{}, errors.New("user not found"))

	u := usecase.NewTodoUseCase(mockRepo, mockCacheRepo)
	err := u.AddTodo(dto.TodoDto{
		ID:      1,
		Content: "Hello",
		UserId:  userId,
	})
	assert.Error(t, err)
	mockCacheRepo.AssertCalled(t, "GetCachedUser", userId)
	mockRepo.AssertCalled(t, "SelectUser", userId)

	mockRepo.AssertNumberOfCalls(t, "InsertItem", 0)
	mockCacheRepo.AssertNumberOfCalls(t, "SetUser", 0)
}

func (s *TodoUsecaseTestSuite) TestGivenCannotInsertItemToDB_WhenAddingItem_ShouldReturnError() {
	userId := uint(1)
	t := s.T()
	mockCacheRepo := new(mocks.TodoCacheRepository)
	mockCacheRepo.On("GetCachedUser", uint(1)).Return(model.UserRedisModel{
		ID:               userId,
		CurrentUsage:     1,
		DailyRecordLimit: 5,
	}, nil)
	mockCacheRepo.On("SetUser", model.UserRedisModel{
		ID:               userId,
		CurrentUsage:     1,
		DailyRecordLimit: 5,
	}).Return(nil)

	mockRepo := new(mocks.TodoRepository)
	mockRepo.On("InsertItem", model.TodoItemModel{
		Content: "Hello",
		IsDone:  false,
		UserID:  userId,
	}).Return(errors.New("cannot insert item"))
	mockRepo.On("SelectUser", uint(1)).Return(model.UserModel{}, errors.New("user not found"))

	u := usecase.NewTodoUseCase(mockRepo, mockCacheRepo)
	err := u.AddTodo(dto.TodoDto{
		ID:      1,
		Content: "Hello",
		UserId:  1,
	})
	assert.Error(t, err)
	mockCacheRepo.AssertCalled(t, "GetCachedUser", userId)
	mockRepo.AssertNumberOfCalls(t, "SelectUser", 0)

	mockRepo.AssertNumberOfCalls(t, "InsertItem", 1)
	mockCacheRepo.AssertNumberOfCalls(t, "SetUser", 0)
}

func TestUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TodoUsecaseTestSuite))
}
