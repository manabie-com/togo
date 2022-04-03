package task

import (
	"net/http"
	"time"

	"github.com/TrinhTrungDung/togo/internal/model"
	"github.com/TrinhTrungDung/togo/pkg/server"
	"gorm.io/gorm"
)

var (
	ErrInvalidTaskAccess = server.NewHTTPError(http.StatusBadRequest, "INVALID_ACCESS", "You do not have permission to this task")
	ErrRetrieveTaskList  = server.NewHTTPInternalError("Cannot retrieve task list")
	ErrRetrieveTask      = server.NewHTTPInternalError("Cannot retrieve task")
	ErrCreateTask        = server.NewHTTPInternalError("Cannot creating task")
	ErrRetrieveUserSub   = server.NewHTTPInternalError("Cannot retrieve user subscription")
)

// List retrieves list of tasks for the current user
func (t *Task) List(authUser *model.AuthUser) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := t.db.Where(&model.Task{UserID: authUser.ID}).Find(&tasks).Error; err != nil {
		return nil, ErrRetrieveTaskList.SetInternal(err)
	}

	return tasks, nil
}

// View retrieves task based on id for the current user
func (t *Task) View(authUser *model.AuthUser, id int) (*model.Task, error) {
	task := &model.Task{}
	if err := t.db.Where("id = ? AND user_id = ?", id, authUser.ID).First(task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrInvalidTaskAccess
		}
		return nil, ErrRetrieveTask.SetInternal(err)
	}

	return task, nil
}

// Create creates new task for the current user
func (t *Task) Create(authUser *model.AuthUser, data TaskData) (*model.Task, error) {
	// Check if the user already uses the daily limit task yet
	sub := &model.Subscription{}
	if err := t.db.Preload("Plan").Where(&model.Subscription{UserID: authUser.ID}).Take(sub).Error; err != nil {
		return nil, ErrRetrieveUserSub.SetInternal(err)
	}

	// We use sliding window technique for counting number of tasks daily
	now := time.Now()
	var tasks []*model.Task
	if err := t.db.Where("user_id = ? AND created_at BETWEEN ? AND ? AND deleted IS NULL", authUser.ID, now.Add(-time.Duration(12)*time.Hour), now.Add(time.Duration(12)*time.Hour)).Find(&tasks).Error; err != nil {
		return nil, ErrRetrieveTaskList.SetInternal(err)
	}

	if len(tasks) >= sub.Plan.MaxTasks {
		return nil, server.NewHTTPError(http.StatusBadRequest, "TASK_LIMIT_REACHING", "You have used maximum daily limit of tasks")
	}

	task := &model.Task{
		Content: data.Content,
		UserID:  authUser.ID,
	}

	if err := t.db.Create(task).Error; err != nil {
		return nil, ErrCreateTask.SetInternal(err)
	}

	return task, nil
}

// Update updates the existed task of current user based on id
func (t *Task) Update(authUser *model.AuthUser, id int, data TaskData) (*model.Task, error) {
	task := &model.Task{}
	// Update task
	if err := t.db.Model(task).Where("id = ? AND user_id = ?", id, authUser.ID).Update("content", data.Content).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrInvalidTaskAccess
		}
		return nil, ErrRetrieveTask.SetInternal(err)
	}

	// Reload the updated task
	if err := t.db.Where("id = ?", id).Take(task).Error; err != nil {
		return nil, ErrRetrieveTask.SetInternal(err)
	}

	return task, nil
}

// Delete makes soft delete the existed task of current user based on id
func (t *Task) Delete(authUser *model.AuthUser, id int) error {
	task := &model.Task{}
	if err := t.db.Model(task).Where("id = ? AND user_id = ?", id, authUser.ID).Delete(task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrInvalidTaskAccess
		}
		return ErrRetrieveTask.SetInternal(err)
	}

	return nil
}
