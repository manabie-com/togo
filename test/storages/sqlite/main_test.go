// +build integrate

package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/sqlite"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	store *sqlite.LiteDB
)

func TestMain(m *testing.M) {
	db, err := sql.Open("sqlite3", "../../../data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()

	store = &sqlite.LiteDB{
		DB: db,
	}

	os.Exit(m.Run())
}

func Test_Task(t *testing.T) {
	userID := "user_id"
	createDate := "new_date"

	ctx := context.Background()

	random := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	taskID := fmt.Sprintf("unique_task_id_%d", random)
	content := fmt.Sprintf("random_content_%d", random)
	task := &storages.Task{
		ID:          taskID,
		Content:     content,
		UserID:      userID,
		CreatedDate: createDate,
	}
	t.Log(task)

	require.NotNil(t, store)
	err := store.AddTask(ctx, task)
	require.Nil(t, err)

	tasks, err := store.RetrieveTasks(ctx,
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{
			String: createDate,
			Valid:  true,
		})
	require.Nil(t, err)
	require.True(t, len(tasks) > 0)

	err = store.DeleteTask(ctx, task)
	require.Nil(t, err)
}

func Test_ValidateUser(t *testing.T) {
	userID := "firstUser"
	password := "example"

	ctx := context.Background()
	valid := store.ValidateUser(ctx,
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{
			String: password,
			Valid:  true,
		})

	assert.True(t, valid)
}
