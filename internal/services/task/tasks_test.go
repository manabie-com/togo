package task

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/manabie-com/togo/internal/entity"
	"github.com/manabie-com/togo/internal/mocks"
	"github.com/manabie-com/togo/pkg/customcontext"
)

// Test_CreateTaskHappyCases to test for creating successfully.
func Test_CreateTaskHappyCases(t *testing.T) {

	taskStorage := new(mocks.TaskStorage)
	userStorage := new(mocks.UserStorage)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	userTest := &entity.User{ID: "123", MaxTodoPerday: 5, Password: "1"}
	taskTest := &entity.Task{
		Content: "content",
		UserID:  userTest.ID,
	}
	createdDate := time.Now().Format("2006-01-02")
	ctx = customcontext.SetUserIDToContext(ctx, userTest.ID)
	content := "content"
	taskStorage.On("AddTask", ctx, taskTest).Return(nil)
	taskStorage.On("GetNumberOfTasks", ctx, userTest.ID, createdDate).Return(2, nil)
	userStorage.On("FindByID", ctx, userTest.ID).Return(userTest, nil)

	taskSvc := NewTaskService(taskStorage, userStorage)
	task, err := taskSvc.Create(ctx, content)
	assert.Equal(t, nil, err)
	assert.Equal(t, taskTest.Content, task.Content)
	assert.Equal(t, taskTest.UserID, task.UserID)
}

func Test_CreateTaskFailCases(t *testing.T) {
	testcases := []struct {
		name              string
		err               error
		numberOfTaskToday int
		responseFindByID  *entity.User
		userIDCtx         string
	}{
		{
			name:              "Should return ErrReachedOutTaskTodoPerDay when number of created tasks today seems enough.",
			err:               ErrReachedOutTaskTodoPerDay,
			responseFindByID:  &entity.User{ID: "123", MaxTodoPerday: 5, Password: "1"},
			numberOfTaskToday: 5,
			userIDCtx:         "123",
		},
		{
			name:              "Should return UserNotExistedError when userID in context is not exist in system.",
			err:               NewUserNotExistedError("1"),
			responseFindByID:  nil,
			numberOfTaskToday: 5,
			userIDCtx:         "1",
		},
	}
	for _, test := range testcases {

		t.Run(test.name, func(t *testing.T) {
			taskStorage := new(mocks.TaskStorage)
			userStorage := new(mocks.UserStorage)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			createdDate := time.Now().Format("2006-01-02")
			ctx = customcontext.SetUserIDToContext(ctx, test.userIDCtx)
			taskStorage.On("GetNumberOfTasks", ctx, test.userIDCtx, createdDate).Return(test.numberOfTaskToday, nil)
			userStorage.On("FindByID", ctx, test.userIDCtx).Return(test.responseFindByID, nil)
			taskSvc := NewTaskService(taskStorage, userStorage)
			task, err := taskSvc.Create(ctx, "content")
			if task != nil {
				t.Fatal("task is not nil")
			}
			assert.Equal(t, test.err, err)

		})
	}
}

func Test_ListTasks(t *testing.T) {
	taskStorage := new(mocks.TaskStorage)
	userStorage := new(mocks.UserStorage)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	userIDTest := "123"
	tasks := []*entity.Task{
		&entity.Task{
			ID:      "1",
			Content: "content1",
			UserID:  "123",
		},
		&entity.Task{
			ID:      "2",
			Content: "content2",
			UserID:  "123",
		},
	}
	createdDate := time.Now().Format("2006-01-02")
	ctx = customcontext.SetUserIDToContext(ctx, userIDTest)
	taskStorage.On("RetrieveTasks", ctx, userIDTest, createdDate).Return(tasks, nil)
	taskSvc := NewTaskService(taskStorage, userStorage)
	actual, err := taskSvc.List(ctx, createdDate)
	assert.Equal(t, nil, err)
	assert.Equal(t, tasks, actual)
}
