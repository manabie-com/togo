package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"pt.example/grcp-test/database"
)

type Task struct {
	Title  string `bson:"title,omitempty"`
	UserId string `bson:"user_id,omitempty"`
}

func (t *Task) GetCollection() *mongo.Collection {
	return database.Client.Database("manabie").Collection("tasks")
}
