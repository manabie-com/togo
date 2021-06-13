package mocks

import (
	"context"
	"github.com/manabie-com/togo/internal/storages/ent"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) FindByUsername(ctx context.Context, username string) (*ent.User, error) {
	ret := m.Called(ctx, username)
	return ret.Get(0).(*ent.User), ret.Error(1)

}

func (m *UserRepositoryMock) FindByUserId(ctx context.Context, userId string) (*ent.User, error) {
	ret := m.Called(ctx, userId)
	return ret.Get(0).(*ent.User), ret.Error(1)

}

func (m *UserRepositoryMock) CreateUser(ctx context.Context, username string, pwd string) (*ent.User, error) {
	ret := m.Called(ctx, username, pwd)
	return ret.Get(0).(*ent.User), ret.Error(1)

}

func (m *UserRepositoryMock) InitDb() {
}
