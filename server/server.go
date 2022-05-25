package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/component"
	"github.com/japananh/togo/component/tokenprovider"
	"github.com/japananh/togo/middleware"
	"github.com/japananh/togo/modules/task/tasktransport/gintask"
	"github.com/japananh/togo/modules/user/usertransport/ginuser"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// Server represents server
type Server struct {
	Port        int
	AppEnv      string
	SecretKey   string
	DBConn      *gorm.DB
	TokenConfig *tokenprovider.TokenConfig
	ServerReady chan bool
}

// Start start http server
func (s *Server) Start() {
	// Create context that listens for the interrupt signal from the OS.
	// Reference: https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	if s.AppEnv == "dev" {
		gin.SetMode(gin.DebugMode)
		r.Use(gin.Logger())
	}

	appCtx := component.NewAppContext(s.DBConn, s.SecretKey, s.TokenConfig)
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/api/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))

	tasks := v1.Group("/tasks", middleware.RequiredAuth(appCtx))
	{
		tasks.POST("/", gintask.CreateTask(appCtx))
	}

	// TODO: How to only show these APIs in development?
	// api for encode uid receives real id and database type, then return fake uid
	// e.g: id: 16, db_type: 2 -> fakeId: 3w5rMJ8raFkfXS
	v1.GET("/encode-uid", func(c *gin.Context) {
		type reqData struct {
			DBType int `form:"db_type" binding:"required"`
			RealId int `form:"id" binding:"required"`
		}

		var d reqData
		if err := c.ShouldBind(&d); err != nil {
			c.JSON(http.StatusBadRequest, "invalid request")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": common.NewUID(uint32(d.RealId), d.DBType, 1),
		})
	})

	// api for decode uid receives fake uid then return real id and database type
	// e.g: fakeId: 3w5rMJ8raFkfXS -> id: 16, db_type: 2
	v1.GET("/decode-uid", func(c *gin.Context) {
		type reqData struct {
			FakeId string `form:"id" binding:"required"`
		}

		var d reqData
		if err := c.ShouldBind(&d); err != nil {
			c.JSON(http.StatusBadRequest, "invalid request")
			return
		}

		realId, err := common.FromBase58(d.FakeId)
		if err != nil {
			c.JSON(http.StatusBadRequest, "invalid request")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":      realId.GetLocalID(),
			"db_type": realId.GetObjectType(),
		})
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Printf("Server run on PORT :%d\n", s.Port)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	if s.ServerReady != nil {
		s.ServerReady <- true
	}

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
