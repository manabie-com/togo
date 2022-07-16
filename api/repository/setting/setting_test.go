package setting

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"testing"
	"time"

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

func TestNewSettingRespository(t *testing.T) {
	require.NotNil(t, NewSettingRespository())
}

func Test_settingRespository_Create(t *testing.T) {
	ctx := context.Background()

	userRespository := user.NewUserRespository()
	settingRespository := NewSettingRespository()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    int(time.Now().Unix()),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			if err := userRespository.Create(ctx, tx, u); err != nil {
				return err
			}

			// Create Setting
			st := &models.Setting{
				MemberID:  u.ID,
				LimitTask: 10,
			}

			require.Nil(t, settingRespository.Create(ctx, tx, st))

			out, err := settingRespository.FindByMemberID(ctx, tx, st.MemberID)

			require.Nil(t, err)
			require.NotNil(t, out)

			assert.Equal(t, out.MemberID, u.ID)

			return nil
		}))
	}

	{
		// Fail case
		require.Error(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Create Setting
			st := &models.Setting{
				MemberID:  -1,
				LimitTask: 10,
			}

			return settingRespository.Create(ctx, tx, st)
		}))
	}
}

func Test_settingRespository_Update(t *testing.T) {
	ctx := context.Background()

	userRespository := user.NewUserRespository()
	settingRespository := NewSettingRespository()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    int(time.Now().Unix()),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			if err := userRespository.Create(ctx, tx, u); err != nil {
				return err
			}

			// Create Setting
			require.Nil(t, settingRespository.Create(ctx, tx, &models.Setting{
				MemberID:  u.ID,
				LimitTask: 10,
			}))

			st, err := settingRespository.FindByMemberID(ctx, tx, u.ID)

			require.Nil(t, err)
			require.NotNil(t, st)

			st.LimitTask = 100
			require.Nil(t, settingRespository.Update(ctx, tx, st))

			out, err := settingRespository.FindByID(ctx, tx, st.ID)

			require.Nil(t, err)
			require.NotNil(t, out)

			assert.Equal(t, out.LimitTask, 100)

			return nil
		}))
	}
}

func Test_settingRespository_FindByMemberID(t *testing.T) {
	ctx := context.Background()

	userRespository := user.NewUserRespository()
	settingRespository := NewSettingRespository()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    int(time.Now().Unix()),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			if err := userRespository.Create(ctx, tx, u); err != nil {
				return err
			}

			// Create Setting
			require.Nil(t, settingRespository.Create(ctx, tx, &models.Setting{
				MemberID:  u.ID,
				LimitTask: 10,
			}))

			st, err := settingRespository.FindByMemberID(ctx, tx, u.ID)

			require.Nil(t, err)
			require.NotNil(t, st)

			assert.Equal(t, st.MemberID, u.ID)
			assert.Equal(t, st.LimitTask, 10)

			return nil
		}))
	}

	{
		// Fail case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			out, err := settingRespository.FindByMemberID(ctx, tx, -1)

			require.NotNil(t, err)
			require.Nil(t, out)

			return nil
		}))
	}
}

func Test_settingRespository_FindByID(t *testing.T) {
	ctx := context.Background()

	userRespository := user.NewUserRespository()
	settingRespository := NewSettingRespository()

	{
		// Success case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			// Init Memeber
			u := &models.User{
				ID:    int(time.Now().Unix()),
				Email: "somthing@gmail.com",
				Name:  "something",
			}

			if err := userRespository.Create(ctx, tx, u); err != nil {
				return err
			}

			// Create Setting
			require.Nil(t, settingRespository.Create(ctx, tx, &models.Setting{
				MemberID:  u.ID,
				LimitTask: 10,
			}))

			st, err := settingRespository.FindByMemberID(ctx, tx, u.ID)
			require.Nil(t, err)
			require.NotNil(t, st)

			st, err = settingRespository.FindByID(ctx, tx, st.ID)
			require.Nil(t, err)
			require.NotNil(t, st)

			assert.Equal(t, st.MemberID, u.ID)
			assert.Equal(t, st.LimitTask, 10)

			return nil
		}))
	}

	{
		// Fail case
		require.Nil(t, db.TransactionForTesting(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
			out, err := settingRespository.FindByID(ctx, tx, -1)

			require.NotNil(t, err)
			require.Nil(t, out)

			return nil
		}))
	}
}
