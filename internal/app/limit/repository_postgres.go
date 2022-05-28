package limit

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dinhquockhanh/togo/internal/pkg/errors"
	db "github.com/dinhquockhanh/togo/internal/pkg/sql/sqlc"
)

type (
	PostgresRepository struct {
		q *db.Queries
	}
)

func NewPostgresRepository(cnn *sql.DB) Repository {
	return &PostgresRepository{
		q: db.New(cnn),
	}
}

func (r *PostgresRepository) GetLimit(ctx context.Context, req *GetLimitReq) (*Limit, error) {
	limit, err := r.q.GetLimit(ctx, &db.GetLimitParams{
		TierID: int16(req.TierID),
		Action: req.Action,
	})

	if err != nil {
		op := "postgresRepository.AssignTask"
		if errors.IsSQLNotFound(err) {
			return nil, errors.NewNotFoundErr(err, op, fmt.Sprintf("limit with tier_id=%d, action=%s", req.TierID, req.Action))
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Limit{
		TierID: int(limit.TierID),
		Action: limit.Action,
		Value:  int(limit.Value),
	}, nil
}
