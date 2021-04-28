package usecase

import (
	"errors"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/validate"
)

type useCase struct {
	repository storages.EntityRepository
}

func InitUseCase(repo storages.EntityRepository) storages.EntityUseCase {
	return &useCase{
		repository: repo,
	}
}

func (u *useCase) GetAuthToken(args map[string]string) (map[string]interface{}, error) {

	var (
		res = map[string]interface{}{}
		err error
	)

	err = u.repository.Login(args["user_id"], args["password"])

	if err != nil {
		return res, errors.New("incorrect user_id/pwd")
	}

	res["token"], err = validate.CreateToken(args["user_id"])

	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *useCase) ListTasks(args map[string]string) (map[string]interface{}, error) {

	var (
		res = map[string]interface{}{}
		err error
	)

	tasks, err := u.repository.GetTasks(args["user_id"], args["created_date"])
	if err != nil {
		return res, err
	}

	res["tasks"] = tasks
	return res, nil
}

func (u *useCase) AddTask(task storages.Task) (map[string]interface{}, error) {

	var (
		chUser                 = make(chan storages.User)
		chTasks                = make(chan []storages.Task)
		res                    = map[string]interface{}{}
		errInsert, errValidate error
	)

	go func() {
		var user storages.User
		user, errValidate = u.repository.GetUserByID(task.UserID)
		chUser <- user
	}()

	go func() {
		var tasks []storages.Task
		tasks, errValidate = u.repository.GetTasks(task.UserID, task.CreatedDate)
		chTasks <- tasks
	}()

	if errValidate != nil {
		res["status"] = 400
		return res, errValidate
	}

	user := <-chUser
	tasks := <-chTasks

	if len(tasks) >= user.MaxTodo {
		res["status"] = 400
		return res, errors.New("The number of task is full")
	}

	errInsert = u.repository.InsertTask(&task)

	if errInsert != nil {
		res["status"] = 500
		return res, errInsert
	}

	res["task"] = task

	return res, nil
}
