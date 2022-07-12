package actions

import (
	"context"

	"pt.example/grcp-test/database"
	"pt.example/grcp-test/models"
)

type TodoTaskParam interface {
	GetAssigneeEmail() string // Used to validate user and link to task
	GetTitle() string         // Set title for task
}

func SaveTodoTask(p TodoTaskParam) (r interface{}, ok bool) {
	t := models.Task{
		Title:         p.GetTitle(),
		AssigneeEmail: p.GetAssigneeEmail(),
	}
	var tr database.Repository = &t
	tr.GetCollection().InsertOne(context.TODO(), t)

	r = "Added"
	ok = true

	return
}
