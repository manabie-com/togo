package repository

import (
	"context"
	"log"
	"os"
	"time"
	"togo/models"

	"go.mongodb.org/mongo-driver/bson"
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

// Find all tasks per current day per user
func (db *mongodb) CountTask(userid string, now time.Time) (int, error) {
	collection := db.conn.Database(os.Getenv("DATABASE_NAME")).Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define filter
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	log.Println(userid)
	filter := bson.D{
		{"created_by", userid},
		{"created_at", bson.D{{"$gte", today}}},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return 0, err
	}
	var found []models.Task
	// var found []bson.M
	if err := cursor.All(context.TODO(), &found); err != nil {
		log.Fatal(err)
	}

	log.Println(len(found))
	return len(found), nil
}
