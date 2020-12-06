package cmd

import (
	"database/sql"
	"fmt"
	authHandler "github.com/HoangVyDuong/togo/internal/handler/auth"
	taskHandler "github.com/HoangVyDuong/togo/internal/handler/task"
	authUsecase "github.com/HoangVyDuong/togo/internal/usecase/auth"
	taskUsecase "github.com/HoangVyDuong/togo/internal/usecase/task"
	userUsecase "github.com/HoangVyDuong/togo/internal/usecase/user"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	taskSQLite "github.com/HoangVyDuong/togo/internal/storages/task/sqlite"
	userSQLite "github.com/HoangVyDuong/togo/internal/storages/user/sqlite"
	authTransport "github.com/HoangVyDuong/togo/internal/transport/auth"
	taskTransport "github.com/HoangVyDuong/togo/internal/transport/task"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"log"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "./data.db")
		if err != nil {
			log.Fatal("error opening db", err)
		}

		taskRepository := &taskSQLite.LiteDB{DB: db}
		userRepository := &userSQLite.LiteDB{DB: db}

		userService := userUsecase.NewService(userRepository)
		taskService := taskUsecase.NewService(taskRepository)
		authService := authUsecase.NewService()

		authHandler := authHandler.NewHander(authService, userService)
		taskHandler := taskHandler.NewHander(taskService)

		router := httprouter.New()
		authTransport.WithHandler(router, authHandler)
		taskTransport.WithHandler(router, taskHandler)


		errs := make(chan error)
		go func() {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			errs <- fmt.Errorf("%s", <-c)
		}()


		errs <- http.ListenAndServe(viper.GetString("server.address"), router)
		log.Print("exit", <- errs)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
