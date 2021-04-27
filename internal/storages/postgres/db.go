package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/manabie-com/togo/internal/app/models"
)

// LiteDB for working with sqllite
type PGsqlImpl struct {
	DB *sql.DB
}

func NewPosgresql(DB *sql.DB) *PGsqlImpl {
	return &PGsqlImpl{DB}
}

// ValidateUser returns tasks if match userID AND password
func (pg *PGsqlImpl) ValidateUser(ctx context.Context, username string, pwd string) (*models.User, error) {
	stmt := `SELECT id, username, password FROM users WHERE username = $1 AND password = $2`
	row := pg.DB.QueryRowContext(ctx, stmt, username, pwd)
	u := &models.User{}
	err := row.Scan(&u.ID, &u.Username, &u.Password)

	if err != nil {
		return nil, err
	}

	return u, nil
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (pg *PGsqlImpl) RetrieveTasks(ctx context.Context, userID uint64, createdDate string) ([]*models.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date::date = $2::date`
	rows, err := pg.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		tsk := &models.Task{}
		err := rows.Scan(&tsk.ID, &tsk.Content, &tsk.UserID, &tsk.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, tsk)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (pg *PGsqlImpl) AddTask(ctx context.Context, t *models.Task) error {
	query := `INSERT INTO tasks (id, content, user_id, created_date) 
			SELECT $1, $2, $3, $4::date 
			FROM users u 
			WHERE u.id = $3 
				AND (SELECT COUNT(id) 
				FROM tasks 
				WHERE user_id = $3 
				AND created_date::date = $4::date) < u.max_todo RETURNING id`

	stmt, err := pg.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, t.ID, t.Content, t.UserID, t.CreatedDate)
	if err != nil {
		return err
	}

	rsAfted, _ := res.RowsAffected()

	if rsAfted < 1 {
		return errors.New("the task daily limit is reached")
	}

	return nil
}
