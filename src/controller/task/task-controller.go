package task

import (
	"togo/src"
	"togo/src/domain/task"
	"togo/src/schema"
)

type TaskController struct {
	taskWokflow task.ITaskWorkflow
}

func (this *TaskController) AddTaskByOwner(context src.IContextService, data *schema.AddTaskRequest) (*schema.AddTaskResponse, error) {
	if err := context.CheckPermission([]string{src.CREATE_TASK}); err != nil {
		return nil, err
	}

	return this.taskWokflow.AddTaskByOwner(context, data)
}

func NewTaskController() ITaskController {
	return &TaskController{
		task.NewTaskWorkflow(),
	}
}
