package mock

import (
	"github.com/qgdomingo/todo-app/model"
)

type TaskRepositoryMock struct {
	TaskList []model.Task
	IsTaskSuccessful bool
	ErrorMessage map[string]string
}


func (t *TaskRepositoryMock) GetTasksDB (searchParam any) ([]model.Task, map[string]string) {
	return t.TaskList, t.ErrorMessage
}

func (t *TaskRepositoryMock) InsertTaskDB (task *model.TaskUserEnteredDetails) (bool, map[string]string) {
	return t.IsTaskSuccessful, t.ErrorMessage
}

func (t *TaskRepositoryMock) UpdateTaskDB (task *model.TaskUserEnteredDetails, id int) (bool, map[string]string) {
	return t.IsTaskSuccessful, t.ErrorMessage
}

func (t *TaskRepositoryMock) DeleteTaskDB (id int) (bool, map[string]string) {
	return t.IsTaskSuccessful, t.ErrorMessage
}