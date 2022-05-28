package errors

import (
	"database/sql"
	"errors"
)

func IsSQLNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
