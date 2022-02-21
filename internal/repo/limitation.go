package repo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/chi07/todo/internal/model"
)

type Limitation struct {
	db *sqlx.DB
}

func NewLimitation(db *sqlx.DB) *Limitation {
	return &Limitation{db: db}
}

func (ul *Limitation) GetByUserID(ctx context.Context, userID uuid.UUID) (*model.Limitation, error) {
	var limitation model.Limitation
	err := ul.db.GetContext(ctx, &limitation, "SELECT id, user_id, limit_tasks FROM limitations where user_id=$1", userID)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "cannot get user limitation from DB")
	}
	return &limitation, nil
}
