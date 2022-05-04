package handlers

import (
	"testing"

	"github.com/manabie-com/togo/models"
)

func TestAddtogo(t *testing.T) {

	togo1 := &models.Togo{Task: "testing1", Userid: 1}

	SetUpUnitTest(togo1)

	user, _ := GetUserById(togo1)

	togo1.Userid = user.Id

	result, err := Addtogo(togo1)

	if err != nil {
		t.Errorf("There are error when add togo with parameter %v", togo1.Task)
	}

	if result == nil || result.CountTasks() == 0 {
		t.Errorf("Can't update count task with parameter %v", len(result.Tasks))
	}

	deletetogo(togo1)

	cleanLimitTask(togo1)

	cleanUnitTest(togo1)
}

func TestAddtogoLimitTask(t *testing.T) {

	togo2 := &models.Togo{Task: "testing2", Userid: 1}

	SetUpUnitTest(togo2)

	GetUserById(togo2)

	user, _ := GetUserById(togo2)

	togo2.Userid = user.Id

	for i := 0; i <= 10; i++ {
		result, err := Addtogo(togo2)
		if user.CountTasks() > 10 {
			if err == nil {
				t.Errorf("Limit task working incorrect %v", result.CountTasks())
			}
		}
	}

	deletetogo(togo2)

	cleanLimitTask(togo2)

	cleanUnitTest(togo2)
}
