package repositories

import (
	"testing"
	"database/sql"
	"context"
	"github.com/stretchr/testify/assert"
)

func SetUp(iDb *sql.DB) error {
	_, err := iDb.Query(`DELETE FROM "user"`)	
	return err
}

func TestUserRepositorySql(t *testing.T) {
	db := ConnectPostgres()
	defer db.Close()
	t.Run("fetch", func (t *testing.T) {
		err := SetUp(db)
		if err != nil {
			t.Fatal(err)
		}

		/// create user
		_, err = db.Query(`INSERT INTO "user" (
			id,
			name,
			task_limit
		) VALUES (
			1,
			'user-1',
			2
		)`,)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("ok 1")

		tx, err := db.Begin()
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.Background()
		userRepository := MakeUserRepositorySql(tx)
		user, err := userRepository.FetchUserById(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}

		err = tx.Commit()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 1, user.Id)
		assert.Equal(t, "user-1", user.Name)
		assert.Equal(t, 2, user.MaxNumberOfTasks)
	})
}