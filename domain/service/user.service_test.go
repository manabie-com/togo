package service_test

import (
	"context"
	"fmt"
	"testing"
	"togo/domain/errdef"
	"togo/domain/model"
	"togo/domain/service"
	"togo/infrastructure/inmemory"
)

type loginInput struct {
	username string
	password string
}

type loginOutput struct {
	Error error
}

type loginTestCase struct {
	Input  loginInput
	Output loginOutput
}

func TestUserService_Login(t *testing.T) {
	repo := inmemory.NewInMemoryUserRepo()
	tokenService := service.NewTokenService(secret)
	userService := service.NewUserService(repo, tokenService)
	// Mock user
	userService.Register(context.Background(), model.User{
		Username: "admin",
		Password: "admin",
	})

	testcases := []loginTestCase{
		{
			Input: loginInput{
				username: "admin",
				password: "admin",
			},
			Output: loginOutput{
				Error: nil,
			},
		},
		{
			Input: loginInput{
				username: "admin",
				password: "admin2",
			},
			Output: loginOutput{
				Error: errdef.InvalidUsernameOrPassword,
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("TestUserService_Login [%d]", i), func(t *testing.T) {
			_, err := userService.Login(context.Background(), tc.Input.username, tc.Input.password)
			if err == tc.Output.Error {
				t.Log("Successfully")
			} else {
				t.Errorf("Error got: %#v - want: %#v", err, tc.Output.Error)
			}
		})

	}
}
