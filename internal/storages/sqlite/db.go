package sqllite

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/model"
	"time"
)

type TaskStore struct {
	DB *sql.DB
}

func NewTaskStore(db *sql.DB) TaskStore {
	return TaskStore{
		DB: db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l TaskStore) RetrieveTasks(ctx context.Context, userID string, createdDate sql.NullString) ([]*model.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, model.NewError(model.ErrListTasks, err.Error())
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		t := &model.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, model.NewError(model.ErrListTasks, err.Error())
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, model.NewError(model.ErrListTasks, err.Error())
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (l TaskStore) AddTask(ctx context.Context, userID string, t *model.Task) (*model.Task, error) {
	t.UserID = userID
	t.CreatedDate = time.Now().UTC().Format("2006-01-02")
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return nil, model.NewError(model.ErrAddTasks, err.Error())
	}

	return t, nil
}

func (l TaskStore) CountTasksByUser(ctx context.Context, userID string, createdDate sql.NullString) (int, error) {
	args := []interface{}{userID}
	stmt := `SELECT count(*) FROM tasks WHERE user_id = ?`
	if createdDate.Valid && len(createdDate.String) != 0 {
		stmt = `SELECT count(*) FROM tasks WHERE user_id = ? and tasks.created_date = ?`
		args = append(args, createdDate)
	}
	row := l.DB.QueryRowContext(ctx, stmt, args...)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, model.NewError(model.ErrCountTasks, err.Error())
	}

	return count, nil
}

type UserStore struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) UserStore {
	return UserStore{
		DB: db,
	}
}

func (l UserStore) Get(ctx context.Context, userID string) (*model.User, error) {
	stmt := `SELECT id, password_hash, max_todo FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID)
	u := &model.User{}
	err := row.Scan(&u.ID, &u.PasswordHash, &u.MaxTodo)
	if err != nil {
		return nil, model.NewError(model.ErrGetUser, err.Error())
	}

	return u, nil

}

func (l UserStore) Create(ctx context.Context, u *model.User) (*model.User, error) {
	return nil, nil
}

// ValidateUser returns tasks if match userID AND password
func (l *UserStore) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &model.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
