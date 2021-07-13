package storages

import (
	"context"
	"database/sql"
	"time"

	"github.com/manabie-com/togo/internal/entity"
	"github.com/manabie-com/togo/pkg/utils/generator"
)

var (
	nowFunc         = time.Now
	idGeneratorFunc = generator.NewUUID
)

// taskStorage for working with postgres
type taskStorage struct {
	db *sql.DB
}

func NewTaskStorage(db *sql.DB) *taskStorage {
	return &taskStorage{
		db: db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (r *taskStorage) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*entity.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := r.db.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		t := &entity.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (r *taskStorage) AddTask(ctx context.Context, t *entity.Task) error {
	now := nowFunc()
	t.ID = idGeneratorFunc()
	t.CreatedDate = now.Format("2006-01-02")
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// GetNumberOfTasks returns number of task created of user at date
// TODO: use cache couter to return
func (r *taskStorage) GetNumberOfTasks(ctx context.Context, userID, date string) (int, error) {
	stmt := `SELECT count(*) FROM tasks WHERE user_id = $1 and created_date = $2`
	row := r.db.QueryRowContext(ctx, stmt, userID, date)
	var count int
	err := row.Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count, nil
}
