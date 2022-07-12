package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"pt.example/grcp-test/database"
)

type User struct {
	Email                 string             `bson:"email,omitempty"`
	MaxAssignedTaskPerDay uint8              `bson:"max_assigned_task_per_day,omitempty"`
	LastAssignedTime      primitive.DateTime `bson:"last_assigned_time,omitempty"`
}

func (t *User) GetCollection() *mongo.Collection {
	return database.Client.Database("manabie").Collection("users")
}
