// todofixture for testing
package todofixture

import (
	"errors"

	"github.com/laghodessa/togo/domain"
	"github.com/laghodessa/togo/domain/todo"
)

// NewUser returns a user fixture with optional override.
//
// Use default
//  user := NewUser()
//
// Override default
//  user := NewUser(func(user *todo.User) {
//  	user.TaskDailyLimit = 3
//  })
func NewUser(override ...func(*todo.User)) todo.User {
	user := todo.User{
		ID:             domain.NewID(),
		TaskDailyLimit: 10,
	}

	// override if neccessary
	if len(override) > 1 {
		panic(errors.New("override must be exactly one function"))
	}
	if len(override) > 0 {
		override[0](&user)
	}
	return user
}
