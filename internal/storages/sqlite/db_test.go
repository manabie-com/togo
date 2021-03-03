package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

const (
	UserId      = "user_id"
	Password    = "password"
	CreatedDate = "2021-03-01"
)

var (
	store *LiteDB
)

func TestMain(m *testing.M) {
	db, err := CreateUnitTestDB()
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()
	store = &LiteDB{
		DB: db,
	}

	os.Exit(m.Run())
}

func Test_AddTask(t *testing.T) {
	ctx := context.Background()

	createdDate := "addTask"

	random := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	taskID := fmt.Sprintf("unique_task_id_%d", random)
	content := fmt.Sprintf("random_content_%d", random)
	task := &storages.Task{
		ID:          taskID,
		Content:     content,
		UserID:      UserId,
		CreatedDate: createdDate,
	}
	t.Log(task)

	require.NotNil(t, store)
	_, err := store.AddTask(ctx, task)
	require.Nil(t, err)

	tasks, err := store.RetrieveTasks(ctx,
		sql.NullString{
			String: UserId,
			Valid:  true,
		},
		sql.NullString{
			String: createdDate,
			Valid:  true,
		})
	require.Nil(t, err)
	require.NotNil(t, tasks)
	require.True(t, len(tasks) > 0)
	require.Equal(t, taskID, tasks[len(tasks)-1].ID)
}

func Test_AddTask_MaxTodo(t *testing.T) {
	ctx := context.Background()

	maxTodo := 5
	createdDate := "addTask_maxTodo"

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i < maxTodo + 3; i++ {
		rInt := random.Int()
		taskID := fmt.Sprintf("unique_task_id_%d", rInt)
		content := fmt.Sprintf("random_content_%d", rInt)
		task := &storages.Task{
			ID:          taskID,
			Content:     content,
			UserID:      UserId,
			CreatedDate: createdDate,
		}

		require.NotNil(t, store)
		_, err := store.AddTask(ctx, task)
		require.Nil(t, err)
	}

	tasks, err := store.RetrieveTasks(ctx,
		sql.NullString{
			String: UserId,
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

func Test_RetrieveTasks(t *testing.T) {
	ctx := context.Background()
	tasks, err := store.RetrieveTasks(
		ctx,
		sql.NullString{
			String: UserId,
			Valid:  true,
		},
		sql.NullString{
			String: CreatedDate,
			Valid:  true,
		},
	)
	require.Nil(t, err)
	require.NotNil(t, tasks)

	for index, task := range tasks {
		t.Log(fmt.Sprintf("%d: %v", index, task))
	}
	require.Equal(t, 1, len(tasks))
}

func Test_ValidateUser(t *testing.T) {
	ctx := context.Background()
	valid := store.ValidateUser(ctx,
		sql.NullString{
			String: UserId,
			Valid:  true,
		},
		sql.NullString{
			String: Password,
			Valid:  true,
		})

	assert.True(t, valid)
}
