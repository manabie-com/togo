package repository

import (
	"ansidev.xyz/pkg/log"
	"github.com/ansidev/togo/domain/task"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/errs"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

func NewPostgresTaskRepository(db *gorm.DB) task.ITaskRepository {
	return &postgresTaskRepository{db}
}

type postgresTaskRepository struct {
	db *gorm.DB
}

func (r *postgresTaskRepository) GetTotalTasksByUserAndDate(userModel user.User, date time.Time) (int64, error) {
	year, month, day := date.Date()
	startTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(year, month, day, 23, 59, 59, 999999999, time.UTC)

	var totalTodayTask int64
	result := r.db.Model(&task.Task{}).Where("user_id = ? AND created_at BETWEEN ? AND ?", userModel.ID, startTime, endTime).Count(&totalTodayTask)

	if result.Error != nil {
		log.Error("Error while querying total tasks of user: ", result.Error)
		return -1, errors.Wrap(errs.ErrDatabaseFailure, result.Error.Error())
	}

	return totalTodayTask, nil
}

func (r *postgresTaskRepository) Create(taskModel task.Task, userModel user.User) (task.Task, error) {
	taskModel.UserID = userModel.ID
	taskModel.User = userModel

	result := r.db.Create(&taskModel)

	if result.Error != nil {
		log.Error("Error while creating task", result.Error)
		return task.Task{}, errors.Wrap(errs.ErrDatabaseFailure, result.Error.Error())
	}

	return taskModel, nil
}
