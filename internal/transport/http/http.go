package http

import (
	"context"
	"fmt"
	"togo/internal/domain"
	"togo/internal/transport"

	"github.com/labstack/echo/v4"
)

// httpServer http transport
type httpServer struct {
	echo        *echo.Echo
	userService domain.UserService
	authService domain.AuthService
}

// NewHTTPServer constructor
func NewHTTPServer(userService domain.UserService, authService domain.AuthService) transport.Server {
	return &httpServer{
		echo.New(),
		userService,
		authService,
	}
}

func (s *httpServer) Load(_ context.Context) error {
	s.echo.GET("/users/:id", s.GetUser)
	s.echo.POST("/auth/register", s.Register)
	s.echo.POST("/auth/login", s.Login)
	return nil
}

func (s *httpServer) Serve(host string, port int) {
	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf("%s:%v", host, port)))
}
