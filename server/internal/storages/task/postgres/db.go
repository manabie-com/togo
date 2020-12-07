package postgres

import (
	"context"
	"database/sql"
	"github.com/HoangVyDuong/togo/internal/storages/task"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"time"
)

type taskRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *taskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) Close() error {
	return tr.DB.Close()
}

const retrieveTasks = `
	SELECT id, content, user_id
	FROM task
	WHERE user_id = $1 AND deleted_at IS NULL
`
// RetrieveTasks returns tasks if match userID AND createDate.
func (tr *taskRepository) RetrieveTasks(ctx context.Context, userID uint64) ([]task.Task, error) {
	rows, err := tr.DB.QueryContext(ctx, retrieveTasks, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []task.Task
	for rows.Next() {
		var i task.Task
		if err := rows.Scan(&i.ID, &i.Content, &i.UserID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		logger.Errorf("[TaskRepository][RetrieveTasks] error: %s", err.Error())
		return nil, define.DatabaseError
	}
	if err := rows.Err(); err != nil {
		logger.Errorf("[TaskRepository][RetrieveTasks] error: %s", err.Error())
		return nil, define.DatabaseError
	}
	logger.Debugf("%v", items)
	return items, nil
}

const addTask = `
	INSERT INTO task (id, content, user_id, created_at)
	VALUES ($1, $2, $3, $4)
`
// AddTask adds a new task to DB
func (tr *taskRepository) AddTask(ctx context.Context, t task.Task, createdAt time.Time) error {
	_, err := tr.DB.ExecContext(ctx, addTask,
		t.ID,
		t.Content,
		t.UserID,
		createdAt,
	)
	if err != nil {
		logger.Errorf("[TaskRepository][AddTask] error: %s", err.Error())
	}
	return err
}

const softDeleteTask = `
UPDATE task SET deleted_at = $1 WHERE id = $2
`
// SoftDeleteTask
func (tr *taskRepository) SoftDeleteTask(ctx context.Context, taskId uint64, deletedAt time.Time) error {
	_, err := tr.DB.ExecContext(ctx, softDeleteTask, deletedAt, taskId)
	if err != nil {
		logger.Errorf("[TaskRepository][SoftDeleteTask] error: %s", err.Error())
	}
	return err
}