package services

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/elliotchance/redismock/v8"
	"github.com/go-redis/redis/v8"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListTasks(t *testing.T) {
	ctx := context.Background()
	storage := new(storeMock)
	s := TaskService{
		Storage: storage,
	}
	expectedTasks := []*storages.Task{
		{
			ID:          "1",
			Content:     "content",
			UserID:      "firstUser",
			CreatedDate: "2021-03-01",
		},
	}
	storage.On("RetrieveTasks", ctx, mock.Anything, mock.Anything).Return(expectedTasks, nil)
	tasks, err := s.ListTasks(ctx, "firstUser", "2021-03-01")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(expectedTasks, tasks)
}

func TestAddTask(t *testing.T) {
	ctx := context.Background()
	r := redismock.NewMock()
	storage := new(storeMock)
	s := TaskService{
		Redis:   r,
		Storage: storage,
	}
	expectedTask := &storages.Task{
		Content: "content",
	}
	jsonTask, err := json.Marshal(expectedTask)
	stringReader := strings.NewReader(string(jsonTask))
	stringReadCloser := ioutil.NopCloser(stringReader)

	storage.On("AddTask", ctx, mock.Anything).Return(nil)
	r.On("IncrBy", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(redis.NewIntResult(1, nil))

	task, err := s.AddTask(ctx, stringReadCloser, "firstUser")

	r.AssertNumberOfCalls(t, "IncrBy", 1)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotEmpty(task.ID)
	assert.NotEmpty(task.UserID)
	assert.NotEmpty(task.CreatedDate)
	assert.Equal(expectedTask.Content, task.Content)
}

func TestIsReachedLimit(t *testing.T) {
	ctx := context.Background()
	r := redismock.NewMock()
	s := TaskService{
		Redis: r,
	}

	r.On("Get", ctx, mock.AnythingOfType("string")).Return(redis.NewStringResult(strconv.Itoa(LIMIT_PER_DAY), nil))
	err := s.IsReachedLimit(ctx, "firstUser")

	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(http.StatusTooManyRequests, err.Code)
	assert.Equal(err.Error(), "daily tasks limit exceeded")
}

func TestIsNotReachedLimit(t *testing.T) {
	ctx := context.Background()
	r := redismock.NewMock()
	s := TaskService{
		Redis: r,
	}

	r.On("Get", ctx, mock.Anything).Return(redis.NewStringResult(strconv.Itoa(LIMIT_PER_DAY-1), nil))
	err := s.IsReachedLimit(ctx, "firstUser")

	assert := assert.New(t)
	assert.Nil(err)
}
