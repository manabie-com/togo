package task

import (
	"example.com/m/v2/internal/models"
)

// TaskUseCaseUsecase is the definition for collection of methods related to the `users` table use case
type TaskUseCase interface {
	FindTaskByUser(userID, createDate string) ([]models.Task, error)
	AddTask(task *models.Task) error
}

type taskUseCaseRepository struct {
	repository TaskRepository
}

// NewUserUsecase returns a UserUsecase attached with methods related to the `users` table use case
func NewTaskUseCase(repository TaskRepository) TaskUseCase {
	return &taskUseCaseRepository{repository: repository}
}

func (r *taskUseCaseRepository) FindTaskByUser(userID, createDate string) ([]models.Task, error) {
	return r.repository.FindTaskByUser(userID, createDate)
}

func (r *taskUseCaseRepository) AddTask(task *models.Task) error {
	return r.repository.AddTask(task)
}
