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

//NewStorage return new psql storage
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

func (s *Storage) FindUserByID(userID string) (domain.User, error) {
	rows, err := s.db.Query("SELECT id,password,max_todo FROM tasks where tasks.user_id =$1", userID)
	empty := domain.User{}
	if err != nil {
		return empty, err
	}
	if !rows.Next() {
		return empty, domain.UserNotFound(userID)
	}
	err = rows.Scan(&empty.ID, &empty.Password, &empty.MaxTasksPerDay)
	if err != nil {
		return empty, err
	}
	return empty, nil
}

func (s *Storage) CreateUser(user domain.User) error {
	_, err := s.db.Exec("INSERT INTO users(id, password, max_todo) VALUES ($1,$2,$3)", user.ID, user.Password, user.MaxTasksPerDay)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetUserTasksPerDay(userID string) (int, error) {
	rows, err := s.db.Query("SELECT max_todo FROM tasks where tasks.user_id =$1", userID)
	if err != nil {
		return 0, err
	}
	if !rows.Next() {
		return 0, domain.UserNotFound(userID)
	}
	var result int
	err = rows.Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
