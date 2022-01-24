package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"togo/handler"
	"togo/pkg/mysql"
	"togo/server"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Start a server",
	Run: func(cmd *cobra.Command, args []string) {
		svr := server.NewServer("Traning Go", 8080)
		restHdl := handler.RestHandler(svr)
		sql := mysql.NewMySQL("MYSQL")

		svr.InitService(sql)
		svr.AddHandler(restHdl)
		if err := svr.Run(); err != nil {
			log.Printf("Server is stopped by %v\n", err.Error())
		}
	},
}