package user

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"testing"

	"manabie/todo/models"
	"manabie/todo/pkg/db"
	"manabie/todo/pkg/utils"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {

	if !flag.Parsed() {
		flag.Parse()
	}

	if err := db.Setup(); err != nil {
		panic(err)
	}

	exit := m.Run()
	if err := db.Teardown(); err != nil {
		panic(err)
	}
	os.Exit(exit)

}

func Test_userRespository_Create(t *testing.T) {
	ctx := context.Background()

	userRespository := NewUserRespository()

	{
		// Success case with id
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    utils.RamdomID(),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			require.Nil(t, userRespository.Create(ctx, tx, u))

			outs, err := userRespository.Find(ctx, tx)
			require.Nil(t, err)

			assert.GreaterOrEqual(t, len(outs), 1)

			return nil
		}))
	}

	{
		// Success case without id
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			require.Nil(t, userRespository.Create(ctx, tx, u))

			outs, err := userRespository.Find(ctx, tx)
			require.Nil(t, err)

			assert.GreaterOrEqual(t, len(outs), 1)

			return nil
		}))
	}
}
