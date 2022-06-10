package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"togo/handler"
	logger2 "togo/pkg/logger"
	"togo/pkg/mysql"
	"togo/server"
)

var envFlag string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Your Server",
	Long:  "Let start a server with your opinion",
	Run: func(cmd *cobra.Command, args []string) {
		loggerPkg := logger2.NewLogger("Logger", "logs.txt")
		if err := loggerPkg.Run(); err != nil {
			log.Fatal(err)
		}

		logger := loggerPkg.Get()

		start, _ := cmd.Flags().GetBool("start")
		if start {
			svr := server.NewServer("Traning Go", 8080)
			restHdl := handler.RestHandler(svr)
			sql := mysql.NewMySQL("MYSQL")

			svr.InitService(sql)
			svr.AddHandler(restHdl)
			svr.AddLogger(logger)

			if err := svr.Run(); err != nil {
				logger.Error().Println(fmt.Sprintf("Server is stopped by %v", err.Error()))
			}
		}

		envVal, _ := cmd.Flags().GetBool("env-list")
		if envVal {
			logger.Info().Println(listEnv())
		}
	},
}

func listEnv() []string {
	return os.Environ()
}

func InitFlags() {
	serverCmd.PersistentFlags().Bool("env-list", false, "Usage to list service's env")
	serverCmd.PersistentFlags().Bool("start", false, "Command to start service with default port 8080")
}
