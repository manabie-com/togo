package services

import (
	"context"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/core"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var todoService *ToDoService

func TestValidateUser(t *testing.T) {
	tests := []struct {
		userID string
		pwd    string
		want   bool
	}{
		{
			userID: "test",
			pwd:    "test",
			want:   true,
		},
		{
			userID: "test",
			pwd:    "gg",
			want:   false,
		},
	}

	for _, test := range tests {
		result := todoService.ValidateUser(context.Background(), test.userID, test.pwd)
		assert.Equal(t, test.want, result)
	}
}

func TestAddNewTask(t *testing.T) {

	ctx := context.Background()
	testUserID := "test"
	err := todoService.AddTask(ctx, testUserID, &entities.Task{
		Content: "test",
	})

	assert.Nil(t, err)

	result, err := todoService.ListTasks(ctx, testUserID, utils.FormatTimeToString(time.Now()))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "test", result[0].Content)
}

func TestListTasks(t *testing.T) {

	ctx := context.Background()
	testUserID := "test"
	err := todoService.TaskRepository.AddTask(ctx, &entities.Task{
		Content:     "test",
		CreatedDate: utils.FormatTimeToString(time.Now()),
	})

	assert.Nil(t, err)

	result, err := todoService.ListTasks(ctx, testUserID, utils.FormatTimeToString(time.Now()))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "test", result[0].Content)
}

func TestIsAllowToAddTask(t *testing.T) {

	ctx := context.Background()
	testUserID := "test"
	for i := 0; i < int(config.LimitAllowTaskPerDay); i++ {
		err := todoService.isAllowToAddTask(ctx, testUserID)
		assert.Nil(t, err, "it should nil success allow to add")
		err = todoService.AddTask(ctx, testUserID, &entities.Task{
			Content: "test",
		})
		assert.Nil(t, err, "it should nil success allow to add")

	}
	err := todoService.isAllowToAddTask(ctx, testUserID)
	assert.NotNil(t, err, "it should not nil success limit allow to add")

	if internalError, ok := err.(*core.InternalError); ok {
		assert.Equal(t, uint8(core.ERROR_CODE_EXCEED_TASK_LIMITS), internalError.ErrCode, "it should be equal")
	}

}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	cleanup()
	os.Exit(ret)
}

func setup() {
	todoService = &ToDoService{
		UserRepository: &storages.MockUserRepository{},
		TaskRepository: &storages.MockTaskRepository{},
	}
}

func cleanup() {
	todoService = nil
}
