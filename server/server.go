package server

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/controllers"
)

type Server struct {
	Router     *gin.Engine
	Controller controllers.ITaskController
}

func NewServer(taskController controllers.ITaskController) (*Server, error) {
	server := &Server{
		Controller: taskController,
	}

	server.configRouter()

	return server, nil
}

func (server *Server) StartServer(address string) error {
	return server.Router.Run(address)
}

func (server *Server) configRouter() {
	router := gin.Default()

	router.GET("/tasks", server.Controller.GetAllTask)
	router.POST("/tasks", server.Controller.CreateTask)

	server.Router = router
}
