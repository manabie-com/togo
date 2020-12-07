package cmd

import (
	userCache "github.com/HoangVyDuong/togo/internal/cache/user/redis"
	authHandler "github.com/HoangVyDuong/togo/internal/handler/auth"
	taskHandler "github.com/HoangVyDuong/togo/internal/handler/task"
	taskRepository "github.com/HoangVyDuong/togo/internal/storages/task/postgres"
	userRepository "github.com/HoangVyDuong/togo/internal/storages/user/postgres"
	authTransport "github.com/HoangVyDuong/togo/internal/transport/auth"
	taskTransport "github.com/HoangVyDuong/togo/internal/transport/task"
	authUsecase "github.com/HoangVyDuong/togo/internal/usecase/auth"
	taskUsecase "github.com/HoangVyDuong/togo/internal/usecase/task"
	userUsecase "github.com/HoangVyDuong/togo/internal/usecase/user"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"github.com/HoangVyDuong/togo/pkg/plugins/cache/redis"
	"github.com/HoangVyDuong/togo/pkg/plugins/database/postgres"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Infof("Starting with config: %v", viper.AllSettings())

		redisClient := redis.NewRedisClient()
		dbClient := postgres.NewDatabase()

		cache := userCache.NewCache(redisClient)

		taskRepository := taskRepository.NewRepository(dbClient)
		userRepository := userRepository.NewRepository(dbClient)

		userService := userUsecase.NewService(cache)
		taskService := taskUsecase.NewService(taskRepository)
		authService := authUsecase.NewService(userRepository)

		authHandler := authHandler.NewHandler(authService)
		taskHandler := taskHandler.NewHander(taskService)

		router := httprouter.New()
		authTransport.MakeHandler(router, authHandler)
		taskTransport.MakeHandler(router, taskHandler, userService, taskService)

		logger.Info("Starting Serve HTTP server")
		if err := http.ListenAndServe(viper.GetString("server.address"), router); err != nil {
			logger.Errorf("Exit by %s", err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
