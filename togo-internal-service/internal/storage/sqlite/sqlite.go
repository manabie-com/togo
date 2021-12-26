package sqlite

import (
	"context"
	"database/sql"
	"strconv"
	"time"
	"togo-internal-service/internal/model"
	"togo-internal-service/internal/storage"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sony/sonyflake"
)

type sqliteDB struct {
	db     *sql.DB
	config storage.StorageConfig
	sf     *sonyflake.Sonyflake
}

var (
	createStm     = `INSERT INTO tasks(id, user_id, title, content, created_time) VALUES(?,?,?,?,?)`
	checkLimitStm = `SELECT COUNT(*) FROM tasks WHERE user_id = ? AND created_time >= ? AND created_time < ?`
	getStm        = `SELECT id, user_id, title, content, created_time FROM tasks WHERE id = ?`
	listStm       = `SELECT id, user_id, title, SUBSTR(content, 1, ?) content, created_time FROM tasks WHERE user_id = ? AND created_time >= ? ORDER BY user_id ASC, created_time DESC LIMIT ? OFFSET ?`
)

func NewSqliteDB(name string, config storage.StorageConfig) (storage.Storage, error) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}

	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: config.SonyflakeStartTime,
	})

	return &sqliteDB{
		db:     db,
		config: config,
		sf:     sf,
	}, nil
}

func (s *sqliteDB) CreateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	if id, err := s.sf.NextID(); err != nil {
		return nil, err
	} else {
		task.ID = strconv.Itoa(int(id))
	}

	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	if err = s.checkLimitTaskPerDay(ctx, tx, task); err != nil {
		return nil, err
	}

	createdTime := task.CreatedTime.Format(time.RFC3339)
	_, err = tx.ExecContext(ctx, createStm, task.ID, task.UserID, task.Title, task.Content, createdTime)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *sqliteDB) checkLimitTaskPerDay(ctx context.Context, tx *sql.Tx, task *model.Task) error {
	date := time.Date(task.CreatedTime.Year(), task.CreatedTime.Month(), task.CreatedTime.Day(), 0, 0, 0, 0, time.UTC)
	nextDate := date.AddDate(0, 0, 1)

	var countPerUserPerDay int
	err := tx.QueryRowContext(ctx, checkLimitStm, task.UserID, date, nextDate).Scan(&countPerUserPerDay)

	if err != nil {
		_ = tx.Commit()
		return err
	}

	if countPerUserPerDay >= s.config.MaxTaskCreatedPerDay {
		_ = tx.Commit()
		return storage.ErrExceedLimitTaskPerDay
	}
	return nil
}

func (s *sqliteDB) GetTask(ctx context.Context, ID string) (*model.Task, error) {
	rows, err := s.db.QueryContext(ctx, getStm, ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		var task model.Task
		err = rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Content, &task.CreatedTime)
		if err != nil {
			return nil, err
		}
		return &task, nil
	}
	return nil, storage.ErrTaskNotFound
}

func (s *sqliteDB) ListTask(ctx context.Context, userID string, date time.Time, limit, offset int) ([]*model.Task, error) {
	rows, err := s.db.QueryContext(ctx, listStm, s.config.SubstrContentLength, userID, date, limit, offset)
	if err != nil {
		return nil, err
	}
	var tasks []*model.Task
	defer rows.Close()
	for rows.Next() {
		var task model.Task
		var createdTime string
		err = rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Content, &createdTime)
		if err != nil {
			return nil, err
		}
		if task.CreatedTime, err = time.Parse(time.RFC3339, createdTime); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil

}

func (s *sqliteDB) Close() error {
	if s != nil && s.db != nil {
		return s.db.Close()
	}
	return nil
}
