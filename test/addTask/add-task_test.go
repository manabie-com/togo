package addTask_test

import (
	"errors"
	"testing"
	"togo/src"
	"togo/src/entity/task"
	"togo/src/entity/user"
	"togo/src/schema"

	"github.com/stretchr/testify/assert"

	taskWork "togo/src/domain/task"
	gErrors "togo/src/infra/error"
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
)

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

type InputMock struct {
	mockTaskRepo *TaskRepositoryMock
	mockUserRepo *UserRepositoryMock
	mockContext  *ContextMock
}

func TestAddTaskByOwner(t *testing.T) {
	testCases := []struct {
		input          *InputMock
		expectedOutput *schema.AddTaskResponse
		expectedError  error
	}{
		// Normal creating - everything should happens without errors
		{
			&InputMock{
				&TaskRepositoryMock{
					GetCreateFunc: func(data *task.Task) (*task.Task, error) {
						return &task.Task{Id: "89f403d9-a48c-4c14-9bc0-fe170d4ae30f"}, nil
					},
					GetCountFunc: func(filter interface{}) (int, error) {
						return 4, nil
					},
				},
				nil,
				&ContextMock{
					GetTokenDataFunc: func() *src.TokenData {
						return &src.TokenData{
							UserId: "firstUser",
						}
					},
				},
			},
			&schema.AddTaskResponse{
				TaskId: "89f403d9-a48c-4c14-9bc0-fe170d4ae30f",
			},
			nil,
		},

		/*
			No permission context - this would never happens because the permission is checked in upper layer
			(btw this should be the list in test case for other unexpected thing happens)
		*/
		{
			&InputMock{
				&TaskRepositoryMock{},
				nil,
				&ContextMock{
					GetTokenDataFunc: func() *src.TokenData {
						return &src.TokenData{
							UserId: "secondUser",
						}
					},
				},
			},
			nil,
			errors.New("USER_NOT_FOUND"),
		},

		// User have number task per day reach limit - should return error
		{
			&InputMock{
				&TaskRepositoryMock{
					GetCountFunc: func(filter interface{}) (int, error) {
						return 5, nil
					},
				},
				nil,
				&ContextMock{
					GetTokenDataFunc: func() *src.TokenData {
						return &src.TokenData{
							UserId: "firstUser",
						}
					},
				},
			},
			nil,
			gErrors.NewBadRequestError(src.MAX_TODO_OVER_LIMIT, errors.New("max todo over limit")),
		},

		// The number task per day over the limit - should return error
		{
			&InputMock{
				&TaskRepositoryMock{
					GetCountFunc: func(filter interface{}) (int, error) {
						return 6, nil
					},
				},
				nil,
				&ContextMock{
					GetTokenDataFunc: func() *src.TokenData {
						return &src.TokenData{
							UserId: "firstUser",
						}
					},
				},
			},
			nil,
			gErrors.NewBadRequestError(src.MAX_TODO_OVER_LIMIT, errors.New("max todo over limit")),
		},

		// UserId not defined in context (by some unexpected way) - should return error
		{
			&InputMock{
				&TaskRepositoryMock{},
				nil,
				&ContextMock{
					GetTokenDataFunc: func() *src.TokenData {
						return &src.TokenData{
							UserId: "",
						}
					},
				},
			},
			nil,
			errors.New("USER_NOT_FOUND"),
		},
	}

	for _, testCase := range testCases {
		testCase.input.mockUserRepo = &UserRepositoryMock{
			GetFindOneFunc: func(options interface{}) (*user.User, error) {
				filter := options.(user.User)
				if filter.ID == "firstUser" {
					return &user.User{ID: "firstUser", MaxTodo: 5}, nil
				}
				return nil, errors.New("USER_NOT_FOUND")
			},
		}

		var context src.IContextService = testCase.input.mockContext

		var taskWorkflow taskWork.ITaskWorkflow = &taskWork.TaskWorkflow{
			Repository:     testCase.input.mockTaskRepo,
			UserRepository: testCase.input.mockUserRepo,
			ErrorFactory:   gErrors.NewErrorFactory(),
		}

		actualOutput, actualError := taskWorkflow.AddTaskByOwner(context, &schema.AddTaskRequest{
			Content: "Hello World",
		})

		t.Logf("\n[expected-output]: %v %v\n[actual-output]: %v %v", testCase.expectedOutput, testCase.expectedError, actualOutput, actualError)

		assert.Equal(t, testCase.expectedOutput, actualOutput)
		assert.Equal(t, testCase.expectedError, actualError)
	}
}
