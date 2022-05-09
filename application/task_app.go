package application

import (
	"github.com/jfzam/togo/domain/entity"
	"github.com/jfzam/togo/domain/repository"
)

type TaskApp struct {
	tr repository.TaskRepository
}

var _ TaskAppInterface = &TaskApp{}

type TaskAppInterface interface {
	SaveTask(*entity.Task) (*entity.Task, map[string]string)
	GetAllTask() ([]entity.Task, error)
	GetTask(uint64) (*entity.Task, error)
}

func (f *TaskApp) SaveTask(Task *entity.Task) (*entity.Task, map[string]string) {
	return f.tr.SaveTask(Task)
}

func (f *TaskApp) GetAllTask() ([]entity.Task, error) {
	return f.tr.GetAllTask()
}

func (f *TaskApp) GetTask(TaskId uint64) (*entity.Task, error) {
	return f.tr.GetTask(TaskId)
}
