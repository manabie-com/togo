package repositories

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/kier1021/togo/api/models"
	"github.com/kier1021/togo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserTaskRepository struct {
	mongoDB            *db.MongoDB
	userTaskCollection *mongo.Collection
}

func NewUserTaskRepository(mongoDB *db.MongoDB) *UserTaskRepository {

	client := mongoDB.GetClient()
	userTaskCollection := client.Database(os.Getenv("DB_MONGODB_NAME")).Collection("user_task")

	return &UserTaskRepository{
		mongoDB:            mongoDB,
		userTaskCollection: userTaskCollection,
	}
}

func (repo *UserTaskRepository) AddTaskToUser(user models.User, userTask models.Task, insDay string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userIDPrimitive, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	task := map[string]interface{}{
		"title":       userTask.Title,
		"description": userTask.Description,
	}

	_, err = repo.userTaskCollection.UpdateOne(
		ctx,
		bson.M{"user_name": user.UserName, "ins_day": insDay},
		bson.D{
			{
				Key:   "$push",
				Value: map[string]interface{}{"tasks": task},
			},
			{
				Key: "$set",
				Value: map[string]interface{}{
					"user_name": user.UserName,
					"ins_day":   insDay,
					"user_id":   userIDPrimitive,
				},
			},
		},
		options.Update().SetUpsert(true),
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserTaskRepository) GetUserTask(filter map[string]interface{}) (*models.UserTask, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := filterToObjectID(filter, "_id", "user_id"); err != nil {
		return nil, err
	}

	var userTask models.UserTask
	res := repo.userTaskCollection.FindOne(ctx, filter)

	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, res.Err()
	}

	res.Decode(&userTask)
	return &userTask, nil
}
