package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/manabie-com/togo/internal/iservices"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type ITaskApi interface {
	ListTasksByUserAndDate(ctx context.Context, req *http.Request) (*iservices.ListTaskResponse, *tools.TodoError)
	AddTask(ctx context.Context, req *http.Request) (*iservices.AddTaskResponse, *tools.TodoError)
}

type TaskApi struct {
	taskService iservices.ITaskService
	requestTool tools.IRequestTool
}

func (ta *TaskApi) ListTasksByUserAndDate(ctx context.Context, req *http.Request) (*iservices.ListTaskResponse, *tools.TodoError) {
	createDate := ta.requestTool.Value(req, "created_date")
	if !createDate.Valid {
		return nil, tools.NewTodoError(400, "not found created_date to check data")
	}
	return ta.taskService.ListTasksByUserAndDate(ctx, iservices.ListTaskRequest{CreatedDate: createDate.String})
}

func (ta *TaskApi) AddTask(ctx context.Context, req *http.Request) (*iservices.AddTaskResponse, *tools.TodoError) {
	t := &iservices.AddTaskRequest{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		return nil, tools.NewTodoError(500, err.Error())
	}
	return ta.taskService.AddTask(ctx, *t)
}

func NewTaskApi(db *sql.DB, contextTool tools.IContextTool, requestTool tools.IRequestTool) TaskApi {
	return TaskApi{
		taskService: services.NewTaskService(storages.NewTaskRepo(db), contextTool),
		requestTool: requestTool,
	}
}
