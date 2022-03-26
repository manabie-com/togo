package postgres

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repository"
)

// postgresDB is a wrapper of *sql.DB for working with PostgreSQL
type postgresDB struct {
	db *sql.DB
}

// NewPostgresRepository returns a database repository with attached methods to interact directly with Postgres database
func NewPostgresRepository(db *sql.DB) repository.DatabaseRepository {
	return &postgresDB{db}
}

// ValidateUser checks if username and password is valid
func (p *postgresDB) ValidateUser(ctx context.Context, username, password sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE username = $1 AND password = $2`
	row := p.db.QueryRowContext(ctx, stmt, username, password)

	u := &models.User{}
	err := row.Scan(&u.ID)

	return err == nil
}

// GetUserByUserName gets user by the specified username
func (p *postgresDB) GetUserByUserName(ctx context.Context, username sql.NullString) (*models.User, error) {
	stmt := `SELECT id, username, max_task_per_day FROM users WHERE username = $1`
	row := p.db.QueryRowContext(ctx, stmt, username)

	u := &models.User{}
	if err := row.Scan(&u.ID, &u.Username, &u.MaxTaskPerDay); err != nil {
		return nil, err
	}

	return u, nil
}

// AddTask creates new task and persist data into DB
func (p *postgresDB) AddTask(ctx context.Context, task *models.Task) error {
	stmt := `INSERT INTO tasks (id, detail, user_id, created_date) VALUES ($1, $2, $3, $4)`
	if _, err := p.db.ExecContext(ctx, stmt, &task.ID, &task.Detail, &task.UserID, &task.CreatedDate); err != nil {
		return err
	}

	return nil
}

// RetrieveTasks return tasks by the specified userID and createdDate
func (p *postgresDB) RetrieveTasks(ctx context.Context, userID uint, createdDate sql.NullString) ([]*models.Task, error) {
	stmt := `SELECT id, detail, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := p.db.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*models.Task{}
	for rows.Next() {
		t := &models.Task{}
		if err := rows.Scan(&t.ID, &t.Detail, &t.UserID, &t.CreatedDate); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
