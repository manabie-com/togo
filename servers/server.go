package servers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/triet-truong/todo/config"
	"github.com/triet-truong/todo/domain"
	"github.com/triet-truong/todo/todo/controller/http_handlers"
	"github.com/triet-truong/todo/todo/usecase"
)

type Server struct {
	e         *echo.Echo
	dbRepo    domain.TodoRepository
	cacheRepo domain.TodoCacheRepository
}

func NewServer(dbRepo domain.TodoRepository, cacheRepo domain.TodoCacheRepository) *Server {
	return &Server{
		dbRepo:    dbRepo,
		cacheRepo: cacheRepo,
	}
}

func (s *Server) Run() {
	e := echo.New()
	e.GET("/hello", func(ctx echo.Context) error {
		resp := `{"message":"success"}`
		ctx.Response().Write(bytes.NewBufferString(resp).Bytes())
		ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return nil
	})

	e.Debug = true

	usecase := usecase.NewTodoUseCase(s.dbRepo, s.cacheRepo)
	handler := http_handlers.NewTodoHandler(usecase)
	e.POST("/user/todo", handler.Add)

	// Start server
	// go func() {
	if err := e.Start(fmt.Sprintf(":%v", config.Port())); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server")
	}
	// }()

	s.e = e
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// if err := e.Shutdown(ctx); err != nil {
	// 	utils.FatalLog(err)
	// }
}

func (s *Server) Stop() {
	s.e.Close()
}
