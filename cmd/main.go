package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"togo/pkg/config"
	"togo/pkg/db_client"
	"togo/pkg/routes"
	"togo/pkg/services"
	"togo/pkg/utils"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	h := db_client.Init(c.DBUrl)
	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "manibie-todo",
		ExpirationHours: 24 * 60,
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.RedisUrl,
		Password: c.RedisPassWord,
		DB:       0,
	})

	s := services.Server{
		H:     h,
		Jwt:   jwt,
		Redis: redisClient,
	}
	routes.RegisterRoutes(r, &s)
	r.Run(c.Port)
}
