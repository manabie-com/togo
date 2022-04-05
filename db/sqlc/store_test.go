package db

import (
	"context"
	"log"
	"testing"
	"time"
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
	n := int(user.DailyCap)
	// errs := make(chan error)
	// results := make(chan CreateTaskTxResult)
	errs := make([]error, 0)
	results := make([]CreateTaskTxResult, 0)
	ctx := context.Background()
	// ctx := context.WithValue(context.Background(), txKey, txName)
	for i := 0; i < n; i++ {
		// txName := fmt.Sprintf("tx %d", i+1)
		// go func() {
		arg := CreateTaskTxParams{
			User:    user,
			Name:    name,
			Content: content,
		}
		result, err := store.CreateTaskTx(ctx, arg)
		user.DailyQuantity = result.User.DailyQuantity
		// errs <- err
		// results <- result
		results = append(results, result)
		errs = append(errs, err)
		time.Sleep(100 * time.Millisecond)
		// }()
	}
	// Check results
	for i := 0; i < n; i++ {
		// err := <-errs
		err := errs[i]
		require.NoError(t, err)
		// result := <-results
		result := results[i]
		require.NotEmpty(t, result)
		// Check user
		require.NotEmpty(t, result.User)
		require.Equal(t, user.Username, result.User.Username)
		// Check task
		require.NotEmpty(t, result.Task)
		require.Equal(t, name, result.Task.Name)
		require.Equal(t, content, result.Task.Content)
		require.NotZero(t, result.Task.CreatedAt)
		// Strange behavior of testify/require on local versus github worker
		// Error:      	Should be zero, but was 0001-01-01 00:00:00 +0000 UTC
		// require.Zero(t, result.Task.ContentChangeAt)
		// Work around
		require.True(t, result.Task.ContentChangeAt.IsZero())
		require.NotZero(t, result.Task.ID)
		_, err = store.GetTask(ctx, result.Task.ID)
		require.NoError(t, err)
		log.Printf(">> step: %d/%d\n", result.User.DailyQuantity, result.User.DailyCap)
	}
	// Check logic result of concurrent transactions
	updatedUser, err := testQueries.GetUser(ctx, user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	log.Printf(">> after: %d/%d\n", updatedUser.DailyQuantity, updatedUser.DailyCap)
	require.Equal(t, int64(n), updatedUser.DailyQuantity)
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
