package model

import (
	"time"
)

type Task struct {
	ID int `json:"task_id"`
	Title string `json:"task_title"`
	Description string `json:"task_desc"`
	Username string `json:"created_by"`
	CreateDate time.Time `json:"date_created"` 
}

type TaskUserEnteredDetails struct {
	Title string `json:"task_title"`
	Description string `json:"task_desc"`
	Username string `json:"created_by"`
}
