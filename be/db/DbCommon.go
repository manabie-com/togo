package db

import (
	"context"
	"time"
	"todo/be/env"
	"todo/be/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionTodo *mongo.Collection

func InitDb() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(env.CONNECTION_STRING)
	client, errDb := mongo.Connect(ctx, clientOptions)

	if utils.IsError(errDb) {
		return false
	}

	database := client.Database(env.DATABASE_NAME)

	collectionTodo = database.Collection("ColTodo")

	const INDEX_ASC = 1
	collectionTodo.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			primitive.E{Key: Todo_UserId, Value: INDEX_ASC},
			primitive.E{Key: Todo_CreatedDate, Value: INDEX_ASC},
		},
	})
	return true
}
