package usecase

import "togo.com/pkg/repository"

type TaskUseCase interface {
	AddTask()
}
type taskUseCase struct {
	repo repository.Repository
}

func NewTaskUseCase(repo repository.Repository) TaskUseCase {
	return taskUseCase{repo: repo}
}

func (t taskUseCase) AddTask() {

}
