package controller

/*
This file will contain all method of the project
*/

const (
	User    = "user"
	Task    = "task"
	NewTask = "new-task"
)

type Controller struct {
	Frequency map[int]int
}

func NewController() *Controller {
	return &Controller{
		Frequency: make(map[int]int),
	}
}
