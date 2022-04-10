package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/namnhatdoan/togo/handlers"
	"github.com/namnhatdoan/togo/services"
	"github.com/namnhatdoan/togo/settings"
	"log"
)

func main() {
	startRestApi()
}

func startRestApi() {
	handler := handlers.InitToGoHandler(&services.ToGoServiceImpl{})

	router := gin.Default()
	api := router.Group("/")
	{
		api.POST("/tasks/", handler.CreateTask)
		api.POST("/config/", handler.SetConfig)
	}
	if err := router.Run(fmt.Sprintf(":%v", settings.RestPort)); err != nil {
		log.Fatal(err)
	}
}
