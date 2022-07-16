package task

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"testing"
	"time"

	"manabie/todo/models"
	"manabie/todo/pkg/apiutils"
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

type mockTaskRepository struct {
	mock.Mock
}

func (ms *mockTaskRepository) Find(ctx context.Context, tx *sql.Tx, memberID int, date string) ([]*models.Task, error) {
	args := ms.Called(ctx, tx, memberID, date)
	return args.Get(0).([]*models.Task), args.Error(1)
}
func (ms *mockTaskRepository) FindForUpdate(ctx context.Context, tx *sql.Tx, memberID int, date string) ([]*models.Task, error) {
	args := ms.Called(ctx, tx, memberID, date)
	return args.Get(0).([]*models.Task), args.Error(1)
}
func (ms *mockTaskRepository) FindByID(ctx context.Context, tx *sql.Tx, ID int, forUpdate bool) (*models.Task, error) {
	args := ms.Called(ctx, tx, ID, forUpdate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}
func (ms *mockTaskRepository) Create(ctx context.Context, tx *sql.Tx, tk *models.Task) error {
	args := ms.Called(ctx, tx, tk)
	return args.Error(0)
}
func (ms *mockTaskRepository) Update(ctx context.Context, tx *sql.Tx, tk *models.Task) error {
	args := ms.Called(ctx, tx, tk)
	return args.Error(0)
}
func (ms *mockTaskRepository) Delete(ctx context.Context, tx *sql.Tx, tk *models.Task) error {
	args := ms.Called(ctx, tx, tk)
	return args.Error(0)
}

type mockSettingRepository struct {
	mock.Mock
}

func (ms *mockSettingRepository) Create(ctx context.Context, tx *sql.Tx, st *models.Setting) error {
	return nil
}
func (ms *mockSettingRepository) Update(ctx context.Context, tx *sql.Tx, st *models.Setting) error {
	return nil
}
func (ms *mockSettingRepository) FindByMemberID(ctx context.Context, tx *sql.Tx, memberID int) (*models.Setting, error) {
	args := ms.Called(ctx, tx, memberID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Setting), args.Error(1)
}
func (ms *mockSettingRepository) FindByID(ctx context.Context, tx *sql.Tx, ID int) (*models.Setting, error) {
	return nil, nil
}

func Test_service_Index(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			exs := []*models.Task{
				{ID: 1, MemberID: 1, Content: "something", TargetDate: time.Time{}},
				{ID: 2, MemberID: 2, Content: "something", TargetDate: time.Time{}},
			}

			mck := new(mockTaskRepository)
			mck.On("Find", ctx, tx, 1, "2022-01-01").Return(exs, nil)

			s := NewTaskService(mck, nil)

			tks, err := s.Index(ctx, 1, "2022-01-01")
			require.Nil(t, err)
			require.NotNil(t, tks)

			assert.Len(t, tks, 2)

			return nil
		}))
	}

	{
		// Fail case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			exs := []*models.Task{}

			mck := new(mockTaskRepository)
			mck.On("Find", ctx, tx, 1, "2022-01-01").Return(exs, errors.New("something"))

			s := NewTaskService(mck, nil)

			_, err := s.Index(ctx, 1, "2022-01-01")
			require.NotNil(t, err)

			return nil
		}))
	}
}

func Test_service_Show(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			exs := &models.Task{
				ID: 1, MemberID: 1, Content: "something", TargetDate: time.Time{},
			}

			mck := new(mockTaskRepository)
			mck.On("FindByID", ctx, tx, 1, false).Return(exs, nil)

			s := NewTaskService(mck, nil)

			tks, err := s.Show(ctx, 1)
			require.Nil(t, err)
			require.NotNil(t, tks)

			assert.Equal(t, tks.ID, exs.ID)

			return nil
		}))
	}

	{
		// Fail case by FindByID
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			mck := new(mockTaskRepository)
			mck.On("FindByID", ctx, tx, 1, false).Return(nil, errors.New("something"))

			s := NewTaskService(mck, nil)

			tks, err := s.Show(ctx, 1)
			require.NotNil(t, err)
			require.Nil(t, tks)

			return nil
		}))

		// Fail case by NotFound
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			mck := new(mockTaskRepository)
			mck.On("FindByID", ctx, tx, 1, false).Return(nil, sql.ErrNoRows)

			s := NewTaskService(mck, nil)

			tks, err := s.Show(ctx, 1)
			require.NotNil(t, err)
			require.True(t, errors.Cause(err) == apiutils.ErrNotFound)
			require.Nil(t, tks)

			return nil
		}))
	}
}

func Test_service_Create(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			ex := &models.Task{
				MemberID:   1,
				Content:    "some",
				TargetDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			}

			mcSet := new(mockSettingRepository)
			mcSet.On("FindByMemberID", ctx, tx, 1).Return(&models.Setting{LimitTask: 2}, nil)
			mcTask := new(mockTaskRepository)
			mcTask.On("FindForUpdate", ctx, tx, 1, "2022-01-01").Return([]*models.Task{{}}, nil)
			mcTask.On("Create", ctx, tx, ex).Return(nil)

			s := NewTaskService(mcTask, mcSet)

			return s.Create(ctx, 1, &models.TaskCreateRequest{
				Content:    "some",
				TargetDate: "2022-01-01",
			})
		}))
	}

	{
		// Fail case
		tests := []struct {
			name      string
			tkMock    func(context.Context, *sql.Tx, int, string, *models.Task) *mockTaskRepository
			stMock    func(context.Context, *sql.Tx, int) *mockSettingRepository
			tkRequest *models.TaskCreateRequest
			tk        *models.Task
		}{
			{
				name: "FindByMemberID",
				tkMock: func(ctx context.Context, tx *sql.Tx, i int, s string, t *models.Task) *mockTaskRepository {
					return new(mockTaskRepository)
				},
				stMock: func(ctx context.Context, tx *sql.Tx, memberId int) *mockSettingRepository {
					mc := new(mockSettingRepository)
					mc.On("FindByMemberID", ctx, tx, memberId).Return(nil, errors.New("something"))
					return mc
				},
				tkRequest: &models.TaskCreateRequest{TargetDate: "2022-01-01"},
				tk:        &models.Task{},
			},
			{
				name: "Setting NotFound",
				tkMock: func(ctx context.Context, tx *sql.Tx, i int, s string, t *models.Task) *mockTaskRepository {
					return new(mockTaskRepository)
				},
				stMock: func(ctx context.Context, tx *sql.Tx, memberId int) *mockSettingRepository {
					mc := new(mockSettingRepository)
					mc.On("FindByMemberID", ctx, tx, memberId).Return(nil, sql.ErrNoRows)
					return mc
				},
				tkRequest: &models.TaskCreateRequest{TargetDate: "2022-01-01"},
				tk:        &models.Task{},
			},
			{
				name: "FindForUpdate",
				tkMock: func(ctx context.Context, tx *sql.Tx, memberId int, targetDate string, t *models.Task) *mockTaskRepository {
					mc := new(mockTaskRepository)
					mc.On("FindForUpdate", ctx, tx, memberId, targetDate).Return([]*models.Task{}, errors.New("something"))
					return mc
				},
				stMock: func(ctx context.Context, tx *sql.Tx, memberId int) *mockSettingRepository {
					mc := new(mockSettingRepository)
					mc.On("FindByMemberID", ctx, tx, memberId).Return(&models.Setting{LimitTask: 2}, nil)
					return mc
				},
				tkRequest: &models.TaskCreateRequest{TargetDate: "2022-01-01"},
			},
			{
				name: "Validate",
				tkMock: func(ctx context.Context, tx *sql.Tx, memberId int, targetDate string, t *models.Task) *mockTaskRepository {
					mc := new(mockTaskRepository)
					mc.On("FindForUpdate", ctx, tx, memberId, targetDate).Return([]*models.Task{{}}, nil)
					return mc
				},
				stMock: func(ctx context.Context, tx *sql.Tx, memberId int) *mockSettingRepository {
					mc := new(mockSettingRepository)
					mc.On("FindByMemberID", ctx, tx, memberId).Return(&models.Setting{LimitTask: 1}, nil)
					return mc
				},
				tkRequest: &models.TaskCreateRequest{TargetDate: "2022-01-01"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				require.Error(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
					s := NewTaskService(
						tt.tkMock(ctx, tx, 1, "2022-01-01", tt.tk),
						tt.stMock(ctx, tx, 1),
					)
					return s.Create(ctx, 1, tt.tkRequest)
				}))
			})
		}
	}
}

func Test_service_Update(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			mc := new(mockTaskRepository)
			mc.On("FindByID", ctx, tx, 1, true).Return(&models.Task{ID: 1}, nil)
			mc.On("Update", ctx, tx, &models.Task{ID: 1}).Return(nil)

			s := NewTaskService(mc, nil)
			require.Nil(t, s.Update(ctx, 1, &models.Task{}))

			return nil
		}))
	}

	{
		// Fail case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			mc := new(mockTaskRepository)
			mc.On("FindByID", ctx, tx, 1, true).Return(nil, sql.ErrNoRows)

			s := NewTaskService(mc, nil)
			require.NotNil(t, s.Update(ctx, 1, &models.Task{}))

			return nil
		}))
	}
}

func Test_service_Delete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			mc := new(mockTaskRepository)
			mc.On("FindByID", ctx, tx, 1, true).Return(&models.Task{ID: 1}, nil)
			mc.On("Delete", ctx, tx, &models.Task{ID: 1}).Return(nil)

			s := NewTaskService(mc, nil)
			require.Nil(t, s.Delete(ctx, 1))

			return nil
		}))
	}

	{
		// Fail case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			mc := new(mockTaskRepository)
			mc.On("FindByID", ctx, tx, 1, true).Return(nil, sql.ErrNoRows)

			s := NewTaskService(mc, nil)
			require.NotNil(t, s.Delete(ctx, 1))

			return nil
		}))
	}
}
