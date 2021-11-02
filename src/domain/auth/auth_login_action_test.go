package auth

import (
	"testing"

	"github.com/quochungphp/go-test-assignment/src/domain/users"
	"github.com/quochungphp/go-test-assignment/src/pkgs/db"
	"github.com/stretchr/testify/assert"
)

func TestLoginAction(t *testing.T) {
	pgSession := db.Init()
	userCreateAction := users.UserCreateAction{pgSession}
	user, err := userCreateAction.Execute("user-test", "123456")
	authLoginAction := AuthLoginAction{pgSession}
	token, err := authLoginAction.Execute("user-test", "123456")
	assert.NoError(t, err)
	assert.EqualValues(t, "user-test", user.Username)
	assert.Greater(t, len(token.AccessToken), 20)
	assert.Greater(t, len(token.RefreshToken), 20)

	// Tear down
	_, err = pgSession.Model(new(users.User)).Where("id = ?", user.ID).Delete()
	assert.NoError(t, err)
}
