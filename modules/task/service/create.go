package service

import (
	"github.com/khoale193/togo/models"
)

type Task struct {
	Name     string
	MemberID int
}

func (a *Task) CreateTask() error {
	(&models.Task{
		Name:     a.Name,
		MemberID: a.MemberID,
	}).CreateTask()
	return nil
}
