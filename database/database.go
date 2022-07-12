package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Name string
var URI string

func Init() {
	URI = "mongodb://mongoadmin:secret@localhost:27017"
	Name = "manabie"

	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))

	if err != nil {
		panic(err)
	}

	// coll := Client.Database("manabie").Collection("users")
	// title := "Back to the Future"

	// coll.InsertOne(
	// 	context.TODO(),
	// 	bson.D{{Key: "title", Value: title}},
	// )
}
