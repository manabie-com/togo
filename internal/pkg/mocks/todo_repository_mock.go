package mocks

import (
	context "context"
	"togo/internal/pkg/domain/entities"

	"github.com/stretchr/testify/mock"
)

type TodoRepository struct {
	mock.Mock
}

// NewMockTodoRepo creates a new mock instance.
func NewMockTodoRepo() *TodoRepository {
	return new(TodoRepository)
}

// Create func mock
func (r *TodoRepository) Create(ctx context.Context, todo entities.Todo) error {
	ret := r.Called(ctx, todo)

	var r0 error
	if rf, ok := ret.Get(0).(func(ctx context.Context, todo entities.Todo) error); ok {
		r0 = rf(ctx, todo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountTodosByDay func mock
func (r *TodoRepository) CountTodosByDay(ctx context.Context, userID int) (int, error) {
	ret := r.Called(ctx, userID)

	var r0 int
	if rf, ok := ret.Get(0).(func(ctx context.Context, userID int) int); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctx context.Context, userID int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}
