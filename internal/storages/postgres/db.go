package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/manabie-com/togo/internal/storages"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

var DailyLimitExceededError = errors.New("Daily tasks limit exceeded")

// PostgresDB for working with PostgreSQL
type PostgresDB struct {
	DB *pgx.Conn
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (pg *PostgresDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	const query = "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2"

	rows, err := pg.DB.Query(ctx, query, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		var t storages.Task

		if err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate); err != nil {
			return nil, err
		}

		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (pg *PostgresDB) AddTask(ctx context.Context, t *storages.Task) error {
	const (
		maxQuery = "SELECT max_todo FROM users WHERE id = $1"
		cntQuery = "SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND created_date = $2"
		insQuery = "INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)"
		maxRetry = 100
	)

	retry := func() error {
		var (
			maxToDo int
			cntToDo int
		)

		tx, err := pg.DB.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
		if err != nil {
			return err
		}
		defer tx.Rollback(ctx)

		if err := tx.QueryRow(ctx, maxQuery, t.UserID).Scan(&maxToDo); err != nil {
			return err
		}

		if err := tx.QueryRow(ctx, cntQuery, t.UserID, t.CreatedDate).Scan(&cntToDo); err != nil {
			return err
		}

		if cntToDo >= maxToDo {
			return DailyLimitExceededError
		}

		if _, err := tx.Exec(ctx, insQuery, t.ID, t.Content, t.UserID, t.CreatedDate); err != nil {
			return err
		}

		if err := tx.Commit(ctx); err != nil {
			return err
		}

		return nil
	}

	for i := 0; i < maxRetry; i++ {
		if err := retry(); err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "40001" {
				// ERROR:  could not serialize access due to read/write dependencies among transactions
				continue
			}

			return err
		}

		break
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (pg *PostgresDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	const query = "SELECT id FROM users WHERE id = $1 AND password = $2"

	var u storages.User

	if err := pg.DB.QueryRow(ctx, query, userID, pwd).Scan(&u.ID); err != nil {
		return false
	}

	return true
}
