package services

import (
	"context"
	"errors"
	"testing"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/assert"
)

var (
	retrieveTasksMock  func(ctx context.Context, userID, createdDate string) ([]*storages.Task, error)
	addTaskMock        func(ctx context.Context, t *storages.Task) error
	validateUserMock   func(ctx context.Context, userID, pwd string) bool
	loadTasksCountMock func(ctx context.Context, userID, createdDate string) (cnt int, err error)
	maxTodo            func(ctx context.Context, userID string) (int, error)
)

type storageRepoMock struct{}

func (s *storageRepoMock) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	return retrieveTasksMock(ctx, userID, createdDate)
}
func (s *storageRepoMock) AddTask(ctx context.Context, t *storages.Task) error {
	return addTaskMock(ctx, t)
}
func (s *storageRepoMock) ValidateUser(ctx context.Context, userID, pwd string) bool {
	return validateUserMock(ctx, userID, pwd)
}
func (s *storageRepoMock) LoadTasksCount(ctx context.Context, userID, createdDate string) (cnt int, err error) {
	return loadTasksCountMock(ctx, userID, createdDate)
}
func (s *storageRepoMock) MaxTodo(ctx context.Context, userID string) (int, error) {
	return maxTodo(ctx, userID)
}

func TestToDoService_ValidateUser(t *testing.T) {
	tdService := NewToDoService(&storageRepoMock{})
	ctx := context.Background()
	validateUserMock = func(ctx context.Context, userID, pwd string) bool {
		return true
	}
	assert.EqualValues(t, true, tdService.ValidateUser(ctx, "1", "1pwd"))
	validateUserMock = func(ctx context.Context, userID, pwd string) bool {
		return false
	}
	assert.EqualValues(t, false, tdService.ValidateUser(ctx, "2", "1pwd"))
	assert.EqualValues(t, false, tdService.ValidateUser(ctx, "2", "2pwd"))
}

func TestToDoService_ListTasks(t *testing.T) {
	tdService := NewToDoService(&storageRepoMock{})
	ctx := context.Background()
	var (
		tasks []*storages.Task
		err   error
	)

	retrieveTasksMock = func(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
		return []*storages.Task{
			{
				ID:          "1",
				Content:     "c1",
				UserID:      "1",
				CreatedDate: "2021-04-30",
			},
			{
				ID:          "2",
				Content:     "c2",
				UserID:      "1",
				CreatedDate: "2021-04-30",
			},
		}, nil

	}
	tasks, err = tdService.ListTasks(ctx, "1", "2021-04-30")
	assert.Nil(t, err)
	assert.NotNil(t, tasks)
	assert.EqualValues(t, 2, len(tasks))
	assert.EqualValues(t, "1", tasks[0].ID)
	assert.EqualValues(t, "c1", tasks[0].Content)
	assert.EqualValues(t, "1", tasks[0].UserID)
	assert.EqualValues(t, "2021-04-30", tasks[0].CreatedDate)
	assert.EqualValues(t, "2", tasks[1].ID)
	assert.EqualValues(t, "c2", tasks[1].Content)
	assert.EqualValues(t, "1", tasks[1].UserID)
	assert.EqualValues(t, "2021-04-30", tasks[1].CreatedDate)

	retrieveTasksMock = func(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
		return nil, nil
	}
	tasks, err = tdService.ListTasks(ctx, "1", "2021-04-30")
	assert.Nil(t, err)
	assert.Nil(t, tasks)

	retrieveTasksMock = func(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
		return nil, errors.New("Some error")
	}
	tasks, err = tdService.ListTasks(ctx, "1", "2021-04-30")
	assert.NotNil(t, err)
	assert.Nil(t, tasks)
}

func TestToDoService_AddTask(t *testing.T) {
	tdService := NewToDoService(&storageRepoMock{})
	ctx := context.Background()
	var (
		task *storages.Task
		err  error
	)

	task, err = tdService.AddTask(ctx, "1", []byte(`{"content": "exampl`))
	assert.NotNil(t, err)
	assert.Nil(t, task)

	addTaskMock = func(ctx context.Context, t *storages.Task) error {
		return nil
	}
	loadTasksCountMock = func(ctx context.Context, userID, createdDate string) (cnt int, err error) {
		return 2, nil
	}
	maxTodo = func(ctx context.Context, userID string) (int, error) {
		return 5, nil
	}
	task, err = tdService.AddTask(ctx, "1", []byte(`{"content": "example"}`))
	assert.Nil(t, err)
	assert.NotNil(t, task)
	assert.EqualValues(t, "1", task.UserID)
	assert.EqualValues(t, "example", task.Content)

	addTaskMock = func(ctx context.Context, t *storages.Task) error {
		return errors.New("some error")
	}
	loadTasksCountMock = func(ctx context.Context, userID, createdDate string) (cnt int, err error) {
		return 2, nil
	}
	maxTodo = func(ctx context.Context, userID string) (int, error) {
		return 5, nil
	}
	task, err = tdService.AddTask(ctx, "1", []byte(`{"content": "example"}`))
	assert.NotNil(t, err)
	assert.Nil(t, task)

	addTaskMock = func(ctx context.Context, t *storages.Task) error {
		return nil
	}
	loadTasksCountMock = func(ctx context.Context, userID, createdDate string) (cnt int, err error) {
		return 0, errors.New("some error")
	}
	maxTodo = func(ctx context.Context, userID string) (int, error) {
		return 5, nil
	}
	task, err = tdService.AddTask(ctx, "1", []byte(`{"content": "example"}`))
	assert.Nil(t, err)
	assert.NotNil(t, task)

	tdService = NewToDoService(&storageRepoMock{})
	addTaskMock = func(ctx context.Context, t *storages.Task) error {
		return nil
	}
	loadTasksCountMock = func(ctx context.Context, userID, createdDate string) (cnt int, err error) {
		return 0, nil
	}
	maxTodo = func(ctx context.Context, userID string) (int, error) {
		return 2, nil
	}
	tdService.AddTask(ctx, "1", []byte(`{"content": "example"}`))
	tdService.AddTask(ctx, "1", []byte(`{"content": "example"}`))
	task, err = tdService.AddTask(ctx, "1", []byte(`{"content": "example"}`))
	assert.NotNil(t, err)
	assert.EqualValues(t, err, ErrUserReachDailyRequestLimit)
	assert.Nil(t, task)

	tdService = NewToDoService(&storageRepoMock{})
	addTaskMock = func(ctx context.Context, t *storages.Task) error {
		return nil
	}
	loadTasksCountMock = func(ctx context.Context, userID, createdDate string) (cnt int, err error) {
		return 0, nil
	}
	maxTodo = func(ctx context.Context, userID string) (int, error) {
		return 0, errors.New("some error")
	}
	task, err = tdService.AddTask(ctx, "1", []byte(`{"content": "example"}`))
	assert.NotNil(t, err)
	assert.Nil(t, task)
}
