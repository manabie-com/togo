package tasks

import (
	"testing"

	"github.com/quochungphp/go-test-assignment/src/domain/users"
	"github.com/quochungphp/go-test-assignment/src/pkgs/db"
	"github.com/quochungphp/go-test-assignment/src/pkgs/token"
	"github.com/stretchr/testify/assert"
)

func TestPreserveM2(t *testing.T) {
	pgSession := db.Init()
	userCreateAction := users.UserCreateAction{pgSession}
	user, err := userCreateAction.Execute("test", "123456")
	assert.NoError(t, err)

	token.AccessUser = token.AccessUserInfo{
		CorrelationID: "fb9ed746-759e-4ce2-af3f-3903100e4760",
		MaxTodo:       user.MaxTodo,
		UserID:        user.ID,
	}
	createTaskAction := TaskCreateAction{pgSession}
	task, err := createTaskAction.Execute("I'm a M10")
	assert.NoError(t, err)
	assert.EqualValues(t, "I'm a M10", task.Content)

	// Tear down
	_, err = pgSession.Model(new(Tasks)).Where("id = ?", task.ID).Delete()
	assert.NoError(t, err)

	_, err = pgSession.Model(new(users.User)).Where("id = ?", user.ID).Delete()
	assert.NoError(t, err)
}
