package postgresql

import (
	"context"
	"fmt"
	"github.com/banhquocdanh/togo/internal/storages"
	"github.com/go-pg/pg/v10"
)

type PostgreSql struct {
	db *pg.DB
}

func NewPostgreSQL(db *pg.DB) *PostgreSql {
	return &PostgreSql{db: db}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (s *PostgreSql) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	var tasks []storages.Task
	err := s.db.Model(&tasks).Where("user_id = ?", userID).Where("created_date = ?", createdDate).Select()

	if err != nil {
		return nil, err
	}

	result := make([]*storages.Task, 0, len(tasks))
	for _, t := range tasks {
		result = append(result, &t)
	}
	return result, err
}

// addTask adds a new task to DB
func (s *PostgreSql) addTask(ctx context.Context, t *storages.Task) error {
	_, err := s.db.Model(t).Insert()
	return err
}

// AddTask adds a new task to DB, with depend on user.MaxTodo
func (s *PostgreSql) AddTask(ctx context.Context, t *storages.Task) error {
	user := &storages.User{
		ID: t.UserID,
	}
	err := s.db.Model(user).WherePK().Select()
	if err != nil {
		return err
	}

	count, err := s.db.Model(&storages.Task{
		UserID:      t.UserID,
		CreatedDate: t.CreatedDate,
	}).Count()

	if uint(count) >= user.MaxTodo {
		return fmt.Errorf("max todo task")
	}

	return s.addTask(ctx, t)
}

// ValidateUser returns tasks if match userID AND password
func (s *PostgreSql) ValidateUser(ctx context.Context, userID, pwd string) bool {
	user := &storages.User{
		ID: userID,
	}
	err := s.db.Model(user).WherePK().Select()
	if err != nil {
		return false
	}

	return user.Password == pwd
}
