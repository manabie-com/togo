package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/httpserver/middleware"
	"github.com/phathdt/libs/go-sdk/plugin/storage/sdkgorm"
	"github.com/phathdt/libs/go-sdk/plugin/storage/sdkredis"
	"github.com/phathdt/libs/go-sdk/plugin/tokenprovider/jwt"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"togo/common"
	"togo/modules/task/tasktransport/taskgin"
	"togo/modules/user/userstorage"
	"togo/modules/user/usertransport/ginuser"
	middleware2 "togo/plugin/middleware"
)

var (
	serviceName = "task-service"
	version     = "1.0.0"
)

func newService() goservice.Service {
	s := goservice.New(
		goservice.WithName(serviceName),
		goservice.WithVersion(version),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
		goservice.WithInitRunnable(sdkredis.NewRedisDB("main", common.PluginRedis)),
		goservice.WithInitRunnable(jwt.NewJWTProvider(common.PluginJWT)),
	)

	return s
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start task service",
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
				users.POST("/login", ginuser.Login(service))
			}

			db := service.MustGet(common.DBMain).(*gorm.DB)
			store := userstorage.NewSQLStore(db)
			middlewareAuth := middleware2.RequireAuth(store, service)

			todos := engine.Group("/api/tasks", middlewareAuth)
			{
				todos.PATCH("/:id", taskgin.UpdateTask(service))
				todos.DELETE("/:id", taskgin.DeleteTask(service))
				todos.POST("", taskgin.CreateTask(service))
				todos.GET("", taskgin.ListTasks(service))
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
