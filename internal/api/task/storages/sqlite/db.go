package sqllite

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/api/task/storages"
	"github.com/manabie-com/togo/internal/api/utils"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate string, page, limit int) ([]*storages.Task, error) {

	limit, offset := utils.GetLimitOffsetFormPageNumber(page, limit)

	userIDSQL := sql.NullString{
		String: userID,
		Valid:  true,
	}
	createdDateSql := sql.NullString{
		String: createdDate,
		Valid:  true,
	}

	query := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2 LIMIT $3 OFFSET $4`
	rows, err := l.DB.QueryContext(ctx, query, userIDSQL, createdDateSql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
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
func (l *LiteDB) AddTask(ctx context.Context, t *storages.Task) error {
	query := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	if _, err := l.DB.ExecContext(ctx, query, &t.ID, &t.Content, &t.UserID, &t.CreatedDate); err != nil {
		return err
	}
	return nil
}
