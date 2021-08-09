package repo

import (
	"github.com/manabie-com/togo/common/context"
	"github.com/manabie-com/togo/domain/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) GetUserByUsername(ctx context.Context, userName string) (*model.User, error) {
	args := u.Called(ctx, userName)
	return args.Get(0).(*model.User), args.Error(1)
}

func (u *UserRepositoryMock) GetUserById(ctx context.Context, id string) (*model.User, error) {
	args := u.Called(ctx, id)
	return args.Get(0).(*model.User), args.Error(1)
}
