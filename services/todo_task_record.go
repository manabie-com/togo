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
		if user.UserId == request.UserId {
			now := time.Now()
			log.Printf("This Date %v", user.LastUpdated)
			if !DateEqual(user.LastUpdated, now) {
				user.DailyTask = 0
			}

			if user.DailyTask >= user.TaskLimit {
				log.Printf("This user has exceeded the tasks limit!")
				return errors.GetError(errors.BadRequestContext, errors.BadRequestMessage, errors.ExceedDailyLimitRecords)
			}
			newTodoTask := users.TodoTask{
				Title:      request.Title,
				Detail:     request.Detail,
				RemindDate: request.RemindDate,
			}
			user.TodoTasks = append(user.TodoTasks, newTodoTask)
			user.DailyTask += 1
			user.LastUpdated = now
			log.Printf("Record successfully to user :%d; currently have %d tasks", request.UserId, len(user.TodoTasks))
			return nil
		}
	}

	return errors.GetError(errors.BadRequestContext, errors.BadRequestMessage, errors.UserIdNotFound)
}

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
