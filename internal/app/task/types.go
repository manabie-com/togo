package task

import (
	"time"
)

type (
	CreateTaskReq struct {
		Name        string `json:"name" binding:"required"`
		Assignee    string `json:"assignee,omitempty"`
		AssignDate  time.Time
		Description string    `json:"description,omitempty"`
		Status      string    `json:"status"  binding:"required"`
		StartDate   time.Time `json:"start_date,omitempty"`
		EndDate     time.Time `json:"end_date,omitempty"`
		Creator     string
	}

	AssignTaskReq struct {
		ID       int    `json:"id" binding:"required,min=1"`
		Assignee string `json:"assignee"  binding:"required,min=3"`
	}

	GetTaskByIdReq struct {
		ID int `uri:"id" binding:"required,min=1"`
	}

	ListTasksReq struct {
		PageNumber int `uri:"id" binding:"required,min=1"`
		PageSize   int `uri:"id" binding:"required,min=1"`
	}

	DeleteTaskByIdReq struct {
		ID int `uri:"id" binding:"required,min=1"`
	}

	Task struct {
		ID          int32     `json:"id"`
		Name        string    `json:"name"`
		Assignee    string    `json:"assignee"`
		AssignDate  time.Time `json:"assign_date"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		Creator     string    `json:"creator"`
		CreatedAt   time.Time `json:"created_at"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
	}
)
