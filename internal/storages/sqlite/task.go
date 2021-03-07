package sqllite

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/model"
	"time"
)

// LiteDB for working with sqllite
type TaskLiteDB struct {
	DB *sql.DB
}

func NewTaskLiteDBStorage(db *sql.DB) model.TaskStorage {
	return &TaskLiteDB{
		DB: db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *TaskLiteDB) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*model.Task, error) {

	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.DB.QueryContext(ctx, stmt, convertToNullString(userID), convertToNullString(createdDate))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		t := &model.Task{}
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
func (l *TaskLiteDB) IsAllowedToAddTask(ctx context.Context, userId string) bool {
	createdDate := time.Now().Format("2006-01-02")
	stmtGetCurrentTasksInDay := ` SELECT COALESCE((SELECT COUNT(*) FROM tasks WHERE user_id = ? and created_date = ? GROUP BY user_id), 0)`
	stmtGetMaxToDo := `SELECT max_todo FROM users WHERE id = ?`

	rowTask := l.DB.QueryRowContext(ctx, stmtGetCurrentTasksInDay, convertToNullString(userId),convertToNullString(createdDate))
	rowUser := l.DB.QueryRowContext(ctx, stmtGetMaxToDo, convertToNullString(userId))
	var maxToDo, tasksLengthInDay int

	err := rowTask.Scan(&tasksLengthInDay)
	err = rowUser.Scan(&maxToDo)
	if  maxToDo == tasksLengthInDay || err != nil {
		return false
	}
	return true
}


// AddTask adds a new task to DB
func (l *TaskLiteDB) AddTask(ctx context.Context, t *model.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

func convertToNullString(value string) sql.NullString{
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}