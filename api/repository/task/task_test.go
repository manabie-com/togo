package task

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"database/sql"
	"manabie/todo/models"
	"manabie/todo/pkg/db"
	"manabie/todo/repository/user"

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

func Test_taskRespository_Find(t *testing.T) {
	ctx := context.Background()

	userRespository := user.NewUserRespository()
	taskRespository := NewTaskRespository()

	{
		// Success case with id
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    int(time.Now().Unix()),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			require.Nil(t, userRespository.Create(ctx, tx, u))

			tk := &models.Task{
				MemberID:   u.ID,
				Content:    "something",
				TargetDate: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			}
			require.Nil(t, taskRespository.Create(ctx, tx, tk))

			out, err := taskRespository.FindByID(ctx, tx, tk.ID, false)
			require.Nil(t, err)
			require.NotNil(t, out)

			return nil
		}))
	}
}

func Test_taskRespository_FindByID(t *testing.T) {
	ctx := context.Background()

	userRespository := user.NewUserRespository()
	taskRespository := NewTaskRespository()

	{
		// Success case with id
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    int(time.Now().Unix()),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			require.Nil(t, userRespository.Create(ctx, tx, u))

			tk := &models.Task{
				MemberID:   u.ID,
				Content:    "something",
				TargetDate: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			}

			require.Nil(t, taskRespository.Create(ctx, tx, tk))

			out, err := taskRespository.FindByID(ctx, tx, tk.ID, true)
			require.Nil(t, err)
			require.NotNil(t, out)

			return nil
		}))
	}

	{
		// Fail case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			out, err := taskRespository.FindByID(ctx, tx, -1, false)
			require.NotNil(t, err)
			require.Nil(t, out)

			return nil
		}))
	}
}

func Test_taskRespository_FindForUpdate(t *testing.T) {
	ctx := context.Background()

	userRespository := user.NewUserRespository()
	taskRespository := NewTaskRespository()

	{
		// Success case with id
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    int(time.Now().Unix()),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			require.Nil(t, userRespository.Create(ctx, tx, u))

			require.Nil(t, taskRespository.Create(ctx, tx, &models.Task{
				MemberID:   u.ID,
				Content:    "something",
				TargetDate: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			}))

			outs, err := taskRespository.FindForUpdate(ctx, tx, u.ID, "2022-01-02")
			require.Nil(t, err)
			require.NotNil(t, outs)

			assert.GreaterOrEqual(t, len(outs), 1)

			return nil
		}))
	}
}

func Test_taskRespository_Create(t *testing.T) {
	ctx := context.Background()

	userRespository := user.NewUserRespository()
	taskRespository := NewTaskRespository()

	{
		// Success case with id
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    int(time.Now().Unix()),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			require.Nil(t, userRespository.Create(ctx, tx, u))

			tk := &models.Task{
				MemberID:   u.ID,
				Content:    "something",
				TargetDate: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			}

			require.Nil(t, taskRespository.Create(ctx, tx, tk))

			outs, err := taskRespository.FindByID(ctx, tx, tk.ID, false)
			require.Nil(t, err)
			require.NotNil(t, outs)

			return nil
		}))
	}

	{
		// Success case with id
		require.Error(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			return taskRespository.Create(ctx, tx, &models.Task{
				MemberID:   -10,
				Content:    "something",
				TargetDate: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			})
		}))
	}
}

func Test_taskRespository_Update(t *testing.T) {
	ctx := context.Background()

	userRespository := user.NewUserRespository()
	taskRespository := NewTaskRespository()

	{
		// Success case with id
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    int(time.Now().Unix()),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			require.Nil(t, userRespository.Create(ctx, tx, u))

			tk := &models.Task{
				MemberID:   u.ID,
				Content:    "something",
				TargetDate: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			}
			require.Nil(t, taskRespository.Create(ctx, tx, tk))

			tk.Content = "updated"

			require.Nil(t, taskRespository.Update(ctx, tx, tk))

			out, err := taskRespository.FindByID(ctx, tx, tk.ID, false)
			require.Nil(t, err)
			require.NotNil(t, out)

			assert.Equal(t, tk.Content, "updated")

			return nil
		}))
	}
}

func Test_taskRespository_Delete(t *testing.T) {
	ctx := context.Background()

	userRespository := user.NewUserRespository()
	taskRespository := NewTaskRespository()

	{
		// Success case with id
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    int(time.Now().Unix()),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			require.Nil(t, userRespository.Create(ctx, tx, u))

			tk := &models.Task{
				MemberID:   u.ID,
				Content:    "something",
				TargetDate: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			}

			require.Nil(t, taskRespository.Delete(ctx, tx, tk))

			_, err := taskRespository.FindByID(ctx, tx, tk.ID, false)
			require.NotNil(t, err)

			return nil
		}))
	}
}
