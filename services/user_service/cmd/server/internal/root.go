package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/phathdt/libs/go-sdk/httpserver/middleware"
	"user_service/common"
	"user_service/modules/usertransport/ginuser"

	"github.com/phathdt/libs/go-sdk/plugin/storage/sdkgorm"
	"github.com/phathdt/libs/go-sdk/plugin/storage/sdkredis"

	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/spf13/cobra"
)

var (
	serviceName = "user-service"
	version     = "1.0.0"
)

func newService() goservice.Service {
	s := goservice.New(
		goservice.WithName(serviceName),
		goservice.WithVersion(version),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
		goservice.WithInitRunnable(sdkredis.NewRedisDB("main", common.PluginRedis)),
	)

	return s
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start user service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recover())

			engine.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})

			users := engine.Group("/api/users")
			{
				users.POST("/register", ginuser.Register(service))
			}
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
