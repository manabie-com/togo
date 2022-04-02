package services

import (
	"context"
	"togo/internal/domain"
	"togo/internal/provider"
	"togo/internal/repository"

	"github.com/stretchr/testify/mock"
)

type mockPasswordHashProvider struct {
	mock.Mock
}

var _ provider.PasswordHashProvider = new(mockPasswordHashProvider)

func (p *mockPasswordHashProvider) HashPassword(password string) (string, error) {
	args := p.Called(password)
	return args.String(0), args.Error(1)
}

func (p *mockPasswordHashProvider) ComparePassword(password, hashPassword string) error {
	args := p.Called(password, hashPassword)
	return args.Error(0)
}

type mockTokenProvider struct {
	mock.Mock
}

var _ provider.TokenProvider = new(mockTokenProvider)

func (p *mockTokenProvider) GenerateToken(data interface{}) (token string, err error) {
	args := p.Called(data)
	return args.String(0), args.Error(1)
}

func (p *mockTokenProvider) VerifyToken(token string) (payload interface{}, err error) {
	args := p.Called(token)
	return args.Get(0), args.Error(1)
}

type mockUserRepository struct {
	mock.Mock
}

var _ repository.UserRepository = new(mockUserRepository)

func (r *mockUserRepository) Create(ctx context.Context, entity *domain.User) (*domain.User, error) {
	args := r.Called(entity)
	if u := args.Get(0); u != nil {
		return u.(*domain.User), nil
	}
	return nil, args.Error(1)
}

func (r *mockUserRepository) FindOne(ctx context.Context, filter *domain.User) (*domain.User, error) {
	args := r.Called(filter)
	if u := args.Get(0); u != nil {
		return u.(*domain.User), nil
	}
	return nil, args.Error(1)
}

func (r *mockUserRepository) UpdateByID(ctx context.Context, id uint, update *domain.User) (*domain.User, error) {
	args := r.Called(id, update)
	if u := args.Get(0); u != nil {
		return u.(*domain.User), nil
	}
	return nil, args.Error(1)
}

type mockTaskRepository struct {
	mock.Mock
}

var _ repository.TaskRepository = new(mockTaskRepository)

func (r *mockTaskRepository) Create(ctx context.Context, entity *domain.Task) (*domain.Task, error) {
	args := r.Called(entity)
	if u := args.Get(0); u != nil {
		return u.(*domain.Task), nil
	}
	return nil, args.Error(1)
}

func (r *mockTaskRepository) Update(ctx context.Context, filter, update *domain.Task) (*domain.Task, error) {
	args := r.Called(filter, update)
	if u := args.Get(0); u != nil {
		return u.(*domain.Task), nil
	}
	return nil, args.Error(1)
}

func (r *mockTaskRepository) Find(ctx context.Context, filter *domain.Task) ([]*domain.Task, error) {
	args := r.Called(filter)
	if u := args.Get(0); u != nil {
		return u.([]*domain.Task), nil
	}
	return nil, args.Error(1)
}

type mockTaskLimitRepository struct {
	mock.Mock
}

var _ repository.TaskLimitRepository = new(mockTaskLimitRepository)

func (r mockTaskLimitRepository) Increase(ctx context.Context, userID uint, limit int) (int, error) {
	args := r.Called(userID, limit)
	return args.Int(0), args.Error(1)
}
