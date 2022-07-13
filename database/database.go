package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Name string
var URI string

func Init() error {
	URI = fmt.Sprintf("mongodb://%v:%v@%v:%v", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	Name = os.Getenv("DB_NAME")

	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))

	if err != nil {
		return err
	}

	return nil
}
