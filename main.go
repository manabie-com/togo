package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/namnhatdoan/togo/handlers"
	"github.com/namnhatdoan/togo/settings"
	"log"
)

func main() {
	startRestApi()
}

func startRestApi() {
	router := gin.Default()
	api := router.Group("/")
	{
		api.POST("/tasks/", handlers.CreateTask)
		api.POST("/config/", handlers.SetConfig)
	}
	if err := router.Run(fmt.Sprintf(":%v", settings.RestPort)); err != nil {
		log.Fatal(err)
	}
}
