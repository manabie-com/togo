package users

import (
	"testing"

	"github.com/quochungphp/go-test-assignment/src/pkgs/db"
	"github.com/stretchr/testify/assert"
)

func TestUserCreateAction(t *testing.T) {
	pgSession := db.Init()
	userCreateAction := UserCreateAction{pgSession}
	user, err := userCreateAction.Execute("user-test", "123456")

	assert.NoError(t, err)
	assert.EqualValues(t, "user-test", user.Username)
	assert.EqualValues(t, 5, user.MaxTodo)

	// Tear down
	_, err = pgSession.Model(new(User)).Where("id = ?", user.ID).Delete()
	assert.NoError(t, err)
}
