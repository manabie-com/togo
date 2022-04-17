package interfaces

import (
	"github.com/qgdomingo/todo-app/model"
)

type ITaskRepository interface {
	GetTasksDB (searchParam any) ([]model.Task, map[string]string)
	InsertTaskDB (task *model.TaskUserEnteredDetails) (bool, map[string]string)
	UpdateTaskDB (task *model.TaskUserEnteredDetails, id int) (bool, map[string]string)
	DeleteTaskDB (id int) (bool, map[string]string)
}