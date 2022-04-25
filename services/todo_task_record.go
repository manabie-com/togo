package services

import (
	"log"
	"time"

	"github.com/SVincentTran/togo/errors"
	"github.com/SVincentTran/togo/forms"
	"github.com/SVincentTran/togo/users"
)

type Service struct {
	users []*users.User
}

func (s Service) RecordTodoTasks(request forms.TodoTaskRequest) error {

	for _, user := range s.users {
		// Check if request User Id is in the temp list
		if user.UserId == request.UserId {
			// Check if the last updated time is not the same as now()
			now := time.Now()
			// If not then reset the daily task limit
			if !DateEqual(user.LastUpdated, now) {
				user.DailyTask = 0
			}

			// Check if the current number of recorded task is exceed task limit or not
			if user.DailyTask >= user.TaskLimit {
				log.Printf("This user has exceeded the tasks limit!")
				return errors.GetError(errors.BadRequestContext, errors.BadRequestMessage, errors.ExceedDailyLimitRecords)
			}

			// Record new to do task
			newTodoTask := users.TodoTask{
				Title:      request.Title,
				Detail:     request.Detail,
				RemindDate: request.RemindDate,
			}
			user.TodoTasks = append(user.TodoTasks, newTodoTask)
			user.DailyTask += 1

			// Update last updated time
			user.LastUpdated = now
			log.Printf("Record successfully to user :%d; currently have %d tasks", request.UserId, len(user.TodoTasks))
			return nil
		}
	}

	return errors.GetError(errors.BadRequestContext, errors.BadRequestMessage, errors.UserIdNotFound)
}

// DateEqual take two Time and compare if they are the same date or not
func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func New() *Service {
	return &Service{
		users.CreateTempUsers(),
	}
}
