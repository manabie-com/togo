package usecases

import (
	"context"
	"errors"
	"testing"
	"togo/internal/pkg/domain/dtos"
	"togo/internal/pkg/domain/entities"
	"togo/internal/pkg/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var ctx = context.Background()

func TestUserRepoCreate(t *testing.T) {
	dueDate := "2022-04-04 15:00"
	req := dtos.CreateTodoRequest{
		Task:    "task1",
		DueDate: &dueDate,
	}
	user := entities.User{
		ID:        1,
		Email:     "test@gmail.com",
		LimitTodo: 5,
	}
	repoMock := mocks.NewMockTodoRepo()
	u := NewTodoUsecase(repoMock)

	t.Run("success", func(t *testing.T) {
		repoMock.On("CountTodosByDay", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		repoMock.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		err := u.Create(ctx, req, user)

		assert.NoError(t, err)
	})

	t.Run("test fail limit todo", func(t *testing.T) {
		repoMock.On("CountTodosByDay", mock.Anything, mock.AnythingOfType("int")).Return(5, nil).Once()

		err := u.Create(ctx, req, user)

		expectedError := errors.New("LIMIT_DAILY_TASK")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
	})
}
