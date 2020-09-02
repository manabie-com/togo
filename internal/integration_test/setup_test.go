package integration_test

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

var toDoService *services.ToDoService

func TestMain(m *testing.M) {
	os.Exit(createServices(m))
}

func createServices(m *testing.M) int {
	db, err := sql.Open("sqlite3", "../../data.db")
	if err != nil {
		log.Fatalf("error opening db %v", err)
	}
	toDoService = &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}
	return m.Run()
}

func truncate(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt := `DELETE FROM users`
	if _, err := tx.Exec(stmt); err != nil {
		tx.Rollback()
		return err
	}

	stmt = `DELETE FROM tasks`
	if _, err := tx.Exec(stmt); err != nil {
		tx.Rollback()
		return err
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

	stmt := `INSERT INTO users (id, password) VALUES (?, ?)`

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
		stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
		_, err := db.Exec(stmt, &task.ID, &task.Content, &task.UserID, &task.CreatedDate)
		if err != nil {
			return nil, err
		}
	}

	return taskItems, nil
}
