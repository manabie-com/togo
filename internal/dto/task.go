package dto

import (
	"github.com/trinhdaiphuc/togo/database/ent"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"time"
)

func Task2TaskEntity(task *ent.Task) *entities.Task {
	return &entities.Task{
		ID:        task.ID,
		Name:      task.Name,
		Content:   task.Content,
		UserID:    task.UserID,
		CreatedAt: task.CreatedAt.Format(time.RFC3339),
		UpdatedAt: task.UpdatedAt.Format(time.RFC3339),
	}
}

func Tasks2TasksEntity(tasks []*ent.Task) []*entities.Task {
	var tasksEntity []*entities.Task
	for _, task := range tasks {
		tasksEntity = append(tasksEntity, Task2TaskEntity(task))
	}
	return tasksEntity
}