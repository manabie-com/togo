package interfaces

import (
	"github.com/qgdomingo/todo-app/model"
)

// These are the function signatures implemented by task_repo.go. 
// 		These are interfaced to allow the implementation of the task_repo_mock.go 
//		for the unit testing of the task_controller.go
type ITaskRepository interface {
	GetTasksDB (searchParam any) ([]model.Task, map[string]string)
	InsertTaskDB (task *model.TaskUserEnteredDetails) (bool, map[string]string)
	UpdateTaskDB (task *model.TaskUserEnteredDetails, id int) (bool, map[string]string)
	DeleteTaskDB (id int) (bool, map[string]string)
}