// +build integrate

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	store *postgres.PostgreDB
)

func TestMain(m *testing.M) {
	db, err := sql.Open("pgx", "postgresql://postgres:example@localhost/postgres")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()

	store = &postgres.PostgreDB{
		DB: db,
	}

	os.Exit(m.Run())
}

func Test_Task(t *testing.T) {
	userID := "firstUser"
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

func Test_AddTask_MaxTodo(t *testing.T) {
	ctx := context.Background()

	maxTodo := 5
	userID := "firstUser"
	createdDate := "addTask_maxTodo"

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i < maxTodo + 3; i++ {
		rInt := random.Int()
		taskID := fmt.Sprintf("unique_task_id_%d", rInt)
		content := fmt.Sprintf("random_content_%d", rInt)
		task := &storages.Task{
			ID:          taskID,
			Content:     content,
			UserID:      userID,
			CreatedDate: createdDate,
		}

		require.NotNil(t, store)
		err := store.AddTask(ctx, task)
		require.Nil(t, err)
	}

	tasks, err := store.RetrieveTasks(ctx,
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{
			String: createdDate,
			Valid:  true,
		})
	require.Nil(t, err)
	require.NotNil(t, tasks)
	require.True(t, len(tasks) == maxTodo)
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
