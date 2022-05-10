package models

import (
	"context"

	"gorm.io/gorm"
)

type TaskModel struct {
	db *gorm.DB
}

type Task struct {
	gorm.Model
	UserID     string `gorm:"type:varchar(256);not null"`
	TaskDetail string
}

func NewTaskModel(db *gorm.DB) *TaskModel {
	return &TaskModel{
		db: db,
	}
}

func (tm *TaskModel) CreateTask(ctx context.Context, userId string, taskDetail string) (newTask *Task, err error) {
	newTask = &Task{
		UserID:     userId,
		TaskDetail: taskDetail,
	}
	if err = tm.db.Create(newTask).Error; err != nil {
		return nil, err
	}
	return newTask, nil
}
