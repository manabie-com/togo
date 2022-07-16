package setting

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

func (ms *mockRepository) Create(ctx context.Context, tx *sql.Tx, st *models.Setting) error {
	args := ms.Called(ctx, tx, st)
	return args.Error(0)
}
func (ms *mockRepository) Update(ctx context.Context, tx *sql.Tx, st *models.Setting) error {
	args := ms.Called(ctx, tx, st)
	return args.Error(0)
}
func (ms *mockRepository) FindByMemberID(ctx context.Context, tx *sql.Tx, memberID int) (*models.Setting, error) {
	args := ms.Called(ctx, tx, memberID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Setting), args.Error(1)
}
func (ms *mockRepository) FindByID(ctx context.Context, tx *sql.Tx, ID int) (*models.Setting, error) {
	args := ms.Called(ctx, tx, ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Setting), args.Error(1)
}

func Test_service_Show(t *testing.T) {
	ctx := context.Background()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			exs := &models.Setting{
				ID: 1,
			}

			mck := new(mockRepository)
			mck.On("FindByMemberID", ctx, tx, 1).Return(exs, nil)

			s := NewSettingService(mck)

			out, err := s.Show(ctx, 1)
			require.Nil(t, err)
			require.NotNil(t, out)

			assert.Equal(t, out.ID, exs.ID)

			return nil
		}))
	}

	{
		// Fail case
		require.NotNil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			mck := new(mockRepository)
			mck.On("FindByMemberID", ctx, tx, 1).Return(&models.Setting{}, errors.New("something"))

			// Success case
			s := NewSettingService(mck)

			_, err := s.Show(ctx, 1)

			return err
		}))

		require.NotNil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			mck := new(mockRepository)
			mck.On("FindByMemberID", ctx, tx, 1).Return(nil, sql.ErrNoRows)

			// Success case
			s := NewSettingService(mck)

			_, err := s.Show(ctx, 1)

			return err
		}))
	}
}

func Test_service_Create(t *testing.T) {
	ctx := context.Background()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			exs := &models.Setting{
				MemberID:  1,
				LimitTask: 0,
			}

			mck := new(mockRepository)
			mck.On("FindByMemberID", ctx, tx, 1).Return(nil, nil)
			mck.On("Create", ctx, tx, exs).Return(nil)

			s := NewSettingService(mck)

			require.Nil(t, s.Create(ctx, 1, &models.SettingCreateRequest{}))

			return nil
		}))
	}

	{
		// Fail case by FindByMemberID
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			exs := &models.Setting{
				MemberID:  1,
				LimitTask: 0,
			}

			mck := new(mockRepository)
			mck.On("FindByMemberID", ctx, tx, 1).Return(nil, errors.New("something"))
			mck.On("Create", ctx, tx, exs).Return(nil)

			s := NewSettingService(mck)

			require.NotNil(t, s.Create(ctx, 1, &models.SettingCreateRequest{}))

			return nil
		}))
		// Fail case by setting exits
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			exs := &models.Setting{
				MemberID:  1,
				LimitTask: 0,
			}

			mck := new(mockRepository)
			mck.On("FindByMemberID", ctx, tx, 1).Return(exs, nil)
			mck.On("Create", ctx, tx, exs).Return(nil)

			s := NewSettingService(mck)

			require.NotNil(t, s.Create(ctx, 1, &models.SettingCreateRequest{}))

			return nil
		}))
	}
}

func Test_service_Update(t *testing.T) {
	ctx := context.Background()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			exs := &models.Setting{
				MemberID:  1,
				LimitTask: 0,
			}

			mck := new(mockRepository)
			mck.On("FindByID", ctx, tx, 1).Return(exs, nil)
			mck.On("Update", ctx, tx, exs).Return(nil)

			s := NewSettingService(mck)

			require.Nil(t, s.Update(ctx, 1, &models.SettingUpdateRequest{}))

			return nil
		}))
	}

	{
		// Fail case by FindByID
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			exs := &models.Setting{
				MemberID:  1,
				LimitTask: 0,
			}

			mck := new(mockRepository)
			mck.On("FindByID", ctx, tx, 1).Return(nil, errors.New("something"))
			mck.On("Update", ctx, tx, exs).Return(nil)

			s := NewSettingService(mck)

			require.NotNil(t, s.Update(ctx, 1, &models.SettingUpdateRequest{}))

			return nil
		}))
		// Fail case by NotFound
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

			exs := &models.Setting{
				MemberID:  1,
				LimitTask: 0,
			}

			mck := new(mockRepository)
			mck.On("FindByID", ctx, tx, 1).Return(nil, sql.ErrNoRows)
			mck.On("Update", ctx, tx, exs).Return(nil)

			s := NewSettingService(mck)

			require.NotNil(t, s.Update(ctx, 1, &models.SettingUpdateRequest{}))

			return nil
		}))
	}
}
