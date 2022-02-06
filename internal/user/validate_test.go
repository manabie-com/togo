package user

import (
	"context"
	"testing"

	"github.com/jmramos02/akaru/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestCanUserInsertSuccess(t *testing.T) {
	db := test.InitTestDB()
	username := "jmramos"
	limit := 2

	ctx := context.TODO()
	ctx = context.WithValue(ctx, "db", db)
	ctx = context.WithValue(ctx, "username", username)

	userModel := test.CreateUserTestData(db, username, limit)
	taskModel1 := test.CreateTaskTestData(db, int(userModel.ID), "Clean the house")

	user := Initialize(ctx)

	result := user.CanUserInsert()
	assert.Equal(t, true, result, "Should be able to insert")

	db.Delete(&taskModel1)
	db.Delete(&userModel)
}
