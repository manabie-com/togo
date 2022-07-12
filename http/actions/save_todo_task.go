package actions

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"pt.example/grcp-test/database"
	"pt.example/grcp-test/models"
)

type TodoTaskParam interface {
	GetAssigneeEmail() string // Used to validate user and link to task
	GetTitle() string         // Set title for task
}

func SaveTodoTask(ctx context.Context, p TodoTaskParam) (r *mongo.InsertOneResult, err error) {
	t := models.Task{
		Title:         p.GetTitle(),
		AssigneeEmail: p.GetAssigneeEmail(),
	}
	var tr database.Repository = &t

	r, err = tr.GetCollection().InsertOne(ctx, t)

	return
}
