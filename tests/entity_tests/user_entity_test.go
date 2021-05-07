package entity_tests

import (
	"github.com/go-playground/assert/v2"
	"log"
	"manabie-com/togo/entity"
	"manabie-com/togo/query"
	"manabie-com/togo/util"
	"testing"
)

func TestSaveUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatalf("Error user refreshing table %v\n", err)
	}
	newUser := entity.User{
		ID:       "firstUser",
		Password: util.HashPassword("firstUser", "password"),
		MaxTodo:  5,
	}
	err = newUser.Create()
	if err != nil {
		t.Errorf("Error while saving a user: %v\n", err)
		return
	}
	assert.Equal(t, newUser.ID, "firstUser")
	assert.Equal(t, newUser.MaxTodo,5)
}

func TestGetUserByID(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatalf("Error user refreshing table %v\n", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	foundUser, err := query.UserByID(user.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.MaxTodo, user.MaxTodo)
}