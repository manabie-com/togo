package storages

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type ITaskRepo interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, *tools.TodoError)
	AddTask(ctx context.Context, t *Task) *tools.TodoError
}

type TaskRepo struct {
	db DBTX
}

const QueryRetrieveTasks = `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *TaskRepo) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, *tools.TodoError) {
	var err error
	stmt := QueryRetrieveTasks
	rows, err := l.db.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		t := &Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, tools.NewTodoError(http.StatusInternalServerError, err.Error())
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}

	return tasks, nil
}

const QueryAddTask = `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`

// AddTask adds a new task to DB
func (l *TaskRepo) AddTask(ctx context.Context, t *Task) *tools.TodoError {
	stmt := QueryAddTask
	_, err := l.db.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func NewTaskRepo(db *sql.DB) ITaskRepo {
	return &TaskRepo{
		db: db,
	}
}
