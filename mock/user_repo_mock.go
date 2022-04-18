package mock

import (
	"github.com/qgdomingo/todo-app/model"
)

// This is the dummy repository implementation of the User functions in user_repo.go
// 		which will be used in the unit test in user_controller_test.go

type UserRepositoryMock struct {
	UserList []model.UserDetails
	IsTaskSuccessful bool
	ErrorMessage map[string]string
}


func (u *UserRepositoryMock) LoginUserDB (user *model.UserLogin) (bool, map[string]string) {
	return u.IsTaskSuccessful, u.ErrorMessage
}

func (u *UserRepositoryMock) RegisterUserDB (user *model.NewUser) (bool, map[string]string) {
	return u.IsTaskSuccessful, u.ErrorMessage
}

func (u *UserRepositoryMock) FetchUserDetailsDB (username string) ([]model.UserDetails, map[string]string) {
	return u.UserList, u.ErrorMessage
}

func (u *UserRepositoryMock) UpdateUserDetailsDB (user *model.UserDetails, username string) (bool, map[string]string) {
	return u.IsTaskSuccessful, u.ErrorMessage
}

func (u *UserRepositoryMock) UserPasswordChangeDB (user *model.UserNewPassword, username string) (bool, map[string]string) {
	return u.IsTaskSuccessful, u.ErrorMessage
}