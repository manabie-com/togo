package repositories

import (
	"context"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/server"
)

type taskRepo struct {
	Database server.DataBase
}

type TaskRepo interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*models.Task, error)
	AddTask(ctx context.Context, t *models.Task) error
	ValidateUser(ctx context.Context, userID, pwd string) bool
	GetMaxTaskPerDay(ctx context.Context, userId string) (int, error)
	CheckTaskPerDayOfAnUser(ctx context.Context, maxToDo int, userId, createdDate string) bool
}

func NewTaskRepo(db server.DataBase) TaskRepo {
	return &taskRepo{Database: db}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *taskRepo) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*models.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := l.Database.Postgres.Query(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		t := &models.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (l *taskRepo) AddTask(ctx context.Context, t *models.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := l.Database.Postgres.Exec(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *taskRepo) ValidateUser(ctx context.Context, userID, pwd string) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := l.Database.Postgres.QueryRow(ctx, stmt, userID, pwd)
	u := &models.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}

func (l *taskRepo) GetMaxTaskPerDay(ctx context.Context, userId string) (int, error) {
	var maxToDo int
	stmt := `SELECT max_todo FROM users WHERE id = $1 `
	row := l.Database.Postgres.QueryRow(ctx, stmt, userId)
	err := row.Scan(&maxToDo)
	if err != nil {
		return 0, err
	}
	return maxToDo, nil
}

func (l *taskRepo) CheckTaskPerDayOfAnUser(ctx context.Context, maxToDo int, userId, createdDate string) bool {
	var totalCount int
	stmt := `SELECT count(id) FROM tasks WHERE user_id = $1 AND created_date = $2`
	err := l.Database.Postgres.QueryRow(ctx, stmt, userId, createdDate).Scan(&totalCount)
	if err != nil {
		return false
	}
	if totalCount >= maxToDo {
		return false
	}
	return true
}
