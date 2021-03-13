package task

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/valonekowd/togo/domain/entity"
	"github.com/valonekowd/togo/domain/errs"
	"github.com/valonekowd/togo/usecase/interfaces"
)

type sqlRepository struct {
	db *sqlx.DB
}

var _ interfaces.TaskDataSource = sqlRepository{}

func NewSQLRepository(db *sqlx.DB) interfaces.TaskDataSource {
	return sqlRepository{db: db}
}

func (r sqlRepository) List(ctx context.Context, userID string, createdDate time.Time) ([]*entity.Task, error) {
	ts := []*entity.Task{}

	query := `
		SELECT
			*
		FROM
			tasks t
		WHERE
			t.user_id = $1
			AND t.created_at >= DATE_TRUNC('day', $2::TIMESTAMP)
			AND t.created_at < DATE_TRUNC('day', $2::TIMESTAMP) + '1 day'::INTERVAL
	`

	err := r.db.SelectContext(ctx, &ts, query, userID, createdDate)

	return ts, err
}

func (r sqlRepository) Add(ctx context.Context, t *entity.Task) error {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO tasks (id, content, user_id, created_at)
		SELECT
			:id,
			:content,
			:user_id,
			:created_at
		WHERE
			(SELECT u.max_todo FROM users u WHERE u.id = :user_id) >
			(SELECT COUNT(t.id)
			FROM tasks t
			WHERE
				t.user_id = :user_id
				AND t.created_at >= CURRENT_DATE
				AND t.created_at < CURRENT_DATE + INTERVAL '1 day')
	`

	res, err := tx.NamedExecContext(ctx, query, t)
	if err != nil {
		return err
	}

	c, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if c == 0 {
		return errs.ErrMaxToDoExceeded
	}

	return tx.Commit()
}
