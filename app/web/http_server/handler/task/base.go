package task

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"

	taskRepo "github.com/manabie-com/togo/app/repo/mongo/task"
	userRepo "github.com/manabie-com/togo/app/repo/mongo/user"
)

// Service return list apis relate to flow
type Service interface {
	Create(c echo.Context) error
	All(c echo.Context) error
}

type service struct {
	userRepo    userRepo.Repository
	taskRepo    taskRepo.Repository
	mongoClient *mongo.Client
}

// NewService return handler instance
func NewService(
	userRepoInstance userRepo.Repository,
	taskRepoInstance taskRepo.Repository,
	mongoClient *mongo.Client,
) Service {
	return &service{
		userRepo:    userRepoInstance,
		taskRepo:    taskRepoInstance,
		mongoClient: mongoClient,
	}
}
