package task_test

import (
	"errors"
	"testing"
	"togo/src"
	"togo/src/entity/task"
	"togo/src/entity/user"
	"togo/src/schema"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	taskWork "togo/src/domain/task"
)

type TaskRepositoryMock struct {
	mock.Mock
}

type UserRepositoryMock struct {
	mock.Mock
}

type ContextServiceMock struct {
	mock.Mock
}

func (csm *ContextServiceMock) GetTokenData() *src.TokenData {
	return &src.TokenData{
		UserId: "firstUser",
	}
}

func (csm *ContextServiceMock) CheckPermission(scopes []string) error {
	return nil
}

func (csm *ContextServiceMock) LoadContext(data interface{}) error {
	return nil
}

func (um *UserRepositoryMock) FindOne(options interface{}) (*user.User, error) {
	filter := options.(user.User)

	if filter.ID == "firstUser" {
		return &user.User{
			ID: "firstUser",
		}, nil
	}

	return nil, errors.New("Error")
}

func (tm *TaskRepositoryMock) Create(data *task.Task) (*task.Task, error) {
	task := &task.Task{
		Id: "89f403d9-a48c-4c14-9bc0-fe170d4ae30f",
	}
	return task, nil
}

func TestAddTaskByOwner(t *testing.T) {
	var workflow taskWork.ITaskWorkflow = &taskWork.TaskWorkflow{
		&TaskRepositoryMock{},
		&UserRepositoryMock{},
	}

	var contextService src.IContextService = &ContextServiceMock{}

	data, _ := workflow.AddTaskByOwner(contextService, &schema.AddTaskRequest{
		Content: "Hello World",
	})

	assert.Equal(t, &schema.AddTaskResponse{
		TaskId: "89f403d9-a48c-4c14-9bc0-fe170d4ae30f",
	}, data)
}
