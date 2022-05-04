package handlers

import (
	"testing"

	"github.com/manabie-com/togo/models"
)

func TestGetUserById(t *testing.T) {

	togo1 := &models.Togo{Task: "testing1", Userid: 1}

	SetUpUnitTest(togo1)

	result, err := GetUserById(togo1)

	if err != nil {
		t.Errorf("There are error when execute GetUserById with parameter %v", togo1.Userid)
	}

	if result == nil || result.Id == 0 {
		t.Errorf("Can't found user when execute GetUserById with parameter %v", togo1.Userid)
	}

	cleanUnitTest(togo1)
}

func TestCreateUser(t *testing.T) {

	togo2 := &models.Togo{Task: "testing2", Userid: 1}

	SetUpUnitTest(togo2)

	deleteUser(togo2)

	result, err := CreateUser(togo2)

	if err != nil {
		t.Errorf("There are error when execute CreateUser with parameter %v", togo2.Userid)
	}

	if result == nil || result.Id == 0 {
		t.Errorf("Can't create user when execute CreateUser with parameter %v", togo2.Userid)
	}

	cleanUnitTest(togo2)
}
