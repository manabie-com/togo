package services

import (
	"context"
	"togo/internal/domain"

	"github.com/stretchr/testify/mock"
)

type mockPasswordHashProvider struct {
	mock.Mock
}

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
