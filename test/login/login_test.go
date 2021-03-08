package login_test

import (
	"errors"
	"testing"
	"togo/src"
	"togo/src/entity/task"
	"togo/src/entity/user"
	"togo/src/schema"

	uWorkflow "togo/src/domain/user"

	"github.com/stretchr/testify/assert"
)

type (
	TaskRepositoryMock struct {
		GetCreateFunc func(data *task.Task) (*task.Task, error)
		GetCountFunc  func(filter interface{}) (int, error)
	}
	UserRepositoryMock struct {
		GetFindOneFunc func(options interface{}) (*user.User, error)
	}
	ContextMock struct {
		GetTokenDataFunc func() *src.TokenData
	}
	JwtMock struct {
		GetCreateTokenFunc func(data *src.TokenData) (string, error)
	}
)

func (this *JwtMock) CreateToken(data *src.TokenData) (string, error) {
	return this.GetCreateTokenFunc(data)
}

func (this *JwtMock) VerifyToken(token string) (*src.TokenData, error) {
	return nil, nil
}

func (csm *ContextMock) GetTokenData() *src.TokenData {
	return csm.GetTokenDataFunc()
}

func (csm *ContextMock) CheckPermission(scopes []string) error {
	return nil
}

func (csm *ContextMock) LoadContext(data interface{}) error {
	return nil
}

func (um *UserRepositoryMock) FindOne(options interface{}) (*user.User, error) {
	return um.GetFindOneFunc(options)
}

func (tm *TaskRepositoryMock) Create(data *task.Task) (*task.Task, error) {
	return tm.GetCreateFunc(data)
}

func (tm *TaskRepositoryMock) Count(filter interface{}) (int, error) {
	return tm.GetCountFunc(filter)
}

type LoginInputMock struct {
	userRepoMock *UserRepositoryMock
	loginRequest *schema.LoginRequest
	jwtService   src.IJWTService
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		input          *LoginInputMock
		expectedOutput *schema.LoginResponse
		expectedError  error
	}{
		// fetch first user - everything should happens without errors
		{
			&LoginInputMock{
				nil,
				&schema.LoginRequest{
					UserId: "firstUser",
				},
				&JwtMock{
					GetCreateTokenFunc: func(data *src.TokenData) (string, error) {
						return "this-is-secret-token", nil
					},
				},
			},
			&schema.LoginResponse{
				UserId: "firstUser",
				Token:  "this-is-secret-token",
			},
			nil,
		},

		// fetch non-exists user - should return error
		{
			&LoginInputMock{
				nil,
				&schema.LoginRequest{
					UserId: "secondUser",
				},
				nil,
			},
			nil,
			errors.New("USER_NOT_FOUND"),
		},

		// UserId is not provided - should return error
		{
			&LoginInputMock{
				nil,
				&schema.LoginRequest{
					UserId: "",
				},
				nil,
			},
			nil,
			errors.New("USER_NOT_FOUND"),
		},
	}

	for _, testCase := range testCases {
		testCase.input.userRepoMock = &UserRepositoryMock{
			GetFindOneFunc: func(options interface{}) (*user.User, error) {
				filter := options.(*user.User)
				if filter.ID == "firstUser" {
					return &user.User{
						ID:       filter.ID,
						Password: "example",
					}, nil
				}
				return nil, errors.New("USER_NOT_FOUND")
			},
		}

		var userWorkflow uWorkflow.IUserWorkflow = &uWorkflow.UserWorkflow{
			Repository: testCase.input.userRepoMock,
			JWTService: testCase.input.jwtService,
		}

		actualOutput, actualError := userWorkflow.Login(testCase.input.loginRequest)

		t.Logf("\n[expected-output]: %v %v\n[actual-output]: %v %v", testCase.expectedOutput, testCase.expectedError, actualOutput, actualError)

		assert.Equal(t, testCase.expectedOutput, actualOutput)
		assert.Equal(t, testCase.expectedError, actualError)
	}
}
