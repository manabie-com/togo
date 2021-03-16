package postgres

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/manabie-com/togo/internal/pkg/config"
	d "github.com/manabie-com/togo/internal/todo/domain"
	"github.com/pkg/errors"
)

type PGTaskRepository struct {
	PGRepository
}

func NewPGTaskRepository(dbConn *sqlx.DB) *PGTaskRepository {
	return &PGTaskRepository{PGRepository{dbConn}}
}

func (r *PGTaskRepository) CreateTaskForUser(ctx context.Context, userID int, taskParam d.TaskCreateParam) (*d.Task, error) {
	task := d.Task{UserID: userID, Content: taskParam.Content}
	tx, err := r.DBConn.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "db error")
	}

	// Lock row until transaction ends with FOR UPDATE
	if _, err := tx.ExecContext(ctx, "SELECT * FROM users WHERE id = $1 FOR UPDATE", userID); err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "db error")
	}

	// Insert new row when
	row := tx.QueryRowxContext(ctx,
		`INSERT INTO tasks (user_id, content) 
		SELECT $1, $2
		FROM tasks
		WHERE user_id = $1 AND created_at >= now()::date AND created_at < now()::date + '1 day'::interval
		HAVING count(id) < $3
		RETURNING *`,
		task.UserID, task.Content, config.GetEnvInt("MAX_TASKS_DAILY"))

	if err := row.StructScan(&task); err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, errors.Wrap(err, "db error")
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "db error")
	}

	return &task, nil
}

func (r *PGTaskRepository) GetTasksForUser(ctx context.Context, userID int, dateStr string) ([]*d.Task, error) {
	rows, err := r.DBConn.QueryxContext(
		ctx,
		"SELECT * FROM tasks WHERE user_id = $1 AND created_at >= $2 AND created_at <= $3",
		userID, dateStr+" 00:00:00", dateStr+" 23:59:59")
	if err != nil {
		return nil, errors.Wrap(err, "db error")
	}

	defer rows.Close()
	tasks := []*d.Task{}
	for rows.Next() {
		t := d.Task{}
		err := rows.StructScan(&t)
		if err != nil {
			return nil, errors.Wrap(err, "parse struct error")
		}

		tasks = append(tasks, &t)
	}

	return tasks, nil
}
