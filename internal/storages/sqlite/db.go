package sqllite

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/storages/entities"
	"github.com/manabie-com/togo/internal/util"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB     *sql.DB
	logger *logs.Logger
}

func NewLitDB() *LiteDB {
	logger := logs.WithPrefix("LiteDB")
	conn := util.Conf.ConnectionString()
	db, err := sql.Open(util.Conf.SqlLiteDriver, conn)
	if err != nil {
		logger.Panic("Create DB occur problem", err.Error())
	}
	return &LiteDB{
		logger: logger,
		DB:     db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate string, limit, offset int) ([]*entities.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ? LIMIT ? OFFSET ?`
	offset = limit * (offset - 1)
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entities.Task
	for rows.Next() {
		t := &entities.Task{}
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
func (l *LiteDB) AddTask(ctx context.Context, t *entities.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd string) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &entities.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		l.logger.Error("Valid user occur problem", err.Error())
		return false
	}

	return true
}
