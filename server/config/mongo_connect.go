package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Establish Connection to MongoDB
// This function will be used to instantiate the Mongo Repositories
func ConnectMongo(uri, port string) *mongo.Client {
	mongodb_uri := fmt.Sprintf("%v:%v", os.Getenv("MONGODB_URI"), os.Getenv("DATABASE_PORT"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	return client
}
