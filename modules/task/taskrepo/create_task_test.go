package taskrepo_test

import (
	"context"
	"errors"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/modules/task/taskmodel"
	"github.com/japananh/togo/modules/task/taskrepo"
	"github.com/japananh/togo/modules/user/usermodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockCreateTaskStore struct{}

func (mockCreateTaskStore) FindTaskByCondition(
	_ context.Context,
	_ map[string]interface{},
	_ ...string,
) (*taskmodel.Task, error) {
	task := taskmodel.Task{
		Title:       "Task 1",
		Description: "Task description",
		CreatedBy:   1,
	}
	return &task, nil
}

func (mockCreateTaskStore) CountUserDailyTask(_ context.Context, createdBy int) (int, error) {
	if createdBy >= 1 {
		return 3, nil
	}
	return 0, errors.New("invalid task creator")
}

func (mockCreateTaskStore) CreateTask(_ context.Context, data *taskmodel.TaskCreate) error {
	data.Id = 2
	return nil
}

type mockUserStore struct{}

func (mockUserStore) FindUser(_ context.Context, conditions map[string]interface{}, _ ...string) (*usermodel.User, error) {
	if val, ok := conditions["id"]; ok && val.(int) > 0 {
		return &usermodel.User{
			Email:          "user@gmail.com",
			Password:       "user@123",
			DailyTaskLimit: 5,
		}, nil
	}
	return nil, common.NewCustomError(nil, "invalid task creator", "ErrInvalidCreatedBy")
}

func TestCreateTaskRepo_CreateTask(t *testing.T) {
	repo := taskrepo.NewCreateTaskRepo(mockCreateTaskStore{}, mockUserStore{})
	err := repo.CreateTask(nil, &taskmodel.TaskCreate{Title: "Task", Description: "Task description", AssigneeId: 1, CreatedBy: 1, ParentId: 1})
	assert.Nil(t, err)
}

func TestCreateTaskRepo_CreateTaskInvalidCreatedBy(t *testing.T) {
	repo := taskrepo.NewCreateTaskRepo(mockCreateTaskStore{}, mockUserStore{})
	err := repo.CreateTask(nil, &taskmodel.TaskCreate{Title: "Task", Description: "Task description", CreatedBy: 0})
	assert.NotNil(t, err)
}
