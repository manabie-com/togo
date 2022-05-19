package repositories

import (
	"manabie.com/internal/models"
	"manabie.com/internal/common"
	"context"
)

type UserRepositoryMock struct {
	Users map[int]models.User
}

func MakeUserRepositoryMock() UserRepositoryMock {
	return UserRepositoryMock{
		Users: map[int]models.User{},
	}
}

func (r *UserRepositoryMock) AddUser(iUser models.User) {
	r.Users[iUser.Id] = iUser
}

func (r *UserRepositoryMock) FetchUserById(iContext context.Context, iId int) (models.User, error) {
	if user, ok := r.Users[iId]; !ok {
		return models.User{}, common.NotFound
	} else {
		return user, nil
	}
}
