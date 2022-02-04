package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	sqlc "github.com/roandayne/togo/db/sqlc"
)

var db *sqlc.Queries
var task Task

type Task struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsComplete bool   `json:"is_complete"`
	FullName   string `json:"fullname"`
}

func (t Task) init() {
	conn, err := sql.Open("postgres", "user=postgres password=postgres dbname=todo_app sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = sqlc.New(conn)
}

func (t Task) GetUser() (sqlc.User, error) {
	un, err := db.GetUserByName(context.Background(), t.FullName)
	if err != nil {
		log.Println("Error: No user found")
		os.Exit(1)
	}

	return un, err
}

func (t Task) GetTasksCount(id int64) (int64, error) {
	tc, err := db.CountTasks(context.Background(), id)
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	return tc, err
}

func (t Task) CreateOneTask(id int64) {
	ts, err := db.CreateTask(context.Background(), sqlc.CreateTaskParams{
		Title:      t.Title,
		Content:    t.Content,
		IsComplete: t.IsComplete,
		UserID:     id,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ts)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	task.init()
	json.NewDecoder(r.Body).Decode(&task)

	un, err := task.GetUser()
	if err != nil {
		log.Fatal(err)
	}

	tc, err := task.GetTasksCount(int64(un.ID))
	if err != nil {
		log.Fatal(err)
	}

	if tc < int64(un.Maximum) {
		ts, err := db.CreateTask(context.Background(), sqlc.CreateTaskParams{
			Title:      task.Title,
			Content:    task.Content,
			IsComplete: task.IsComplete,
			UserID:     int64(un.ID),
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ts)
	} else {
		log.Println("You have reached the maximum allowed tasks for today! You can only add task count of ", un.Maximum)
		os.Exit(1)
	}
}
