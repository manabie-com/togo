package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB holds the client for mongo
type MongoDB struct {
	client *mongo.Client
}

// NewMongoDB is the constructor for MongoDB
func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

// GetClient returns the mongo client
func (db *MongoDB) GetClient() *mongo.Client {
	return db.client
}

// Connect connects the application to MongoDB
func (db *MongoDB) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_MONGO_URI")))
	if err != nil {
		return err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	db.client = client

	return nil
}
