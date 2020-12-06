package cmd

import (
	"database/sql"
	userCache "github.com/HoangVyDuong/togo/internal/cache/user/redis"
	authHandler "github.com/HoangVyDuong/togo/internal/handler/auth"
	taskHandler "github.com/HoangVyDuong/togo/internal/handler/task"
	taskRepository "github.com/HoangVyDuong/togo/internal/storages/task/sqlite"
	userRepository "github.com/HoangVyDuong/togo/internal/storages/user/sqlite"
	authTransport "github.com/HoangVyDuong/togo/internal/transport/auth"
	taskTransport "github.com/HoangVyDuong/togo/internal/transport/task"
	authUsecase "github.com/HoangVyDuong/togo/internal/usecase/auth"
	taskUsecase "github.com/HoangVyDuong/togo/internal/usecase/task"
	userUsecase "github.com/HoangVyDuong/togo/internal/usecase/user"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Starting")
		db, err := sql.Open("sqlite3", "./data.db")
		if err != nil {
			log.Fatal("error opening db", err)
		}

		redisClient := redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.address"),
			PoolSize: viper.GetInt("redis.connection_pool"),
		})

		cache := userCache.NewCache(redisClient)

		taskRepository := taskRepository.NewRepository(db)
		userRepository := userRepository.NewRepository(db)

		userService := userUsecase.NewService(userRepository, cache)
		taskService := taskUsecase.NewService(taskRepository)
		authService := authUsecase.NewService()

		authHandler := authHandler.NewHander(authService, userService)
		taskHandler := taskHandler.NewHander(taskService)

		router := httprouter.New()
		authTransport.MakeHandler(router, authHandler)
		taskTransport.MakeHandler(router, taskHandler, userService, taskService)

		if err := http.ListenAndServe(viper.GetString("server.address"), router); err != nil {
			logger.Errorf("Exit by %s", err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
