package services

import (
	"log"

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
			log.Printf("Record successfully to user :%d; currently have %d tasks", request.UserId, len(user.TodoTasks))
			return nil
		}
	}

	return errors.GetError(errors.BadRequestContext, errors.BadRequestMessage, errors.UserIdNotFound)
}

func New() *Service {
	return &Service{
		users.CreateTempUsers(),
	}
}
