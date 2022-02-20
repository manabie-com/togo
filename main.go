package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var Mode = os.Getenv("MODE")
var port = os.Getenv("PORT")
var dbHost = os.Getenv("DB_HOST")
var dbPort = os.Getenv("DB_PORT")
var dbDatabase = os.Getenv("DB_DATABASE")
var dbUsername = os.Getenv("DB_USERNAME")
var dbPassword = os.Getenv("DB_PASSWORD")

func init() {
	gin.SetMode(Mode)
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)
	fmt.Println(dbInfo)
}

func main() {
	router := gin.Default()
	router.GET("/health-check", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"service": "API Todo",
			"status":  1,
		})
	})

	router.Run(":" + port)
}
