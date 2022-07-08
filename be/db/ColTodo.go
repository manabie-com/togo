package db

import (
	"context"
	"todo/be/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Todo_UserId      = "UserId"
	Todo_CreatedDate = "CreatedDate"
)

type Todo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId      string             `json:"UserId,omitempty" bson:"UserId,omitempty"`
	CreatedDate int64              `json:"CreatedDate,omitempty" bson:"CreatedDate,omitempty"`
	Text        string             `json:"Text,omitempty" bson:"Text,omitempty"`
}

func Todo_add(ctx context.Context, todo Todo) error {
	_, err := collectionTodo.InsertOne(ctx, todo)
	return err
}

func Todo_count(ctx context.Context, query bson.M) int64 {
	count, err := collectionTodo.CountDocuments(ctx, query)
	if utils.IsError(err) {
		return -1
	}
	return count
}

func Todo_delete(ctx context.Context, query bson.M) error {
	_, err := collectionTodo.DeleteMany(ctx, query)
	return err
}
