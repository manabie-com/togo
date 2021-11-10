package postgresql

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/domain/entity"
)

type taskPostgresqlRepository struct {
	DB *sql.DB
}

// NewTaskPostgresqlRepository will create an implementation of task.TaskRepository
func NewTaskPostgresqlRepository(db *sql.DB) domain.TaskRepository {
	return taskPostgresqlRepository{
		DB: db,
	}
}

func (t taskPostgresqlRepository) Create(ctx context.Context, content string, username string) error {
	query := `INSERT INTO tasks(content, username) VALUES ($1, $2)`
	stmt, err := t.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, content, username)
	if err != nil {
		return err
	}
	return nil
}

func (t taskPostgresqlRepository) CountTaskInDayByUsername(ctx context.Context, username string) (int, error) {
	query := `SELECT count(*) FROM tasks WHERE username = $1 AND DATE(created_at) = DATE(NOW());`
	stmt, err := t.DB.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	var count int
	row := stmt.QueryRowContext(ctx, username)
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t taskPostgresqlRepository) GetTaskByUsernameAndDate(ctx context.Context, username string, date string) ([]entity.Task, error) {
	query := `SELECT id, content, username, created_at FROM tasks WHERE username = $1 AND DATE(created_at)=DATE($2)`
	stmt, err := t.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, username, date)
	if err != nil {
		return nil, err
	}
	result := make([]entity.Task, 0)
	for rows.Next() {
		task := entity.Task{}
		err := rows.Scan(
			&task.Id,
			&task.Content,
			&task.Username,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, task)
	}
	return result, nil
}
