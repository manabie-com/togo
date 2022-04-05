package main

import (
	"fmt"
	"os"
	"togo/internal/app/router"
	"togo/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	ginMode := gin.DebugMode

	dbConfig := database.DBConfig{
		Host: os.Getenv("DB_HOST"),
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Port: os.Getenv("DB_PORT"),
	}
	dbConn, err := database.NewDatabase(dbConfig)

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		database.Close(dbConn)
	}()

	gin.SetMode(ginMode)
	engine := gin.New()

	r := &router.Router{
		Engine: engine,
		DB:     dbConn.DB,
	}

	r.InitRoute()
	engine.Run(":" + port)
}
