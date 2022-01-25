package main

import (
	"log"
	"os"

	"api_service/handler"
	"api_service/middleware"

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

	router.POST("/login", handler.LoginAccount)
	router.POST("/account/create", handler.CreateAccount)

	authorized := router.Group("/")
	authorized.Use(middleware.TokenAuth)
	{
		authorized.POST("/logout", handler.LogoutAccount)
		authorized.GET("/account/show", handler.ShowAccount)
		authorized.PUT("/account/update", handler.UpdateAccount)

		authorized.POST("todo/create", handler.CreateTodo)
		authorized.GET("todo/get/:id", handler.GetTodo)
		authorized.PUT("todo/update/:id", handler.UpdateTodo)
		authorized.DELETE("todo/delete/:id", handler.DeleteTodo)
	}

	log.Fatal(router.Run(":" + os.Getenv("REST_API_PORT")))
}
