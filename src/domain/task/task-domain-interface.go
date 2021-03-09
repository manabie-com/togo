package task

import (
	"togo/src"
	"togo/src/schema"
)

type ITaskWorkflow interface {
	AddTaskByOwner(context src.IContextService, data *schema.AddTaskRequest) (*schema.AddTaskResponse, error)
}
