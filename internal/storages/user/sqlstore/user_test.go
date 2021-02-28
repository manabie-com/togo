package sqlstore

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	cmsqlmock "github.com/manabie-com/togo/pkg/common/cmsql/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mock sqlmock.Sqlmock
	userStore *UserStore
)

func TestMain(m *testing.M) {
	var db *sql.DB
	db, mock = cmsqlmock.SetupMock()
	userStore = NewUserStore(db)
	m.Run()
	cmsqlmock.TeardownMock(db)
}

func TestStore_FindByID(t *testing.T) {
	t.Run("TestStore_FindByID", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, password, max_todo FROM users WHERE id = $1").
			WithArgs("00001").
			WillReturnRows(sqlmock.NewRows([]string{"id", "password", "max_todo"}).
				AddRow("00001", "example", 5))

		user, err := userStore.FindByID(context.Background(), sql.NullString{
			String: "00001",
			Valid:  true,
		})

		assert.Nil(t, err)
		assert.Equal(t, "00001", user.ID)
		assert.Equal(t, "example", user.Password)
		assert.Equal(t, 5, user.MaxTodo)
	})

	t.Run("TestStore_FindByID_Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, password, max_todo FROM users WHERE id = $1").
			WithArgs("00002").
			WillReturnRows(sqlmock.NewRows([]string{"id", "password", "max_todo"}))

		user, err := userStore.FindByID(context.Background(), sql.NullString{
			String: "00002",
			Valid:  true,
		})

		assert.Nil(t, user)
		assert.Error(t, err)
	})
}
