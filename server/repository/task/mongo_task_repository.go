package repository

import (
	"context"
	"os"
	"time"
	"togo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// Define a MongoDB struct with the mongo.Client as the attribute
type mongodb struct {
	conn *mongo.Client
}

// Define a Constructor to inject the connection to the repository
func NewMongoRepository(conn *mongo.Client) TaskRepository {
	return &mongodb{
		conn: conn,
	}
}

// Add a new `Task` into the database
func (db *mongodb) CreateTask(task *models.Task) (*models.Task, error) {
	collection := db.conn.Database(os.Getenv("DATABASE_NAME")).Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, task)
	return task, err
}
