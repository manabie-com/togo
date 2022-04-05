package mocks

import (
	context "context"
	"togo/internal/pkg/domain/dtos"
	"togo/internal/pkg/domain/entities"

	"github.com/stretchr/testify/mock"
)

type TodoUsecase struct {
	mock.Mock
}

// NewMockTodoUsecase creates a new mock instance.
func NewMockTodoUsecase() *TodoUsecase {
	return new(TodoUsecase)
}

// Create func mock
func (u *TodoUsecase) Create(ctx context.Context, req dtos.CreateTodoRequest, user entities.User) error {
	ret := u.Called(ctx, req, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(ctx context.Context, req dtos.CreateTodoRequest, user entities.User) error); ok {
		r0 = rf(ctx, req, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
