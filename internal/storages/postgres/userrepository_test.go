package postgres

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateUser(t *testing.T) {
	valid := testQueries.ValidateUser(context.Background(), createValidUserParams(t))
	require.Equal(t, true, valid)

	inValid := testQueries.ValidateUser(context.Background(), createInValidUserParams(t))
	require.Equal(t, false, inValid)
}

func createInValidUserParams(t *testing.T) ValidateUserParams {
	user := ValidateUserParams{
		ID:       sql.NullString{String: "user_id", Valid: true},
		Password: sql.NullString{String: "password", Valid: true},
	}

	return user
}

func createValidUserParams(t *testing.T) ValidateUserParams {
	user := ValidateUserParams{
		ID:       sql.NullString{String: "firstUser", Valid: true},
		Password: sql.NullString{String: "example", Valid: true},
	}

	return user
}
