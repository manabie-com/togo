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

	ctx := context.TODO()
	ctx = context.WithValue(ctx, "db", db)
	ctx = context.WithValue(ctx, "username", username)

	testCases := []struct {
		Limit           int
		ExepectedResult bool
		Description     string
	}{
		{
			Limit:           2,
			ExepectedResult: true,
			Description:     "Should be able to insert",
		},
		{
			Limit:           1,
			ExepectedResult: false,
			Description:     "Should not be able to insert",
		},
	}

	for _, tc := range testCases {

		userModel := test.CreateUserTestData(db, username, tc.Limit)
		taskModel1 := test.CreateTaskTestData(db, int(userModel.ID), "Clean the house")

		user := Initialize(ctx)

		result := user.CanUserInsert()
		assert.Equal(t, result, tc.ExepectedResult, tc.Description)

		db.Delete(&taskModel1)
		db.Delete(&userModel)
	}

}
