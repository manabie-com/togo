package dto

import (
	"time"

	"github.com/trinhdaiphuc/togo/database/ent"
	"github.com/trinhdaiphuc/togo/internal/entities"
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
	tasksEntity := make([]*entities.Task, len(tasks))
	for idx, task := range tasks {
		tasksEntity[idx] = Task2TaskEntity(task)
	}
	return tasksEntity
}
