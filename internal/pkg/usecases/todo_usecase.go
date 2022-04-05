package usecases

import (
	"context"
	"fmt"
	"time"
	"togo/internal/pkg/domain/dtos"
	"togo/internal/pkg/domain/entities"
	"togo/internal/pkg/repositories"
)

type TodoUsecase interface {
	Create(ctx context.Context, req dtos.CreateTodoRequest, user entities.User) error
}

type todoUsecase struct {
	repo repositories.TodoRepository
}

func (u *todoUsecase) Create(ctx context.Context, req dtos.CreateTodoRequest, user entities.User) error {
	count, err := u.repo.CountTodosByDay(ctx, user.ID)
	if err != nil {
		return err
	}

	fmt.Println(count)
	if count >= user.LimitTodo {
		return fmt.Errorf("LIMIT_DAILY_TASK")
	}
	var dueDate time.Time
	if req.DueDate != nil {
		dueDate, err = time.Parse("2006-01-02 15:04", *req.DueDate)
		if err != nil {
			return err
		}
	}

	todo := entities.Todo{
		UserID:  user.ID,
		Task:    req.Task,
		DueDate: &dueDate,
	}

	return u.repo.Create(ctx, todo)
}

// NewTodoUsecase
func NewTodoUsecase(repo repositories.TodoRepository) TodoUsecase {
	return &todoUsecase{
		repo: repo,
	}
}
