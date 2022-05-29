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

type taskCreateInput struct {
	Title       string
	Description string
}

type taskCreateOutput struct {
	Error error
}

type taskCreateTestCase struct {
	input  taskCreateInput
	output taskCreateOutput
}

func TestTaskServiceImpl_CreateTask(t *testing.T) {
	repo := inmemory.NewInMemoryTaskRepository()
	taskSvc := service.NewTaskService(repo)
	testcases := []taskCreateTestCase{
		{
			input: taskCreateInput{
				Title:       "Task1",
				Description: "Task1",
			},
			output: taskCreateOutput{
				Error: nil,
			},
		},
		{
			input: taskCreateInput{
				Title:       "Task2",
				Description: "Task1",
			},
			output: taskCreateOutput{
				Error: nil,
			},
		},
		{
			input: taskCreateInput{
				Title:       "Task3",
				Description: "Task1",
			},
			output: taskCreateOutput{
				Error: nil,
			},
		},
		{
			input: taskCreateInput{
				Title:       "Task4",
				Description: "Task1",
			},
			output: taskCreateOutput{
				Error: nil,
			},
		},

		{
			input: taskCreateInput{
				Title:       "Task1",
				Description: "Task1",
			},
			output: taskCreateOutput{
				Error: errdef.DupplicateTask,
			},
		},
		{
			input: taskCreateInput{
				Title:       "Task5",
				Description: "Task1",
			},
			output: taskCreateOutput{
				Error: nil,
			},
		},
		{
			input: taskCreateInput{
				Title:       "Task7",
				Description: "Task1",
			},
			output: taskCreateOutput{
				Error: errdef.LimitTaskCreated,
			},
		},
	}

	testUser := model.User{
		Id:       1,
		Username: "test",
		Password: "test",
		Token:    "",
		Limit:    5,
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("TestTaskServiceImpl_CreateTask [%d]", i), func(t *testing.T) {
			_, err := taskSvc.CreateTask(context.Background(), &testUser, &model.Task{Title: tc.input.Title, Description: tc.input.Description})
			if err != tc.output.Error {
				t.Errorf("Error TestTaskServiceImpl_CreateTask [%d] got %#v want %#v", i, err, tc.output.Error)
			}
		})
	}

}
