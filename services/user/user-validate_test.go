package user

import (
	"errors"
	"testing"

	"github.com/manabie-com/backend/entity"
	"github.com/manabie-com/backend/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

const (
	userId = "userid-1"
)

func CreateTask() entity.Task {

	return entity.Task{
		ID:          uuid.NewV4().String(),
		UserID:      userId,
		Content:     "Task-" + uuid.NewV4().String(),
		Status:      "pendding",
		CreatedDate: utils.GetToday(),
	}
}
func TestUserValidate(t *testing.T) {
	tenTasks := make([]entity.Task, 10)

	for i := 0; i < 10; i++ {
		newTask := CreateTask()
		tenTasks = append(tenTasks, newTask)
	}

	userValidate := UserServiceValidate{
		todayTaks: tenTasks,
		userId:    userId,
		limitTask: 5,
	}

	assert.Error(t, errors.New("The tasks is reached to 10 tasks per day"), userValidate.IsAllowedAddTask())

	zeroTasks := make([]entity.Task, 0)
	userValidate = UserServiceValidate{
		todayTaks: zeroTasks,
		userId:    userId,
		limitTask: 5,
	}

	assert.Nil(t, nil, userValidate.IsAllowedAddTask())
}
