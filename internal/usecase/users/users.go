package users

import (
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/reqres/users"
)

// AssignUserTasks assign tasks for user
func AssignUserTasks(userID int16, request *users.AssignTaskRequest) error {
	// check user is exist
	user, err := model.NewUser().
		GetUserTasksTodayByUserID(userID)
	if err != nil {
		return err
	}

	task := model.NewTask()
	// check all task ids is not assign yet
	isNotAssign, err := task.IsNotAssign(request.TaskIDs)
	if err != nil {
		return err
	}
	if !isNotAssign {
		return model.ErrSomeTasksAreNotSatisfying
	}

	// check user is exceeding task limit
	if len(user.Tasks)+len(request.TaskIDs) > int(user.LimitTaskPerDay) {
		return model.ErrExceedingTaskLimit
	}

	// update tasks for user
	return task.Assign(userID, request.TaskIDs)
}
