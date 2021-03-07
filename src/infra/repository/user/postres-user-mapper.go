package user

import "togo/src/entity/user"

type UserMapper struct {
}

func (um *UserMapper) ToDatabase(user *user.User) interface{} {
	return nil
}

func (um *UserMapper) ToEntity(data interface{}) (*user.User, error) {
	return nil, nil
}
