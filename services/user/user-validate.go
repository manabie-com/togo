package user

import (
	"log"
	"strconv"

	"github.com/manabie-com/backend/entity"
	"github.com/manabie-com/backend/repository"
	"github.com/manabie-com/backend/utils"
)

type I_UserServiceValidate interface {
	IsAllowedAddTask() *utils.ErrorRest
}

type UserServiceValidate struct {
	todayTaks []entity.Task
	userId    string
	limitTask int
}

var (
	repo repository.I_Repository
)

func NewUserServiceValidate(repository repository.I_Repository, userId string) (I_UserServiceValidate, *utils.ErrorRest) {
	repo = repository
	limitTask := 10

	cf, errConfig := utils.LoadConfig("../..")
	if errConfig != nil {
		log.Fatalf("Con: %v", errConfig)
	}
	limitTask, err := strconv.Atoi(cf.LIMIT_TASK_PER_DAY)

	if err != nil {
		log.Fatalf("Con: %v", errConfig)
	}

	today := utils.GetToday()
	tasks, errFind := repo.FindTaskByUserIdAndDay(userId, today)
	if err != nil {
		return nil, errFind
	}

	return &UserServiceValidate{
		todayTaks: tasks,
		userId:    userId,
		limitTask: limitTask,
	}, nil
}

func (uv *UserServiceValidate) IsAllowedAddTask() *utils.ErrorRest {

	if len(uv.todayTaks) >= uv.limitTask {
		messge := "The tasks is reached to " + strconv.Itoa(uv.limitTask) + " tasks per day"
		return utils.ErrBadRequest(messge)
	}

	return nil
}
