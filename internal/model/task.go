package model

// Task represents the task model
type Task struct {
	Base
	Content string
	UserID  int
}
