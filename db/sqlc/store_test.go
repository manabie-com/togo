package db

import (
	"context"
	"log"
	"testing"
	"togo/util"

	"github.com/stretchr/testify/require"
)

func TestCreateTaskTxResult(t *testing.T) {
	store := NewStore(testDB)
	user := createRandomUser(t)
	name := util.RandomName()
	content := util.RandomString(100)
	user.DailyCap = util.RandomInt(1, 20)
	user.DailyQuantity = 0
	updateRandomedUser(t, user.Username, user.DailyCap, user.DailyQuantity)
	log.Printf(">> before: %d/%d\n", user.DailyQuantity, user.DailyCap)
	// Run n concurrent creations
	n := 1
	errs := make(chan error)
	results := make(chan CreateTaskTxResult)
	ctx := context.Background()
	// ctx := context.WithValue(context.Background(), txKey, txName)
	for i := 0; i < n; i++ {
		// txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			result, err := store.CreateTaskTx(ctx, CreateTaskTxParams{
				User:    user,
				Name:    name,
				Content: content,
			})
			errs <- err
			results <- result
		}()
	}
	// Check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)
		// Check user
		require.NotEmpty(t, result.User)
		require.Equal(t, user.Username, result.User.Username)
		// Check task
		require.NotEmpty(t, result.Task)
		require.Equal(t, name, result.Task.Name)
		require.Equal(t, content, result.Task.Content)
		require.NotZero(t, result.Task.CreatedAt)
		require.Zero(t, result.Task.ContentChangeAt)
		require.NotZero(t, result.Task.ID)
		_, err = store.GetTask(ctx, result.Task.ID)
		require.NoError(t, err)
		// Check logic
		var expectedQuantity int64
		if user.DailyQuantity >= user.DailyCap {
			expectedQuantity = user.DailyQuantity
		} else {
			count, err := store.CountTasksCreatedToday(ctx, result.User.Username)
			require.NoError(t, err)
			if result.User.DailyQuantity != count {
				expectedQuantity = 1
			} else {
				expectedQuantity = user.DailyQuantity + 1
			}
		}
		require.Equal(t, result.User.DailyQuantity, expectedQuantity)
		updatedUser, err := testQueries.GetUser(ctx, result.User.Username)
		require.NoError(t, err)
		require.NotEmpty(t, updatedUser)
		log.Printf(">> after: %d/%d\n", updatedUser.DailyQuantity, updatedUser.DailyCap)
		require.LessOrEqual(t, updatedUser.DailyQuantity, updatedUser.DailyCap)
	}
}

func TestCreateTaskTxResultDailyCapExceed(t *testing.T) {
	store := NewStore(testDB)
	user := createRandomUser(t)
	name := util.RandomName()
	content := util.RandomString(100)
	user.DailyCap = util.RandomInt(1, 20)
	user.DailyQuantity = user.DailyCap
	updateRandomedUser(t, user.Username, user.DailyCap, user.DailyQuantity)
	log.Printf(">> before: %d/%d\n", user.DailyQuantity, user.DailyCap)
	// Run n concurrent creations
	n := 1
	errs := make(chan error)
	results := make(chan CreateTaskTxResult)
	ctx := context.Background()
	// ctx := context.WithValue(context.Background(), txKey, txName)
	for i := 0; i < n; i++ {
		// txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			result, err := store.CreateTaskTx(ctx, CreateTaskTxParams{
				User:    user,
				Name:    name,
				Content: content,
			})
			errs <- err
			results <- result
		}()
	}
	// Check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.Error(t, err)
		require.EqualError(t, err, "daily limit exceed")
		result := <-results
		require.Empty(t, result.Task)
		updatedUser, err := testQueries.GetUser(ctx, result.User.Username)
		require.NoError(t, err)
		require.NotEmpty(t, updatedUser)
		log.Printf(">> after: %d/%d\n", updatedUser.DailyQuantity, updatedUser.DailyCap)
		// Should be no increase in quantity when the request is rejected
		require.Equal(t, updatedUser.DailyQuantity, updatedUser.DailyCap)
	}
}
