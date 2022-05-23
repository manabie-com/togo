package repository

import "togo/domain/model"

type UserRepository interface {
	Create(u model.User) error
	Get(username string) (u model.User, err error)
}

type TaskRepository interface {
	Create(u model.Task) error
}

type TokenResponse interface {
}
