package task

import (
	"togo/src"
	"togo/src/schema"
)

type ITaskController interface {
	AddTaskByOwner(context src.IContextService, data *schema.AddTaskRequest) (*schema.AddTaskResponse, error)
}
