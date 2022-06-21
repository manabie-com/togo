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

func setUpRoute() *gin.Engine {
	r := gin.Default()

	mongoClient := storage.StartMongo()
	ctx, _ := context.WithTimeout(context.Background(), storage.MongoClientTimeout)
	err := mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	storage.MongoClient = mongoClient

	cache.RedisClient = cache.StartRedis()
	redisCache := cache.NewRedis(cache.RedisClient)

	repository := task.NewRepository(mongoClient)
	service := task.NewService(repository, redisCache)
	taskHandler := taskHandler.NewHandler(service)

	r.POST("/api/v1/task/record", taskHandler.HandleRecordTask)
	return r
}

func main() {
	r := setUpRoute()
	r.Run(":8083") // listen and serve on 0.0.0.0:8083
}
