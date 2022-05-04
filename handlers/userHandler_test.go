package handlers

import (
	"testing"

	"github.com/manabie-com/togo/models"
)

func TestGetUserById(t *testing.T) {

	todo1 := &models.Togo{Task: "testing1", Userid: 1}

	SetUpUnitTest(todo1)

	result, err := GetUserById(todo1)

	if err != nil {
		t.Errorf("There are error when execute GetUserById with parameter %v", todo1.Userid)
	}

	if result == nil || result.Id == 0 {
		t.Errorf("Can't found user when execute GetUserById with parameter %v", todo1.Userid)
	}

	cleanUnitTest(todo1)
}

func TestCreateUser(t *testing.T) {

	todo2 := &models.Togo{Task: "testing2", Userid: 1}

	SetUpUnitTest(todo2)

	deleteUser(todo2)

	result, err := CreateUser(todo2)

	if err != nil {
		t.Errorf("There are error when execute CreateUser with parameter %v", todo2.Userid)
	}

	if result == nil || result.Id == 0 {
		t.Errorf("Can't create user when execute CreateUser with parameter %v", todo2.Userid)
	}

	cleanUnitTest(todo2)
}
