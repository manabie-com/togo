package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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

func (t Task) init() (*sqlc.Queries, error) {
	conn, err := sql.Open("postgres", "user=postgres password=postgres dbname=todo_app sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = sqlc.New(conn)

	return db, err
}

func (t Task) GetUser() (sqlc.User, error) {
	un, err := db.GetUserByName(context.Background(), t.FullName)
	if err != nil {
		log.Fatal("Error: No user found")
	}

	return un, err
}

func (t Task) GetTasksCount(id int64) (int64, error) {
	tc, err := db.CountTasks(context.Background(), id)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return tc, err
}

func (t Task) CreateOneTask(id int64) (sqlc.Task, error) {
	ts, err := db.CreateTask(context.Background(), sqlc.CreateTaskParams{
		Title:      t.Title,
		Content:    t.Content,
		IsComplete: t.IsComplete,
		UserID:     id,
	})
	if err != nil {
		log.Fatal(err)
	}

	return ts, err
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	_, err := task.init()
	json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Fatal(err)
	}

	un, err := task.GetUser()
	if err != nil {
		log.Fatal(err)
	}
	tc, err := task.GetTasksCount(int64(un.ID))
	if err != nil {
		log.Fatal(err)
	}

	if tc < int64(un.Maximum) {
		_, err := task.CreateOneTask(int64(un.ID))
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(&task)
	} else {
		w.Write([]byte("You have reached the maximum allowed tasks for today!"))
	}
}
