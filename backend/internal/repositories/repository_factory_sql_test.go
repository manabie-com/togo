package repositories

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"fmt"
	"manabie.com/internal/common"
	"manabie.com/internal/models"
	"time"
)

func verifyTaskCount(expected int, db *sql.DB, t *testing.T) {
	result, err := db.Query(`SELECT COUNT(*) FROM "tasks" WHERE user_id=1`)
	if err != nil {
		t.Fatal(err)
	}
	defer result.Close()
	result.Next()
	count := 0
	result.Scan(&count)
	require.Equal(t, expected, count)
	result.Close()
}

func TestRepositoryFactorySql(t *testing.T) {
	db := ConnectPostgres()
	defer db.Close()


	t.Run("can create repositories", func (t *testing.T) {
		factory := MakeRepositoryFactorySql(db)
		ctx := context.Background()
		testCreateRepositories := func (oDone chan error) {
			err, txErr := factory.StartTransactionAuto(
				ctx, 
				ReadCommitted,
				func (iTransactionId TransactionId) error  {
					time.Sleep(500 * time.Millisecond)
					_, err := factory.GetTaskRepository(iTransactionId)
					if err != nil {
						return err
					}

					time.Sleep(500 * time.Millisecond)
					_, err = factory.GetUserRepository(iTransactionId)
					if err != nil {
						return err
					}

					return nil
				},
			)
			require.Equal(t, nil, err)
			require.Equal(t, nil, txErr)
			oDone <- nil
		}

		/// create multiple transactions simultaneously
		numberOfTx := 50
		channels := make([]chan error, 0, numberOfTx)
		for i := 0; i < numberOfTx; i++ {
			channel := make(chan error)
			go testCreateRepositories(channel)
			channels = append(channels, channel)
		}

		for i := 0; i < numberOfTx; i++ {
			res := <- channels[i]
			require.Equal(t, nil, res)
		}
	})

	t.Run("commit if ok", func (t *testing.T) {
		factory := MakeRepositoryFactorySql(db)
		SetUpTaskRepositorySqlTest(db)

		verifyTaskCount(0, db, t)

		ctx := context.Background()
		err, txErr := factory.StartTransactionAuto(
			ctx, 
			ReadCommitted,
			func (iTransactionId TransactionId) error  {
				repository, err := factory.GetTaskRepository(iTransactionId)
				if err != nil {
					t.Fatal(err)
				}

				user := models.MakeUser(1, "test-user-1", 2)
				tasks := []models.Task{}

				for i := 0; i < 1; i++ {
					tasks = append(tasks, models.MakeTask(
						-1,
						fmt.Sprintf("title-%d",i),
						fmt.Sprintf("content-%d",i),
						common.MakeTime(time.Unix(1, 0)),
						nil,
					))
				}
				tasks, err = repository.CreateTaskForUser(ctx, user, tasks)
				if err != nil {
					t.Fatal(err)
				}

				require.Equal(t, 1, len(tasks))
				return nil
			},
		)
		verifyTaskCount(1, db, t)
		require.Equal(t, nil, err)
		require.Equal(t, nil, txErr)
	})


	t.Run("rollback if error", func (t *testing.T) {
		factory := MakeRepositoryFactorySql(db)
		SetUpTaskRepositorySqlTest(db)

		verifyTaskCount(0, db, t)

		ctx := context.Background()
		mockError := fmt.Errorf("error")
		err, txErr := factory.StartTransactionAuto(
			ctx, 
			ReadCommitted,
			func (iTransactionId TransactionId) error  {
				repository, err := factory.GetTaskRepository(iTransactionId)
				if err != nil {
					t.Fatal(err)
				}

				user := models.MakeUser(1, "test-user-1", 2)
				tasks := []models.Task{}

				for i := 0; i < 1; i++ {
					tasks = append(tasks, models.MakeTask(
						-1,
						fmt.Sprintf("title-%d",i),
						fmt.Sprintf("content-%d",i),
						common.MakeTime(time.Unix(1, 0)),
						nil,
					))
				}
				tasks, err = repository.CreateTaskForUser(ctx, user, tasks)
				if err != nil {
					t.Fatal(err)
				}

				require.Equal(t, 1, len(tasks))
				return mockError
			},
		)

		verifyTaskCount(0, db, t)
		require.Equal(t, mockError, err)
		require.Equal(t, nil, txErr)
	})
}