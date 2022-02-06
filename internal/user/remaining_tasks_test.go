package user

import (
	"context"
	"testing"

	"github.com/jmramos02/akaru/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestRemainingTasksForTheDay(t *testing.T) {
	//initialize db
	db := test.InitTestDB()
	userID := 1

	ctx := context.TODO()
	ctx = context.WithValue(ctx, "db", db)
	ctx = context.WithValue(ctx, "username", "jmramos")

	//seed data into the db
	task := test.CreateTaskTestData(db, userID, "Buy Milk")

	user := Initialize(ctx)
	numberOfTasks := user.GetRemainingTasksForTheDay(userID)

	//Expect it to be one
	assert.Equal(t, 1, numberOfTasks, "Should get one for the current number of task")

	//remove the seed
	db.Delete(&task)
}
