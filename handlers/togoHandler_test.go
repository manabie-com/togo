package handlers

import (
	"testing"

	"github.com/manabie-com/togo/models"
)

func TestAddTodo(t *testing.T) {

	todo1 := &models.Togo{Task: "testing1", Userid: 1}

	SetUpUnitTest(todo1)

	user, _ := GetUserById(todo1)

	todo1.Userid = user.Id

	result, err := AddTodo(todo1)

	if err != nil {
		t.Errorf("There are error when add todo with parameter %v", todo1.Task)
	}

	if result == nil || result.CountTasks() == 0 {
		t.Errorf("Can't update count task with parameter %v", len(result.Tasks))
	}

	deleteTodo(todo1)

	cleanLimitTask(todo1)

	cleanUnitTest(todo1)
}

func TestAddTodoLimitTask(t *testing.T) {

	todo2 := &models.Togo{Task: "testing2", Userid: 1}

	SetUpUnitTest(todo2)

	GetUserById(todo2)

	user, _ := GetUserById(todo2)

	todo2.Userid = user.Id

	for i := 0; i <= 10; i++ {
		result, err := AddTodo(todo2)
		if user.CountTasks() > 10 {
			if err == nil {
				t.Errorf("Limit task working incorrect %v", result.CountTasks())
			}
		}
	}

	deleteTodo(todo2)

	cleanLimitTask(todo2)

	cleanUnitTest(todo2)
}
