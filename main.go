package main

import (
	"context"
	"log"

	"togo/cache"
	recordHandler "togo/handler/record"
	"togo/record"


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
	record.UserCollection = storage.MongoClient.Database(storage.TogoDbName).Collection(storage.UserConfigTableName)
	record.RecCollection = storage.MongoClient.Database(storage.TogoDbName).Collection(storage.UserTaskTableName)

	cache.RedisClient = cache.StartRedis()
	redisCache := cache.NewRedis(cache.RedisClient)

	repository := record.NewRepository()
	service := record.NewService(repository, redisCache)
	recordHandler := recordHandler.NewHandler(service)

	r.POST("/api/v1/task/record", recordHandler.HandleRecordTask)
	return r
}

func main() {
	r := SetUpRoute()
	r.Run(ServicePort)
}
