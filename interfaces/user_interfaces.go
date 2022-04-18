package interfaces

import (
	"github.com/qgdomingo/todo-app/model"
)

// These are the function signatures implemented by user_repo.go. 
// 		These are interfaced to allow the implementation of the user_repo_mock.go 
//		for the unit testing of the user_controller.go
type IUserRepository interface {
	LoginUserDB (user *model.UserLogin) (bool, map[string]string)
	RegisterUserDB (user *model.NewUser) (bool, map[string]string)
	FetchUserDetailsDB (username string) ([]model.UserDetails, map[string]string)
	UpdateUserDetailsDB (user *model.UserDetails, username string) (bool, map[string]string)
	UserPasswordChangeDB (user *model.UserNewPassword, username string) (bool, map[string]string)
}