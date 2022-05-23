package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"togo/domain/service"
	"togo/infrastructure/inmemory"
	httpInterface "togo/interface/http"
)

func main() {
	userRepo := inmemory.NewInMemoryUserRepo()
	userService := service.NewUserService(userRepo)
	userHttpController := httpInterface.NewUserController(userService)
	e := gin.New()
	userGroup := e.Group("/users")
	userGroup.POST("/register", userHttpController.Register)
	e.Run(fmt.Sprintf("0.0.0.0:%d", 8080))
}
