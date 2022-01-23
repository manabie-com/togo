package main

import (
	"log"
	"os"

	"api_service/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	// Load environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func main() {

	router := gin.Default()

	router.POST("/login", handler.Login)
	router.POST("/account/create", handler.Create)

	authorized := router.Group("/")
	// authorized.Use(middleware.TokenAuth)
	{
		authorized.POST("/logout", handler.Logout)
		authorized.GET("/account/show", handler.Show)
		authorized.PUT("/account/update", handler.Update)
	}

	log.Fatal(router.Run(":" + os.Getenv("REST_API_PORT")))
}
