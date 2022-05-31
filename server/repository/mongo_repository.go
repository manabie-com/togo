package repository

import (
	"context"
	"os"
	"time"
	"togo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	conn *mongo.Client
}

func NewMongoRepository(conn *mongo.Client) TaskRepository {
	return &mongodb{
		conn: conn,
	}
}

// Implement Create method for Task Repository
func (db *mongodb) Create(task *models.Task) (*models.Task, error) {
	collection := db.conn.Database(os.Getenv("DATABASE_NAME")).Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, task)
	return task, err
}
