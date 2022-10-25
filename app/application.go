package app

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/backend/controller/tasks"
)

type Server struct {
	Router     *gin.Engine
	Controller tasks.I_TaskController
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewServer(ctl tasks.I_TaskController) (*Server, error) {
	server := &Server{
		Controller: ctl,
	}
	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.Use(CORS())

	router.GET("/tasks", s.Controller.GetTaskAll)
	router.POST("/tasks", s.Controller.CreateTask)
	router.PATCH("/tasks/:id", s.Controller.UpdateTask)
	router.DELETE("/tasks/:id", s.Controller.DeleteTask)

	s.Router = router
}

func (s *Server) Start(address string) error {
	return s.Router.Run(address)
}
