package task

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserConfig contains user limit config info
type UserConfig struct {
	UserId string `bson:"user_id"`
	Limit  int    `bson:"limit"`
}

// UserTask contains users' task record info
type UserTask struct {
	UserId    string    `bson:"user_id"`
	Task      string    `bson:"task"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type RecordRepository struct {
	mongoClient *mongo.Client
}

func NewRepository(mongoClient *mongo.Client) *RecordRepository {
	return &RecordRepository{
		mongoClient: mongoClient,
	}
}

func (r *RecordRepository) InsertUserTask(userId, task string, updatedAt time.Time) error {
	db := r.mongoClient.Database(TogoDbName)
	coll := db.Collection(UserTaskTableName)

	userTask := UserTask{
		UserId: userId,
		Task: task,
		UpdatedAt: updatedAt,
	}
	_, err := coll.InsertOne(context.Background(), userTask)

	if err != nil {
		return fmt.Errorf("insert user task error: %v", err)
	}
	return nil
}

func (r *RecordRepository) GetUserConfig(userId string) (*UserConfig, error) {
	db := r.mongoClient.Database(TogoDbName)
	coll := db.Collection(UserConfigTableName)

	filter := bson.D{{"user_id", userId}}
	var result UserConfig
	err := coll.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

