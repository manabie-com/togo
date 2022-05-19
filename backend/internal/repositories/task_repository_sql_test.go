package repositories

import (
	"testing"
	"manabie.com/internal/models"
	"manabie.com/internal/common"
	"database/sql"
	"context"
	"github.com/stretchr/testify/require"
	"fmt"
	"time"
)

func SetUpTaskRepositorySqlTest(iDb *sql.DB) error {
	_, err := iDb.Query(`DELETE FROM "tasks"`)	
	if err != nil {
		return err
	}
	_, err = iDb.Query(`DELETE FROM "user"`)	
	if err != nil {
		return err
	}
	_, err = iDb.Query(`INSERT INTO "user" (
		id,
		name, 
		task_limit
	) VALUES (
		1,
		'test-user-1',
		2
	)`)

	return err
}

func TestTaskRepositorySql(t *testing.T) {
	db := ConnectPostgres()
	defer db.Close()

	var cstHoChiMinh, err = time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		t.Fatal(err)
	}
	prevLocal := time.Local
	time.Local = cstHoChiMinh
	defer func() {
		time.Local = prevLocal
	} ()

	t.Run("create task ok", func (t *testing.T) {
		err := SetUpTaskRepositorySqlTest(db)
		if err != nil {
			t.Fatal(err)
		}

		tx, err := db.Begin()
		defer tx.Commit()
		if err != nil {
			t.Fatal(err)
		}
		context := context.Background()

		
		repository := MakeTaskRepositorySql(tx)
		user := models.MakeUser(1, "test-user-1", 2)
		tasks := []models.Task{}
		now := time.Now()
		year, month, day := now.Date()
		mockCreatedTime := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
		for i := 0; i < 100; i++ {
			tasks = append(tasks, models.MakeTask(
				-1,
				fmt.Sprintf("title-%d",i),
				fmt.Sprintf("content-%d",i),
				common.MakeTime(mockCreatedTime),
				nil,
			))
		}
		tasks, err = repository.CreateTaskForUser(context, user, tasks)
		if err != nil {
			t.Fatal(err)
		}

		require.Equal(t, 100, len(tasks))
		for i, task := range tasks {
			require.NotEqual(t, -1, task.Id)
			require.Equal(t, fmt.Sprintf("title-%d",i), task.Title)
			require.Equal(t, fmt.Sprintf("content-%d",i), task.Content)
			require.True(t, common.MakeTime(mockCreatedTime).IsEqual(task.CreatedTime))
			require.Equal(t, &user, task.Owner)
		}
		result, err := tx.Query(`SELECT COUNT(*) FROM "tasks" WHERE user_id=1`)
		if err != nil {
			t.Fatal(err)
		}
		defer result.Close()
		result.Next()
		count := 0
		result.Scan(&count)
		require.Equal(t, 100, count)
		result.Close()

		t.Run("fetch tasks of user", func (t *testing.T) {
			tasks, err := repository.FetchTasksForUser(context, user)
			if err != nil {
				t.Fatal(err)
			}

			for i, task := range tasks {
				index := len(tasks) - i - 1
				require.NotEqual(t, -1, task.Id)
				require.Equal(t, fmt.Sprintf("title-%d",index), task.Title)
				require.Equal(t, fmt.Sprintf("content-%d",index), task.Content)
				require.True(t, common.MakeTime(mockCreatedTime).IsEqual(task.CreatedTime))
				require.Equal(t, &user, task.Owner)
			}
		})

		t.Run("fetch number of created tasks", func (t *testing.T) {
			tasks := []models.Task{}

			now := time.Now()
			year, month, day := now.Date()
			mockCreatedTime := time.Date(year, month, day + 1, 0, 0, 0, 0, now.Location())

			for i := 0; i < 100; i++ {
				if i == 99 {
					mockCreatedTime = time.Date(year, month, day, 23, 59, 59, 99999999, now.Location())
				}
				tasks = append(tasks, models.MakeTask(
					-1,
					fmt.Sprintf("title-%d",i),
					fmt.Sprintf("content-%d",i),
					common.MakeTime(mockCreatedTime),
					nil,
				))
			}
			_, err = repository.CreateTaskForUser(context, user, tasks)
			if err != nil {
				t.Fatal(err)
			}

			t.Run("for user", func (t *testing.T) {
				count, err := repository.FetchNumberOfTaskForUser(context, user)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, 200, count)
			})


			t.Run("for user today", func (t *testing.T) {
				{
					timeToRetrieve := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

					count, err := repository.FetchNumberOfTaskForUserCreatedOnDay(context, user, common.MakeTime(timeToRetrieve))
					if err != nil {
						t.Fatal(err)
					}
					require.Equal(t, 101, count)
				}

				{
					timeToRetrieve := time.Date(year, month, day + 1, 0, 0, 0, 0, now.Location())

					count, err := repository.FetchNumberOfTaskForUserCreatedOnDay(context, user, common.MakeTime(timeToRetrieve))
					if err != nil {
						t.Fatal(err)
					}
					require.Equal(t, 99, count)
				}
			})

		})
	})

}