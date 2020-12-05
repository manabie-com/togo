package cmd

import (
	"database/sql"
	"github.com/HoangVyDuong/togo/internal/usecase/task"

	sqllite "github.com/HoangVyDuong/togo/internal/storages/task/sqlite"
	"github.com/HoangVyDuong/togo/internal/transport/handler"
	"github.com/HoangVyDuong/togo/internal/usecase/auth"

	"github.com/spf13/cobra"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "./data.db")
		if err != nil {
			log.Fatal("error opening db", err)
		}

		http.ListenAndServe(":5050", &handler.ToDoHandler{
			JWTKey: "wqGyEBBfPK9w3Lxw",
			AuthService: auth.NewService(&sqllite.LiteDB{DB: db}),
			TaskService: task.NewService(&sqllite.LiteDB{DB: db}),
		})
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
