package integration_test

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
)

var toDoService services.ToDoService
var DB *sql.DB

func TestMain(m *testing.M) {
	os.Exit(createServices(m))
}

func createServices(m *testing.M) int {

	dbconnecter, err := config.GetDBConnecter()
	if err != nil {
		log.Println(err)
		return 0
	}

	db, err := dbconnecter.Connect()
	if err != nil {
		log.Fatal("error opening db", err)
	}
	toDoService = services.ToDoService{
		Router: mux.NewRouter(),
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.DBStore{
			DB: db,
		},
	}

	toDoService.Router.HandleFunc("/login", toDoService.LoginHandler)
	toDoService.Router.HandleFunc("/tasks", toDoService.Validate(toDoService.GetTasksHandler)).Methods(http.MethodGet)
	toDoService.Router.HandleFunc("/tasks", toDoService.Validate(toDoService.CreateTaskHandler)).Methods(http.MethodPost)
	DB = db
	return m.Run()
}

func refreshUser(db *sql.DB, user *storages.User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt := `DELETE FROM tasks WHERE user_id=$1`
	if _, err := tx.Exec(stmt, user.ID); err != nil {
		tx.Rollback()
		return err
	}

	stmt = `DELETE FROM users WHERE id=$1`
	if _, err := tx.Exec(stmt, user.ID); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func refreshSingleTask(db *sql.DB, task *storages.Task) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt := `DELETE FROM tasks WHERE id=$1`
	if _, err := tx.Exec(stmt, task.ID); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func refreshTasks(db *sql.DB, tasks []storages.Task) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		stmt := `DELETE FROM tasks WHERE id=$1`
		if _, err := tx.Exec(stmt, task.ID); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func seedUser(db *sql.DB) (*storages.User, error) {
	u := storages.User{
		ID:       "UserTest1",
		Password: "example",
	}

	stmt := `INSERT INTO users (id, password) VALUES ($1, $2)`

	_, err := db.Exec(stmt, &u.ID, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func seedTaskItems(db *sql.DB, u *storages.User) ([]storages.Task, error) {
	now := time.Now()
	taskItems := []storages.Task{
		{
			ID:          uuid.New().String(),
			Content:     "write content here",
			UserID:      u.ID,
			CreatedDate: now.Format("2006-01-02"),
		},
	}

	for _, task := range taskItems {
		stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
		_, err := db.Exec(stmt, &task.ID, &task.Content, &task.UserID, &task.CreatedDate)
		if err != nil {
			return nil, err
		}
	}

	return taskItems, nil
}
