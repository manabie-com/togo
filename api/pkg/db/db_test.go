package db

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestSetup(t *testing.T) {

	{
		dbManager = nil
		// Success case
		require.NoError(t, Setup())

		require.NoError(t, Teardown())
	}

	{
		// Fail case
		dbManager = &Manager{}
		require.Error(t, Setup())
		// Clear
		dbManager = nil
	}
}

func TestManager_Teardown(t *testing.T) {

	{
		dbManager = nil
		// Success case
		require.NoError(t, Setup())

		require.NoError(t, Teardown())
	}

	{
		// Fail case
		dbManager = nil
		require.Error(t, Teardown())
	}
}

func TestManager_transaction(t *testing.T) {
	ctx := context.Background()

	{
		// Sccess case
		require.Nil(t, Setup())

		// Commit
		require.NoError(t, dbManager.transaction(ctx, &sql.TxOptions{}, func(ctx context.Context, tx *sql.Tx) error {
			_, err := tx.Exec("SELECT id FROM member LIMIT 1")
			return err
		}))

		// Rollback
		require.Error(t, Transaction(ctx, &sql.TxOptions{}, func(ctx context.Context, tx *sql.Tx) error {
			_, _ = tx.Exec("SELECT id FROM member LIMIT 1")
			return errors.New("Roll back")
		}))

		require.Nil(t, Teardown())
	}
}

func TestTransactionForTesting(t *testing.T) {
	require.Nil(t, Setup())

	{
		var id int64
		// Success case
		{
			ctx := context.Background()
			require.NoError(t, TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

				return Transaction(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
					rlt, err := tx.Exec(`INSERT INTO member (email, name) VALUES ($1, $2) RETURNING id`, "testing@example.com", "testing")
					if err != nil {
						return err
					}

					id, err = rlt.RowsAffected()
					if err != nil {
						return err
					}

					return nil
				})

			}))
		}

		{
			ctx := context.Background()
			require.Error(t, Transaction(ctx, &sql.TxOptions{}, func(ctx context.Context, tx *sql.Tx) error {
				row := tx.QueryRow("SELECT id FROM member WHERE id = $1", id)

				var id int
				if err := row.Scan(&id); err != nil {
					return err
				}

				if id != 0 {
					return errors.New("TransactionForTesting not rollback")
				}

				return nil
			}))
		}

	}

	require.Nil(t, Teardown())
	ctx := context.Background()

	{
		// Fail case
		require.Error(t, TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			return nil
		}))
	}
}
