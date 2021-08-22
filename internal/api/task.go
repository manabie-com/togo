package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/manabie-com/togo/internal/dto"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/repos"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type ITaskApi interface {
	ListTasksByUserAndDate(ctx context.Context, req *http.Request) (*dto.ListTaskResponse, *tools.TodoError)
	AddTask(ctx context.Context, req *http.Request) (*dto.AddTaskResponse, *tools.TodoError)
}

type TaskApi struct {
	taskService dto.ITaskService
}

func (ta *TaskApi) ListTasksByUserAndDate(ctx context.Context, req *http.Request) (*dto.ListTaskResponse, *tools.TodoError) {
	createDate := tools.Value(req, "created_date")
	if !createDate.Valid {
		return nil, tools.NewTodoError(400, "not found created_date to check data")
	}
	return ta.taskService.ListTasksByUserAndDate(ctx, dto.ListTaskRequest{CreatedDate: createDate.String})
}

func (ta *TaskApi) AddTask(ctx context.Context, req *http.Request) (*dto.AddTaskResponse, *tools.TodoError) {
	t := &dto.AddTaskRequest{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		return nil, tools.NewTodoError(500, err.Error())
	}
	return ta.taskService.AddTask(ctx, *t)
}

func NewTaskApi(db *sql.DB) TaskApi {
	return TaskApi{
		taskService: services.NewTaskService(repos.NewTaskRepo(db)),
	}
}
