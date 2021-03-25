package postgres

import (
	"context"
	"database/sql"
)

const validateUser = `-- name: ValidateUser :one
SELECT id FROM users WHERE id = $1 AND password = $2
`

type ValidateUserParams struct {
	ID       sql.NullString `json:"id"`
	Password sql.NullString `json:"password"`
}

// ValidateUser returns tasks if match userID AND password
func (q *Queries) ValidateUser(ctx context.Context, arg ValidateUserParams) bool {
	row := q.queryRow(ctx, q.validateUserStmt, validateUser, arg.ID, arg.Password)

	var id string
	err := row.Scan(&id)

	if err != nil {
		return false
	}

	return true
}
