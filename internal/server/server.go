package server

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie/project/internal/router"
	"github.com/manabie/project/middleware"
)

type server struct {
	router     router.Router
	middleware middleware.AccessController
	engine     *gin.Engine
}

type Server interface {
	RunServer() error
}

func NewServer(router router.Router, middleware middleware.AccessController, engine *gin.Engine) Server {
	return &server{
		router:     router,
		middleware: middleware,
		engine:     engine,
	}
}

func(s *server) RunServer() error {
	group := s.engine.Group("/api")

	group.POST("/login", s.router.Login).POST("/register", s.router.SignUp)

	groupTask := group.Group("/task")
	groupTask.POST("/:user_id", s.middleware.Authenticate(), s.router.CreateTask)
	groupTask.PUT("/:id", s.middleware.Authenticate(), s.router.UpdateTask)
	groupTask.DELETE("/:id", s.middleware.Authenticate(), s.router.DeleteTask)
	groupTask.GET("/", s.middleware.Authenticate(), s.router.TaskAll)
	groupTask.GET("/:id", s.middleware.Authenticate(), s.router.TaskById)

	return nil
}