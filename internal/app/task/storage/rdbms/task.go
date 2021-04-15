package rdbms

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/app/task/model"
	"github.com/manabie-com/togo/internal/util"
)

const SQLRetrieveTasks = `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
const SQLAddTask = `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
const SQLLimitReached = `SELECT COUNT(id) FROM users WHERE id = $1 AND max_todo <= (SELECT COUNT(id) FROM tasks WHERE user_id = $2 AND created_date = $3)`

type TaskStorage struct {
	db         *sql.DB
	driverName string
}

func NewTaskStorage(db *sql.DB, driverName string) *TaskStorage {
	return &TaskStorage{db: db, driverName: driverName}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (t TaskStorage) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]model.Task, error) {
	stmt := util.PrepareQuery(t.driverName, SQLRetrieveTasks)
	rows, err := t.db.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		t := model.Task{}
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
func (t TaskStorage) AddTask(ctx context.Context, task model.Task) error {
	stmt := util.PrepareQuery(t.driverName, SQLAddTask)
	_, err := t.db.ExecContext(ctx, stmt, &task.ID, &task.Content, &task.UserID, &task.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

func (t TaskStorage) LimitReached(ctx context.Context, userID, createdDate string) (bool, error) {
	stmt := util.PrepareQuery(t.driverName, SQLLimitReached)
	rows, err := t.db.QueryContext(ctx, stmt, userID, userID, createdDate)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	c := 0
	for rows.Next() {
		err = rows.Scan(&c)
		if err != nil {
			return false, err
		}
	}
	if err := rows.Err(); err != nil {
		return false, err
	}
	return c == 1, nil
}
