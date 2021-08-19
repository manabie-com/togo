package repositories

import (
	"gorm.io/gorm"
	models "github.com/manabie-com/togo/internal/models"
)

type TaskRepository struct {
	DB *gorm.DB
}

func ProvideTaskRepository(DB *gorm.DB) TaskRepository {
	return TaskRepository{DB: DB}
}

func (repo *TaskRepository) FindAll() []models.Task {
	var tasks []models.Task
	repo.DB.Find(&tasks)

	return tasks
}

func (repo *TaskRepository) FindByID(id string) models.Task {
	var task models.Task
	repo.DB.First(&task, id)

	return task
}

func (repo *TaskRepository) Save(task models.Task) models.Task {
	repo.DB.Save(&task)

	return task
}

func (repo *TaskRepository) Delete(task models.Task) {
	repo.DB.Delete(&task)
}