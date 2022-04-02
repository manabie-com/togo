package http

import (
	"context"
	"fmt"
	"togo/internal/transport"
	"togo/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// httpServer http transport
type httpServer struct {
	echo        *echo.Echo
	// userService domain.UserService
	// authService domain.AuthService
	// taskService domain.TaskService
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
	taskUsecase usecase.TaskUsecase
}

// NewHTTPServer constructor
func NewHTTPServer(
	userUsecase usecase.UserUsecase,
	authUsecase usecase.AuthUsecase,
	taskUsecase usecase.TaskUsecase,
) transport.Server {
	return &httpServer{
		echo.New(),
		userUsecase,
		authUsecase,
		taskUsecase,
	}
}

func (s *httpServer) Load(_ context.Context) error {
	s.echo.Use(middleware.Logger())
	s.echo.Use(httpErrorResolver)
	s.echo.POST("/auth/register", s.Register)
	s.echo.POST("/auth/login", s.Login)
	s.echo.GET("/users/me", s.GetMe, s.authGuard)
	s.echo.PATCH("/users/me", s.UpdateMe, s.authGuard)
	s.echo.POST("/tasks", s.AddTask, s.authGuard)
	s.echo.GET("/tasks", s.GetTasks, s.authGuard)
	s.echo.PATCH("/tasks/:id", s.UpdateTask, s.authGuard)
	return nil
}

func (s *httpServer) Serve(host string, port int) {
	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf("%s:%v", host, port)))
}
