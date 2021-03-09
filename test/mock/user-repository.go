package mock

import (
	"togo/src/entity/user"
)

type UserRepositoryMock struct {
	GetFindOneFunc func(options interface{}) (*user.User, error)
}

func (this *UserRepositoryMock) FindOne(options interface{}) (*user.User, error) {
	return this.GetFindOneFunc(options)
}

func New_UserRepository_With_FindOneOK() *UserRepositoryMock {
	return &UserRepositoryMock{
		GetFindOneFunc: func(options interface{}) (*user.User, error) {
			return &user.User{
				ID:      "firstUser",
				MaxTodo: 5,
			}, nil
		},
	}
}

func New_UserRepository_With_FindOneNotFound() *UserRepositoryMock {
	return &UserRepositoryMock{
		GetFindOneFunc: func(options interface{}) (*user.User, error) {
			return nil, ERROR_404
		},
	}
}
