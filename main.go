package main

import (
	"context"
	"log"

	"togo/cache"
	taskHandler "togo/handler/task"
	"togo/task"

	"github.com/gin-gonic/gin"

	"togo/storage"
)

const (
	ServicePort = ":8083"
)

func SetUpRoute() *gin.Engine {
	r := gin.Default()

	mongoClient := storage.StartMongo()
	ctx, _ := context.WithTimeout(context.Background(), storage.MongoClientTimeout)
	err := mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	storage.MongoClient = mongoClient
	task.UserCollection = storage.MongoClient.Database(storage.TogoDbName).Collection(storage.UserConfigTableName)
	task.RecordCollection = storage.MongoClient.Database(storage.TogoDbName).Collection(storage.UserTaskTableName)

	cache.RedisClient = cache.StartRedis()
	redisCache := cache.NewRedis(cache.RedisClient)

	repository := task.NewRepository()
	service := task.NewService(repository, redisCache)
	taskHandler := taskHandler.NewHandler(service)

	r.POST("/api/v1/task/record", taskHandler.HandleRecordTask)
	return r
}

func main() {
	r := SetUpRoute()
	r.Run(ServicePort)
}
