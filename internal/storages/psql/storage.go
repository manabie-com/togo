//Package psql
package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/domain"
)

type Storage struct {
	db *sql.DB
}

type Config struct {
	ConnString string
}

func NewStorage(c Config) (*Storage, error) {
	db, err := sql.Open("postgres", c.ConnString)
	if err != nil {
		return nil, err
	}
	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) AddTaskWithLimitPerDay(task domain.Task, limit int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	rows, err := tx.Query("SELECT id FROM tasks where tasks.user_id =$1 and date(tasks.created_date)= current_date for update", task.UserID)
	if err != nil {
		fmt.Println(err)
	}
	result := 0
	for rows.Next() {
		result++
		rows.Scan(&struct{}{})
	}
	if result >= limit {
		return domain.TaskLimitReached{}
	}
	_, err = tx.Exec(`INSERT INTO public.tasks(id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`, task.ID, task.Content, task.UserID, task.CreatedDate)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *Storage) GetTasksByUserID(userID string, offset, limit int) ([]domain.Task, error) {
	rows, err := s.db.Query("SELECT id,content,user_id,created_date FROM tasks where tasks.user_id =$1 limit $2 offset $3", userID, limit, offset)
	if err != nil {
		return nil, err
	}
	result := []domain.Task{}

	for rows.Next() {
		var t domain.Task
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (s *Storage) FindUserByID(string) (domain.User, error) {
	return domain.User{}, nil
}
func (s *Storage) CreateUser(Id string, Password string) error   { return nil }
func (s *Storage) GetUserTasksPerDay(userID string) (int, error) { return 0, nil }
