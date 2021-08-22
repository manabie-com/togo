package repositories

import (
	"fmt"
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
	result := repo.DB.Find(&tasks)
	if result.Error != nil {
		panic("Error when final all records");
	}
	return tasks
}

func (repo *TaskRepository) FindByID(id string) models.Task {
	var task models.Task
	repo.DB.First(&task, id)

	return task
}

func (repo *TaskRepository) Create(task models.Task) models.Task {
	result := repo.DB.Create(&task)
	if result.Error != nil {
		fmt.Println(result.Error)
		panic("Error, insert task to into database!")
	}
	return task
}

func (repo *TaskRepository) Delete(task models.Task) {
	repo.DB.Delete(&task)
}