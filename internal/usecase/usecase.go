package usecase

import (
	"errors"
	"github.com/manabie/project/internal/repository"
	"github.com/manabie/project/model"
	"github.com/manabie/project/pkg/hash"
	"github.com/manabie/project/pkg/jwt"
)

type usecase struct {
	repo repository.Repository
	jwt  jwt.TokenUser
	hash hash.Hash
}

type Usecase interface {
	Login(user model.User) (string , error)
	SignUp(user model.User) error
	CreateTask(task model.Task, idUser int) error
	UpdateTask(id int, task model.Task) error
	DeleteTask(id int) error
	TaskAll() ([]model.Task, error)
	TaskById(id int) (model.Task, error)
}

func NewUsecase(repo repository.Repository, jwt jwt.TokenUser, hash hash.Hash) Usecase {
	return &usecase{
		repo: repo,
		jwt: jwt,
		hash: hash,
	}
}

func(u *usecase) Login(user model.User) (string , error) {
	hashPassword,err  := u.repo.CheckAccount(user)
	if err != nil {
		return "", errors.New("user does not exist")
	}
	if checkPass := u.hash.CheckPassword(user.Password, hashPassword); checkPass != nil {
		return "" ,errors.New("password entered incorrectly")
	}
	token, _ := u.jwt.GenerateToken(user.Username)
	return token, nil
}

func(u *usecase) SignUp(user model.User) error {
	check := u.repo.CheckAccountExists(user)
	if check != nil {
		return errors.New("user already exists")
	}
	hashPassword, err := u.hash.HashPassword(user.Password)
	if err != nil {
		return errors.New("password failed")
	}
	account:= model.User{
		Id:           user.Id,
		Username:     user.Username,
		Password:     hashPassword,
	}
	err = u.repo.CreateAccount(account)
	if err != nil {
		return errors.New("create failed")
	}
	return nil
}

func(u *usecase) CreateTask(task model.Task, idUser int) error {
	countMaxTodo := u.repo.CountTask(idUser)
	if countMaxTodo == 5 {
		return errors.New("user created task exceeds 5")
	}

	maxTodo := u.repo.TodoUser(idUser)

	if err := u.repo.UpdateTodoUser(idUser, maxTodo + 1); err != nil {
		return err
	}

	if err := u.repo.CreateTask(task); err != nil {
		return errors.New("create task failed")
	}
	return nil
}

func(u *usecase) UpdateTask(id int, task model.Task) error {
	if err := u.repo.UpdateTask(id, task); err != nil {
		return err
	}
	return nil
}

func(u *usecase) DeleteTask(id int) error {
	if err := u.repo.DeleteTask(id); err != nil {
		return err
	}
	return nil
}

func(u *usecase) TaskAll() ([]model.Task, error) {
	tasks, err := u.repo.TaskAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func(u *usecase) TaskById(id int) (model.Task, error) {
	task, err := u.repo.TaskById(id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}