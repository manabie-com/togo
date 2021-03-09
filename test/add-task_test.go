package test

import (
	"testing"
	"togo/src"
	"togo/src/domain/task"
	"togo/src/schema"
	"togo/test/mock"

	"github.com/stretchr/testify/assert"
)

type InputMock struct {
	taskRepositoryMock *mock.TaskRepositoryMock
	userRepositoryMock *mock.UserRepositoryMock
	contextMock        *mock.ContextMock
}

func TestAddTaskByOwner(t *testing.T) {
	testCases := []struct {
		testName       string
		input          *InputMock
		expectedOutput *schema.AddTaskResponse
		expectedError  error
	}{
		{
			"Create task success",
			&InputMock{
				taskRepositoryMock: mock.New_TaskRepository_CreateOK_CountLessThan5(),
				userRepositoryMock: mock.New_UserRepository_With_FindOneOK(),
				contextMock:        mock.New_ContextMock_With_Valid_UserID(),
			},
			&schema.AddTaskResponse{TaskId: "89f403d9-a48c-4c14-9bc0-fe170d4ae30f"},
			nil,
		},

		{
			"Cannot create task because UserId is not exists in database",
			&InputMock{
				taskRepositoryMock: mock.New_TaskRepository_CreateOK_CountLessThan5(),
				userRepositoryMock: mock.New_UserRepository_With_FindOneNotFound(),
				contextMock:        mock.New_ContextMock_With_Valid_UserID(),
			},
			nil,
			mock.ERROR_404,
		},

		{
			"Cannot create tasl because user's task-number reach limit",
			&InputMock{
				taskRepositoryMock: mock.New_TaskRepository_With_CreateOK_CountEqual5(),
				userRepositoryMock: mock.New_UserRepository_With_FindOneOK(),
				contextMock:        mock.New_ContextMock_With_Valid_UserID(),
			},
			nil,
			mock.ERROR_400,
		},

		{
			"Cannot create tasl because user's task-number larger than limit",
			&InputMock{
				taskRepositoryMock: mock.New_TaskRepository_With_CreateOK_CountLargerThan5(),
				userRepositoryMock: mock.New_UserRepository_With_FindOneOK(),
				contextMock:        mock.New_ContextMock_With_Valid_UserID(),
			},
			nil,
			mock.ERROR_400,
		},

		{
			"Cannot create tasl because internal error",
			&InputMock{
				taskRepositoryMock: mock.New_TaskRepository_With_TaskCount_Error(),
				userRepositoryMock: mock.New_UserRepository_With_FindOneOK(),
				contextMock:        mock.New_ContextMock_With_Valid_UserID(),
			},
			nil,
			mock.ERROR_500,
		},
	}

	var errorFactoryMock src.IErrorFactory = mock.NewErrorFactoryMock()

	for _, testCase := range testCases {

		t.Run(testCase.testName, func(t *testing.T) {
			var context src.IContextService = testCase.input.contextMock

			var taskWorkflow task.ITaskWorkflow = &task.TaskWorkflow{
				Repository:     testCase.input.taskRepositoryMock,
				UserRepository: testCase.input.userRepositoryMock,
				ErrorFactory:   errorFactoryMock,
			}

			actualOutput, actualError := taskWorkflow.AddTaskByOwner(context, &schema.AddTaskRequest{
				Content: "Hello World",
			})

			// t.Logf("\n[expected-output]: %v %v\n[actual-output]: %v %v", testCase.expectedOutput, testCase.expectedError, actualOutput, actualError)

			assert.Equal(t, testCase.expectedOutput, actualOutput)
			assert.Equal(t, testCase.expectedError, actualError)
		})
	}
}
