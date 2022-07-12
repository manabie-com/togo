package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"pt.example/grcp-test/database"
)

type Task struct {
	Title         string `bson:"title,omitempty"`
	AssigneeEmail string `bson:"assignee_email,omitempty"`
}

func (t *Task) GetCollection() *mongo.Collection {
	return database.Client.Database("manabie").Collection("tasks")
}
