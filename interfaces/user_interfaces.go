package interfaces

import (
	"github.com/qgdomingo/todo-app/model"
)

type IUserRepository interface {
	LoginUserDB (user *model.UserLogin) (bool, map[string]string)
	RegisterUserDB (user *model.NewUser) (bool, map[string]string)
	FetchUserDetailsDB (username string) ([]model.UserDetails, map[string]string)
	UpdateUserDetailsDB (user *model.UserDetails, username string) (bool, map[string]string)
	UserPasswordChangeDB (user *model.UserNewPassword, username string) (bool, map[string]string)
}