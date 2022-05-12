package task

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	domain "github.com/nvhai245/togo/internal/domain"

	_ "github.com/mattn/go-sqlite3"
)

type repository struct {
	db *sql.DB
}

// NewRepository returns a new domain
func NewRepository(dialect, dsn string, idleConn, maxConn int) domain.TaskRepository {
	db, err := sql.Open(dialect, dsn)
	if err != nil {
		log.Println("error connect to database")
		return nil
	}
	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	migrateFilePath, err := filepath.Abs("script/seed.sql")
	if err != nil {
		panic(err)
	}
	query, err := ioutil.ReadFile(migrateFilePath)
	if err != nil {
		panic(err)
	}
	if _, err = db.Exec(string(query)); err != nil {
		panic(err)
	}

	return &repository{db}
}

// Close closes the connection
func (r *repository) Close() {
	r.db.Close()
}

// CheckTaskByUserID checks whether a user can create a task
func (r *repository) CheckTaskByUserID(ctx context.Context, userID int64) error {
	var (
		limit int
		count int
	)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, "SELECT daily_limit FROM users WHERE id = ?", userID).Scan(&limit)
	if err != nil {
		return errors.New("user does not exist")
	}
	err = r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM tasks WHERE user_id = ? AND (created_at BETWEEN Date('now') AND Date('now','+1 day','-0.001 second'))", userID).Scan(&count)
	if err != nil {
		return err
	}
	if count >= limit {
		return errors.New(fmt.Sprintf("exceeded daily %v tasks limit", limit))
	}
	return nil
}

// CreateTask creates a task for user
func (r *repository) CreateTask(ctx context.Context, userID int64, content string) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	task := new(domain.Task)
	query := "INSERT INTO tasks (user_id, content) VALUES (?, ?) RETURNING *"

	err := r.db.QueryRowContext(ctx, query, userID, content).Scan(&task.ID, &task.UserID, &task.Content, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// GetAllTaskByUserID get all tasks of an user
func (r *repository) GetAllTaskByUserID(ctx context.Context, userID int64) ([]*domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM tasks WHERE user_id = ? ORDER BY created_at", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*domain.Task
	for rows.Next() {
		task := &domain.Task{}
		err = rows.Scan(&task.ID, &task.UserID, &task.Content, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, task)
	}

	return results, nil
}
