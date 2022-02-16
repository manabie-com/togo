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

func (repo *UserTaskRepository) CreateUser(user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data := make(map[string]interface{})

	data["user_name"] = user.UserName
	data["max_tasks"] = user.MaxTasks
	data["ins_day"] = user.InsDay
	data["tasks"] = []map[string]interface{}{}
	data["_id"] = primitive.NewObjectID()

	res, err := repo.userTaskCollection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo *UserTaskRepository) AddTaskToUser(userTask models.UserTask) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	task := map[string]interface{}{
		"title":       userTask.Title,
		"description": userTask.Description,
	}

	_, err := repo.userTaskCollection.UpdateOne(
		ctx,
		bson.M{"user_name": userTask.UserName, "ins_day": userTask.InsDay},
		bson.D{
			{
				Key:   "$push",
				Value: map[string]interface{}{"tasks": task},
			},
			{
				Key: "$set",
				Value: map[string]interface{}{
					"user_name": userTask.UserName,
					"ins_day":   userTask.InsDay,
					"max_tasks": userTask.MaxTasks,
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

func (repo *UserTaskRepository) GetUser(filter map[string]interface{}) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if id, ok := filter["_id"].(string); ok {
		idPrimitive, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		filter["_id"] = idPrimitive
	}

	var user models.User
	res := repo.userTaskCollection.FindOne(ctx, filter)

	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, res.Err()
	}

	res.Decode(&user)
	return &user, nil
}
