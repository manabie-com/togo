package cmd

import (
	"log"
	"os"

	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/plugin/storage/sdkgorm"
	"github.com/phathdt/libs/go-sdk/plugin/storage/sdkredis"
	"github.com/phathdt/libs/go-sdk/plugin/tokenprovider/jwt"
	"github.com/phathdt/libs/togo_appgrpc"
	"github.com/spf13/cobra"
	"togo/cmd/server/internal/handlers"
	"togo/common"
)

var (
	serviceName = "task-service"
	version     = "1.0.0"
)

func NewService() goservice.Service {
	s := goservice.New(
		goservice.WithName(serviceName),
		goservice.WithVersion(version),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
		goservice.WithInitRunnable(sdkredis.NewRedisDB("main", common.PluginRedis)),
		goservice.WithInitRunnable(jwt.NewJWTProvider(common.PluginJWT)),
		goservice.WithInitRunnable(togo_appgrpc.NewUserClient(common.PluginGrpcUserClient)),
	)

	return s
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start task service",
	Run: func(cmd *cobra.Command, args []string) {
		service := NewService()

		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		service.HTTPServer().AddHandler(handlers.NewHandlers(service))

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
