package usecase

import (
	"errors"
	"github.com/stretchr/testify/require"
	"manabieAssignment/internal/core/entity"
	"manabieAssignment/internal/mocks"
	"testing"
	"time"
)

func TestTodoUseCase_CreateTodo_Success(t *testing.T) {
	todoEntity := entity.Todo{
		UserID:    1,
		Name:      "todo_name",
		Content:   "todo_content",
		CreatedAt: time.Now(),
	}

	mockTodoRepository := new(mocks.TodoRepository)
	mockUserRepository := new(mocks.UserRepository)

	mockUserRepository.On("IsUserExisted", todoEntity.UserID).Return(nil)
	mockUserRepository.On("IsUserHavingMaxTodo", todoEntity.UserID, todoEntity.CreatedAt).Return(nil)
	mockTodoRepository.On("CreateTodo", todoEntity).Return(uint(1), nil)

	todoUC := NewTodoUseCase(mockTodoRepository, mockUserRepository)
	todoId, err := todoUC.CreateTodo(todoEntity)
	require.NoError(t, err)
	require.Equal(t, uint(1), todoId)
}

func TestTodoUseCase_CreateTodo_UserIdIsInvalid(t *testing.T) {
	todoEntity := entity.Todo{
		UserID:    0,
		Name:      "todo_name",
		Content:   "todo_content",
		CreatedAt: time.Now(),
	}

	mockTodoRepository := new(mocks.TodoRepository)
	mockUserRepository := new(mocks.UserRepository)
	todoEntity.UserID = 0
	todoUC := NewTodoUseCase(mockTodoRepository, mockUserRepository)
	_, err := todoUC.CreateTodo(todoEntity)
	require.Error(t, err)
	require.Equal(t, errors.New("invalid userId"), err)
}

func TestTodoUseCase_CreateTodo_NameIsEmpty(t *testing.T) {
	todoEntity := entity.Todo{
		UserID:    1,
		Name:      "",
		Content:   "todo_content",
		CreatedAt: time.Now(),
	}

	mockTodoRepository := new(mocks.TodoRepository)
	mockUserRepository := new(mocks.UserRepository)
	todoEntity.Name = ""
	todoUC := NewTodoUseCase(mockTodoRepository, mockUserRepository)
	_, err := todoUC.CreateTodo(todoEntity)
	require.Error(t, err)
	require.Equal(t, errors.New("name is empty"), err)
}

func TestTodoUseCase_CreateTodo_UserIsNotExisted(t *testing.T) {
	todoEntity := entity.Todo{
		UserID:    1,
		Name:      "todo_name",
		Content:   "todo_content",
		CreatedAt: time.Now(),
	}

	mockTodoRepository := new(mocks.TodoRepository)
	mockUserRepository := new(mocks.UserRepository)

	mockUserRepository.On("IsUserExisted", todoEntity.UserID).Return(errors.New("user is not existed"))

	todoUC := NewTodoUseCase(mockTodoRepository, mockUserRepository)
	_, err := todoUC.CreateTodo(todoEntity)
	require.Error(t, err)
	require.Equal(t, errors.New("user is not existed"), err)
}

func TestTodoUseCase_CreateTodo_UserHaveTooManyTodos(t *testing.T) {
	todoEntity := entity.Todo{
		UserID:    1,
		Name:      "todo_name",
		Content:   "todo_content",
		CreatedAt: time.Now(),
	}

	mockTodoRepository := new(mocks.TodoRepository)
	mockUserRepository := new(mocks.UserRepository)

	mockUserRepository.On("IsUserExisted", todoEntity.UserID).Return(nil)
	mockUserRepository.On("IsUserHavingMaxTodo", todoEntity.UserID, todoEntity.CreatedAt).Return(errors.New("user have too many todos"))

	todoUC := NewTodoUseCase(mockTodoRepository, mockUserRepository)
	_, err := todoUC.CreateTodo(todoEntity)
	require.Error(t, err)
	require.Equal(t, errors.New("user have too many todos"), err)
}

func TestTodoUseCase_CreateTodo_CreateTodoFail(t *testing.T) {
	todoEntity := entity.Todo{
		UserID:    1,
		Name:      "todo_name",
		Content:   "todo_content",
		CreatedAt: time.Now(),
	}

	mockTodoRepository := new(mocks.TodoRepository)
	mockUserRepository := new(mocks.UserRepository)

	mockUserRepository.On("IsUserExisted", todoEntity.UserID).Return(nil)
	mockUserRepository.On("IsUserHavingMaxTodo", todoEntity.UserID, todoEntity.CreatedAt).Return(errors.New("user have too many todos"))
	mockTodoRepository.On("CreateTodo", todoEntity).Return(0, errors.New("some errors"))

	todoUC := NewTodoUseCase(mockTodoRepository, mockUserRepository)
	_, err := todoUC.CreateTodo(todoEntity)
	require.Error(t, err)
}
