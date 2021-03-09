package test

import (
	"testing"
	"togo/src"
	"togo/src/domain/user"
	"togo/src/schema"
	"togo/test/mock"

	"github.com/stretchr/testify/assert"
)

type LoginInputMock struct {
	userRepositoryMock *mock.UserRepositoryMock
	jwtService         src.IJWTService
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		testName       string
		input          *LoginInputMock
		expectedOutput *schema.LoginResponse
		expectedError  error
	}{
		{
			"Login success",
			&LoginInputMock{
				userRepositoryMock: mock.New_UserRepository_With_FindOneOK(),
				jwtService:         mock.New_JwtMock_With_CreateTokenOK(),
			},
			&schema.LoginResponse{
				UserId: "firstUser",
				Token:  mock.TOKEN,
			},
			nil,
		},

		{
			"UserID is not exists in database",
			&LoginInputMock{
				userRepositoryMock: mock.New_UserRepository_With_FindOneNotFound(),
				jwtService:         mock.New_JwtMock_With_CreateTokenOK(),
			},
			nil,
			mock.ERROR_404,
		},

		{
			"Create token error",
			&LoginInputMock{
				userRepositoryMock: mock.New_UserRepository_With_FindOneOK(),
				jwtService:         mock.New_JwtMock_With_CreateTokenError(),
			},
			nil,
			mock.ERROR_500,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.testName, func(t *testing.T) {
			var userWorkflow user.IUserWorkflow = &user.UserWorkflow{
				Repository: testCase.input.userRepositoryMock,
				JWTService: testCase.input.jwtService,
			}

			actualOutput, actualError := userWorkflow.Login(&schema.LoginRequest{})

			// t.Logf("\n[expected-output]: %v %v\n[actual-output]: %v %v", testCase.expectedOutput, testCase.expectedError, actualOutput, actualError)

			assert.Equal(t, testCase.expectedOutput, actualOutput)
			assert.Equal(t, testCase.expectedError, actualError)
		})
	}
}
