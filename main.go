package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/controllers"
	"github.com/manabie-com/togo/internal/pkg/middleware"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/server"
	"github.com/manabie-com/togo/internal/services"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	server.InitServerConfig("")
	server.Database.InitDatabase()
}

func main() {
	r := gin.New()
	taskRepo := repositories.NewTaskRepo(server.Database)

	toDoService := services.NewToDoService(taskRepo)
	c := controllers.New(toDoService)
	// cors
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET, PUT, POST, DELETE"},
		AllowHeaders:     []string{"Origin ,Cookie ,Authorization, Content-Type"},
		ExposeHeaders:    []string{""},
		AllowCredentials: true,
	}))
	r.Use(middleware.GinTokenMiddleware())
	r.POST("/login", c.Login)
	r.GET("/tasks", c.ListTask)
	r.POST("/tasks", c.AddTask)
	r.Run(fmt.Sprintf(":%d", server.Config.Port))
}
