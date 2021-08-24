package test

import (
	"context"
	"github.com/manabie-com/togo/internal/iservices"
	mockRepo "github.com/manabie-com/togo/internal/mocks/storages"
	mockTool "github.com/manabie-com/togo/internal/mocks/tools"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestListTaskByUserAndDate(t *testing.T) {
	t.Run("Get list task by user and date success", func(t *testing.T) {
		req := iservices.ListTaskRequest{CreatedDate: "2021-8-21"}
		resExpect := []iservices.Task{
			{ID: "1", Content: "valid", UserID: "1", CreatedDate: req.CreatedDate},
			{ID: "2", Content: "valid2", UserID: "1", CreatedDate: req.CreatedDate},
		}
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("1", nil)
		taskRepo := mockRepo.ITaskRepo{}
		taskRepo.On("RetrieveTasksStore", context.TODO(),
			storages.RetrieveTasksParams{UserID: "1", CreatedDate: req.CreatedDate}).
			Return([]storages.Task{
				{ID: "1", Content: "valid", UserID: "1", CreatedDate: req.CreatedDate},
				{ID: "2", Content: "valid2", UserID: "1", CreatedDate: req.CreatedDate},
			}, nil)
		taskService := services.NewTaskService(&taskRepo, &contextTool)
		res, err := taskService.ListTasksByUserAndDate(context.TODO(), req)
		require.Nil(t, err)
		require.NotNil(t, res)
		assert.Equal(t, resExpect, res.Data)
	})
	t.Run("Get list task by user and date fail by convert context", func(t *testing.T) {
		req := iservices.ListTaskRequest{CreatedDate: "2021-8-21"}
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("", tools.NewTodoError(500, "fail to convert ctx"))
		taskRepo := mockRepo.ITaskRepo{}
		taskService := services.NewTaskService(&taskRepo, &contextTool)
		res, err := taskService.ListTasksByUserAndDate(context.TODO(), req)
		require.Nil(t, res)
		require.Error(t, err, "fail to convert ctx")
	})
	t.Run("Get list task by user and date fail by get data from repo", func(t *testing.T) {
		req := iservices.ListTaskRequest{CreatedDate: "2021-8-21"}
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("1", nil)
		taskRepo := mockRepo.ITaskRepo{}
		taskRepo.On("RetrieveTasksStore", context.TODO(),
			storages.RetrieveTasksParams{UserID: "1", CreatedDate: req.CreatedDate}).
			Return(nil, tools.NewTodoError(500, "fail to get data from repo"))
		taskService := services.NewTaskService(&taskRepo, &contextTool)
		res, err := taskService.ListTasksByUserAndDate(context.TODO(), req)
		require.Nil(t, res)
		require.Error(t, err, "fail to get data from repo")
	})
}

func TestAddTask(t *testing.T) {
	t.Run("Add Task success", func(t *testing.T) {
		req := iservices.AddTaskRequest{Content: "valid content"}
		now := time.Now()
		resExpect := iservices.Task{Content: "valid content", ID: "1", UserID: "1", CreatedDate: now.Format("2006-01-02")}
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("1", nil)
		taskRepo := mockRepo.ITaskRepo{}
		taskRepo.On("AddTaskStore", context.TODO(), mock.Anything).Return(nil)
		taskService := services.NewTaskService(&taskRepo, &contextTool)
		res, err := taskService.AddTask(context.TODO(), req)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.NotNil(t, res.Data)
		assert.Equal(t, resExpect.Content, res.Data.Content)
		assert.Equal(t, resExpect.UserID, res.Data.UserID)
		assert.Equal(t, resExpect.CreatedDate, res.Data.CreatedDate)
	})
	t.Run("Add Task Fail by convert context", func(t *testing.T) {
		req := iservices.AddTaskRequest{Content: "valid content"}
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("", tools.NewTodoError(500, "fail to convert ctx"))
		taskRepo := mockRepo.ITaskRepo{}
		taskService := services.NewTaskService(&taskRepo, &contextTool)
		res, err := taskService.AddTask(context.TODO(), req)
		require.Nil(t, res)
		require.Error(t, err, "fail to convert ctx")
	})
	t.Run("Add Task Fail by add to repo", func(t *testing.T) {
		req := iservices.AddTaskRequest{Content: "valid content"}
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("1", nil)
		taskRepo := mockRepo.ITaskRepo{}
		taskRepo.On("AddTaskStore", context.TODO(), mock.Anything).Return(tools.NewTodoError(500, "fail to add task to repo"))
		taskService := services.NewTaskService(&taskRepo, &contextTool)
		res, err := taskService.AddTask(context.TODO(), req)
		require.Nil(t, res)
		require.NotNil(t, err)
		require.Error(t, err, "fail to add task to repo")
	})
}
