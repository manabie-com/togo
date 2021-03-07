package integration_tests

import (
	"database/sql"
	"github.com/manabie-com/togo/internal/model"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"testing"
	"time"
)
var (
	dbConn  *sql.DB
)
func TestMain(m *testing.M) {
	var err error
	dbConn, err = sql.Open("sqlite3", "./data_test.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	os.Exit(m.Run())
}

func refreshTasksTable() error {
	stmt, err := dbConn.Prepare("DELETE FROM tasks;")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("Error delete tasks table: %s", err)
	}
	return nil
}

func seedOneTask() (model.Task, error) {
	task := model.Task{
		ID:     "1",
		Content:      "Task 1",
		UserID: "firstUser",
		CreatedDate: time.Now().Format("2006-01-02"),
	}
	stmt, err := dbConn.Prepare("INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	_, createErr := stmt.Exec(task.ID, task.Content, task.UserID, task.CreatedDate)
	if createErr != nil {
		log.Fatalf("Error creating message: %s", createErr)
	}

	return task, nil
}

func seedTasks() ([]model.Task, error) {
	tasks := []model.Task{
		{
			ID:     "2",
			Content:      "Task 2",
			UserID: "firstUser",
			CreatedDate: time.Now().Format("2006-01-02"),
		},
		{
			ID:     "3",
			Content:      "Task 3",
			UserID: "firstUser",
			CreatedDate: time.Now().Format("2006-01-02"),
		},
	}

	for i := range tasks {
		stmt, err := dbConn.Prepare("INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)")
		if err != nil {
			return nil, err
		}
		stmt.QueryRow(tasks[i].ID, tasks[i].Content, tasks[i].UserID, tasks[i].CreatedDate)

	}

	return tasks, nil
}