package main

import (
	"context"

	taskHandler "togo/handler/task"
	"togo/cache"
	"togo/task"

	"log"

	"github.com/gin-gonic/gin"
	"togo/storage"
)

func main() {
	r := gin.Default()

	mongoClient := storage.StartMongo()
	ctx, cancel := context.WithTimeout(context.Background(), storage.MongoClientTimeout)
	defer cancel()
	err := mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	storage.MongoClient = mongoClient
	defer mongoClient.Disconnect(ctx)

	cache.RedisClient = cache.StartRedis()
	redisCache := cache.NewRedis(cache.RedisClient)

	repository := task.NewRepository(mongoClient)
	service := task.NewService(repository, *redisCache)
	taskHandler := taskHandler.NewHandler(service)

	r.POST("/api/v1/task/record", taskHandler.HandleRecordTask)

	r.Run() // listen and serve on 0.0.0.0:8080
}
