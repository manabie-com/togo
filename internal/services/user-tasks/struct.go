package user_tasks

import "github.com/google/uuid"

type createTaskRequest struct {
	UserID  uuid.UUID `json:"user_id" binding:"required"`
	Content string    `json:"content" binding:"required"`
}

type updateTaskRequest struct {
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	TaskLimit int       `json:"task_limit" binding:"required"`
}
