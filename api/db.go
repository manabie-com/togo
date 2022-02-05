package api

import (
	"context"
	"database/sql"
	"log"
	"os"

	sqlc "github.com/roandayne/togo/db/sqlc"
)

var db *sqlc.Queries
var task Task

type Task struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsComplete bool   `json:"is_complete"`
	Username   string `json:"username"`
}

func Initialize() {
	database_url := os.Getenv("DATABASE_URL")
	conn, err := sql.Open("postgres", database_url)
	if err != nil {
		log.Fatal(err)
	}
	db = sqlc.New(conn)
}

func (t Task) GetUser() (sqlc.User, error) {
	u, err := db.GetUserByName(context.Background(), t.Username)

	return u, err
}

func (t Task) GetTasksCount(id int64) int64 {
	tc, err := db.CountUserTasks(context.Background(), id)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return tc
}

func (t Task) CreateOneTask(id int64) sqlc.Task {
	ts, err := db.CreateTask(context.Background(), sqlc.CreateTaskParams{
		Title:      t.Title,
		Content:    t.Content,
		IsComplete: t.IsComplete,
		UserID:     id,
	})
	if err != nil {
		log.Fatal(err)
	}

	return ts
}
