// Package createtodo implements create todo feature
package createtodo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Set of error variables for the feature.
var (
	ErrExceededDailyMaximumTodos = errors.New("user has exceeded maximum todos allowed for today")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	WithinTran(ctx context.Context, fn func(s Storer) error) error
	GetUserDailyMaxTodo(ctx context.Context, userID uuid.UUID) (int, error)
	GetUserTodayTodoCount(ctx context.Context, userID uuid.UUID) (int, error)
	CreateTodo(ctx context.Context, todo Todo) error
}

// Feature manages the api for create todo.
type Feature struct {
	storer Storer
}

// NewFeature constructs a feature for create todo api access.
func NewFeature(storer Storer) *Feature {
	return &Feature{
		storer: storer,
	}
}

// Create inserts a new user into the database.
func (f *Feature) CreateTodo(ctx context.Context, nt NewTodo) (Todo, error) {
	now := time.Now()

	todo := Todo{
		ID:          uuid.New(),
		Title:       nt.Title,
		Content:     nt.Content,
		UserID:      nt.UserID,
		DateCreated: now,
		DateUpdated: now,
	}

	tran := func(s Storer) error {
		maxTodo, err := s.GetUserDailyMaxTodo(ctx, todo.UserID)
		if err != nil {
			return fmt.Errorf("getuserdailymaxtodo: %w", err)
		}

		todoCount, err := s.GetUserTodayTodoCount(ctx, todo.UserID)
		if err != nil {
			return fmt.Errorf("getusertodaytodocount: %w", err)
		}

		if todoCount >= maxTodo {
			return ErrExceededDailyMaximumTodos
		}

		if err := s.CreateTodo(ctx, todo); err != nil {
			return fmt.Errorf("createtodo: %w", err)
		}

		return nil
	}

	if err := f.storer.WithinTran(ctx, tran); err != nil {
		return Todo{}, fmt.Errorf("tran: %w", err)
	}

	return todo, nil
}
