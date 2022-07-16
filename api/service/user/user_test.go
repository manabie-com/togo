package user

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"testing"

	"manabie/todo/models"
	"manabie/todo/pkg/db"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type mockRepository struct {
	mock.Mock
}

func (ms *mockRepository) Find(ctx context.Context, tx *sql.Tx) ([]*models.User, error) {
	args := ms.Called(ctx, tx)
	return args.Get(0).([]*models.User), args.Error(1)
}
func (ms *mockRepository) Create(ctx context.Context, tx *sql.Tx, u *models.User) error { return nil }

func Test_service_Index(t *testing.T) {
	ctx := context.Background()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			exs := []*models.User{
				{ID: 1, Email: "example_1@gmail.com", Name: "example_1"},
				{ID: 2, Email: "example_2@gmail.com", Name: "example_2"},
			}

			mck := new(mockRepository)
			mck.On("Find", ctx, tx).Return(exs, nil)

			s := NewUserService(mck)

			outs, err := s.Index(ctx)
			require.Nil(t, err)
			require.NotNil(t, outs)

			assert.Len(t, outs, 2)
			return nil
		}))
	}

	{
		// Fail case
		require.NotNil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			mck := new(mockRepository)
			mck.On("Find", ctx, tx).Return([]*models.User{}, errors.New("something"))

			// Success case
			s := NewUserService(mck)

			_, err := s.Index(ctx)

			return err
		}))
	}
}
