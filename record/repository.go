package record

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserCollection *mongo.Collection
	RecCollection *mongo.Collection
)

// UserConfig contains user limit config info
type UserConfig struct {
	UserId string `bson:"user_id"`
	Limit  int    `bson:"limit"`
}

// UserTask contains users' record record info
type UserTask struct {
	UserId    string    `bson:"user_id"`
	Task      string    `bson:"task"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type RecRepository struct {
}

func NewRepository() *RecRepository {
	return &RecRepository{
	}
}

func (r *RecRepository) InsertUserTask(userId, task string, updatedAt time.Time) error {
	userTask := UserTask{
		UserId: userId,
		Task: task,
		UpdatedAt: updatedAt,
	}
	_, err := RecCollection.InsertOne(context.Background(), userTask)

	if err != nil {
		return fmt.Errorf("insert user record error: %v", err)
	}
	return nil
}

func (r *RecRepository) GetUserConfig(userId string) (*UserConfig, error) {
	filter := bson.D{{"user_id", userId}}
	var result UserConfig
	err := UserCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

