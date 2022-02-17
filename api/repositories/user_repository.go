package repositories

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/kier1021/togo/api/models"
	"github.com/kier1021/togo/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	mongoDB        *db.MongoDB
	userCollection *mongo.Collection
}

func NewUserRepository(mongoDB *db.MongoDB) *UserRepository {
	client := mongoDB.GetClient()
	userCollection := client.Database(os.Getenv("DB_MONGODB_NAME")).Collection("user")

	return &UserRepository{
		mongoDB:        mongoDB,
		userCollection: userCollection,
	}
}

func (repo *UserRepository) CreateUser(user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data := make(map[string]interface{})

	data["user_name"] = user.UserName
	data["max_tasks"] = user.MaxTasks
	data["_id"] = primitive.NewObjectID()

	res, err := repo.userCollection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo *UserRepository) GetUser(filter map[string]interface{}) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := filterToObjectID(filter, "_id"); err != nil {
		return nil, err
	}

	var user models.User
	res := repo.userCollection.FindOne(ctx, filter)

	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, res.Err()
	}

	res.Decode(&user)
	return &user, nil
}

func (repo *UserRepository) GetUsers(filter map[string]interface{}) (users []models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := filterToObjectID(filter, "_id"); err != nil {
		return nil, err
	}

	cur, err := repo.userCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
