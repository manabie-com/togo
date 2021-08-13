package impl

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/repository"
	"github.com/manabie-com/togo/internal/repository/impl"
	"strconv"
	"time"
)

type TaskServiceImpl struct{
	repository repository.TasksRepository
	userRespository repository.UsersRepository
}

func NewTaskServiceImpl(db *gorm.DB) *TaskServiceImpl {
	return &TaskServiceImpl{impl.NewTaskRepositoryImpl(db), impl.NewUserRepositoryImpl(db)}
}

func (s *TaskServiceImpl) RetrieveTasks(id string, createdDate string) (model.TaskList, error) {

	return s.repository.GetByIdAndCreateDate(id, createdDate)
}

func (s *TaskServiceImpl) AddTask(task *model.Task, userID string) (*model.Task, error) {
	user, err := s.userRespository.GetUserById(userID)
	if err != nil{
		return nil, err
	}
	now := time.Now()
	task.CreatedDate = now.Format("2006-01-02")
	count, err := s.repository.CountByIdAndCreateDate(userID, task.CreatedDate)
	maxTodo, err := strconv.Atoi(user.MaxTodo)
	if err != nil {
		return nil, err
	}
	if err != nil{
		return nil, err
	} else if count >= maxTodo {

		return nil, errors.New(fmt.Sprintf("User insert greater than %s records" , user.MaxTodo))
	}
	task.ID = uuid.New().String()
	task.UserID = userID
	return s.repository.Save(task)
}