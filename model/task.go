package model

import (
	"time"
)

// Task model that is used for the data fetched from the database
type Task struct {
	ID int `json:"task_id"`
	Title string `json:"task_title"`
	Description string `json:"task_desc"`
	Username string `json:"created_by"`
	CreateDate time.Time `json:"date_created"` 
}

// Task model that is used for insert and update of task
type TaskUserEnteredDetails struct {
	Title string `json:"task_title"`
	Description string `json:"task_desc"`
	Username string `json:"created_by"`
}

