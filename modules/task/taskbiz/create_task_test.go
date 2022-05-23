package taskbiz_test

import (
	"context"
	"github.com/japananh/togo/modules/task/taskbiz"
	"github.com/japananh/togo/modules/task/taskmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockCreateTaskStore struct{}

type mockUserStore struct{}

type mockCreateTaskRepo struct {
	store     mockCreateTaskStore
	userStore mockUserStore
}

func (mockCreateTaskRepo) CreateTask(_ context.Context, _ *taskmodel.TaskCreate) error {
	return nil
}

func TestCreateTaskBiz_CreateTask(t *testing.T) {
	repo := mockCreateTaskRepo{store: mockCreateTaskStore{}, userStore: mockUserStore{}}
	biz := taskbiz.NewCreateTaskBiz(repo)
	err := biz.CreateTask(nil, &taskmodel.TaskCreate{Title: "Task 1", Description: "Task Description", CreatedBy: 1})
	assert.Nil(t, err)
}
