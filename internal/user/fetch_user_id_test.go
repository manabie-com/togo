package user

import (
	"context"
	"testing"

	"github.com/jmramos02/akaru/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestFetchUserID(t *testing.T) {
	db := test.InitTestDB()
	username := "jmramos"
	limit := 2

	ctx := context.TODO()
	ctx = context.WithValue(ctx, "db", db)
	ctx = context.WithValue(ctx, "username", username)

	userModel := test.CreateUserTestData(db, username, limit)

	user := Initialize(ctx)

	assert.Equal(t, user.GetUserID(), int(userModel.ID), "Should have the same user id")

	db.Delete(&userModel)
}
