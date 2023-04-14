package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/phathdt/libs/go-sdk/httpserver/middleware"
	"github.com/phathdt/libs/go-sdk/plugin/appgrpc"
	"github.com/phathdt/libs/go-sdk/plugin/tokenprovider/jwt"
	"google.golang.org/grpc"
	"user_service/common"
	"user_service/modules/usertransport/ginuser"
	"user_service/modules/usertransport/usergrpc"

	"github.com/phathdt/libs/go-sdk/plugin/storage/sdkgorm"
	"github.com/phathdt/libs/go-sdk/plugin/storage/sdkredis"

	goservice "github.com/phathdt/libs/go-sdk"
	protos "github.com/phathdt/libs/togo_proto/out/proto"
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
		goservice.WithInitRunnable(jwt.NewJWTProvider(common.PluginJWT)),
		goservice.WithInitRunnable(appgrpc.NewGRPCServer(common.PluginGrpcServer)),
	)

	return s
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start user service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		serviceLogger := service.Logger("service")

		service.MustGet(common.PluginGrpcServer).(interface {
			SetRegisterHdl(hdl func(*grpc.Server))
		}).SetRegisterHdl(func(server *grpc.Server) {
			protos.RegisterUserServiceServer(server, usergrpc.NewUserGrpcServer(service))
		})

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
