package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	d "github.com/manabie-com/togo/internal/todo/domain"
	"github.com/pkg/errors"
)

type PGTaskRepository struct {
	PGRepository
}

func NewPGTaskRepository(dbConn *sqlx.DB) *PGTaskRepository {
	return &PGTaskRepository{PGRepository{dbConn}}
}

func (t *PGTaskRepository) CreateTaskForUser(ctx context.Context, userID int, taskParam d.TaskCreateParam) (*d.Task, error) {
	task := d.Task{UserID: userID, Content: taskParam.Content}
	_, err := t.DBConn.NamedExecContext(
		ctx,
		"INSERT INTO tasks (user_id, content) VALUES (:user_id, :content)",
		task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *PGTaskRepository) GetTasksForUser(ctx context.Context, userID int, dateStr string) ([]*d.Task, error) {
	rows, err := t.DBConn.QueryxContext(
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

func (t *PGTaskRepository) GetTaskCount(ctx context.Context, userID int, dateStr string) (int, error) {
	var count int
	err := t.DBConn.GetContext(
		ctx, &count,
		"SELECT count(id) FROM tasks WHERE user_id = $1 AND created_at >= $2 AND created_at <= $3",
		userID, dateStr+" 00:00:00", dateStr+" 23:59:59")
	if err != nil {
		return 0, errors.Wrap(err, "db error")
	}

	return count, nil
}
