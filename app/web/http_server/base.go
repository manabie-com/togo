package http_server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/manabie-com/togo/app/common/config"
	gHandler "github.com/manabie-com/togo/app/common/gstuff/handler"

	userRepo "github.com/manabie-com/togo/app/repo/mongo/user"

	userHandler "github.com/manabie-com/togo/app/web/http_server/handler/user"
)

var cfg = config.GetConfig()

// Run start httpServer, init required elements
func Run() {
	// mongo client
	// init mongo
	cfg.Mongo.Get("app").Init()

	userRepoCollection := userRepo.InitColletion()
	userRepoInstance := userRepo.NewRepoManager(userRepoCollection)

	// init services
	userSrv := userHandler.NewService(userRepoInstance)

	// init api server
	server := NewAPIServer(
		userSrv,
	)

	// start server
	server.start()
}

type apiServer struct {
	userSrv userHandler.Service
}

// NewAPIServer : Tạo mới đối tuợng api server
func NewAPIServer(
	userSrv userHandler.Service,
) *apiServer {
	return &apiServer{
		userSrv: userSrv,
	}
}

// using to start Echo API Webserver
func (app *apiServer) start() {

	// init echo server
	e := echo.New()
	e.Validator = gHandler.NewValidator()
	e.HTTPErrorHandler = gHandler.EchoError

	// setup middlewares
	app.setMiddleware(e)

	// setup route
	app.initRoute(e)

	// start api server
	go func() {
		if err := e.Start(":80"); err != nil {
			log.Println("=> shutting down the server")
			log.Println(fmt.Sprintf("%v\n", err.Error()))
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Println("=> shutting down the server", err.Error())
	}

	return
}

// setup Middleware for API server
func (app *apiServer) setMiddleware(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))
}
