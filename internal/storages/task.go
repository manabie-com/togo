package storages

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type ITaskRepo interface {
	RetrieveTasksStore(ctx context.Context, arg RetrieveTasksParams) ([]Task, *tools.TodoError)
	AddTaskStore(ctx context.Context, arg AddTaskParams) *tools.TodoError
}

type TaskRepo struct {
	*Queries
	db *sqlx.DB
}

func (tr *TaskRepo) RetrieveTasksStore(ctx context.Context, arg RetrieveTasksParams) ([]Task, *tools.TodoError) {
	tasks, err := tr.RetrieveTasks(ctx, arg)
	if err != nil {
		return nil, tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}
	return tasks, nil
}

func (tr *TaskRepo) AddTaskStore(ctx context.Context, arg AddTaskParams) *tools.TodoError {
	err := tr.AddTask(ctx, arg)
	if err != nil {
		return tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func NewTaskRepo(db DBTX) ITaskRepo {
	return &TaskRepo{
		Queries: New(db),
	}
}
