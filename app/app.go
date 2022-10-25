package main

import (
	gormPkg "ansidev.xyz/pkg/gorm"
	"ansidev.xyz/pkg/log"
	"ansidev.xyz/pkg/rds"
	"context"
	"database/sql"
	"fmt"
	authController "github.com/ansidev/togo/auth/controller/http"
	"github.com/ansidev/togo/config"
	"github.com/ansidev/togo/constant"
	"github.com/ansidev/togo/gingo"
	taskController "github.com/ansidev/togo/task/controller/http"
	"github.com/ansidev/togo/wire"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	appEnv string
	sqlDb  *sql.DB
	gormDb *gorm.DB
	rdb    *rds.RedisDB
)

func init() {
	log.InitLogger("console")

	appEnv = os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = constant.DefaultProdEnv
	}

	if appEnv == constant.DefaultProdEnv {
		config.LoadConfig("/app", constant.DefaultProdConfig, &config.AppConfig)
	} else {
		config.LoadConfig(".", constant.DefaultDevConfig, &config.AppConfig)
	}
}

func main() {
	// Flush log buffer if necessary
	defer log.Sync()

	if appEnv == constant.DefaultProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gingo.DefaultRouter()

	// Default route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"name":        constant.AppName,
			"version":     constant.AppVersion,
			"releaseDate": constant.AppReleaseDate,
		})
	})

	initInfrastructureServices()
	initControllers(router)

	server := &http.Server{
		Addr:    initAddress(),
		Handler: router,
	}

	// Listen from a different goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create channel to signify a signal being sent
	exit := make(chan os.Signal, 1)
	// When an interrupt or termination signal is sent, notify the channel
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	// Block the main thread until an interrupt is received
	<-exit
	log.Info("Gracefully shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

func initInfrastructureServices() {
	sqlDb = wire.InitSqlClient(config.AppConfig.SqlDbConfig)
	redisClient := wire.InitRedisClient(config.AppConfig.RedisConfig)
	dialector := postgres.New(postgres.Config{
		Conn:                 sqlDb,
		PreferSimpleProtocol: true,
	})
	gormDb = gormPkg.InitGormDb(dialector)
	rdb = rds.NewRedisDB(context.Background(), redisClient)
}

func initControllers(router *gin.Engine) {
	authService := wire.InitAuthService(gormDb, rdb, time.Duration(config.AppConfig.TokenTTL)*time.Second)
	authController.NewAuthController(router, authService)

	taskService := wire.InitTaskService(gormDb)
	taskController.NewTaskController(router, authService, taskService)
}

func initAddress() string {
	return fmt.Sprintf("%s:%d", config.AppConfig.Host, config.AppConfig.Port)
}
