package service

import (
	"errors"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
	"togo/database"
	"togo/domain"
)

type TaskService struct {
	repo     domain.ITaskRepository
	userRepo domain.IUserRepository
}

func NewTaskService(repo domain.ITaskRepository, userRepo domain.IUserRepository) domain.ITaskService {
	return &TaskService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (service *TaskService) Create(params domain.TaskParams) (domain.Task, error) {
	task := domain.Task{
		Content:   params.Content,
		CreatedAt: time.Time{},
	}

	tx := database.DB.Begin()
	service.repo.SetTx(tx)
	service.userRepo.SetTx(tx)
	createdTask, dbErr := func() (createdTask domain.Task, err error) {
		createdTask, err = service.repo.Create(task)
		if err != nil {
			return
		}

		if params.UserId != 0 {
			user, findErr := service.userRepo.FindById(params.UserId)
			if errors.Is(findErr, gorm.ErrRecordNotFound) {
				err = domain.ErrUserNotFound
				return
			}

			err = service.AssignToUser(createdTask, user)
			if err != nil {
				return
			}
		}

		if params.UserEmail != "" {
			user, findErr := service.userRepo.FindByEmail(params.UserEmail)
			if errors.Is(findErr, gorm.ErrRecordNotFound) {
				defaultTaskLimit, _ := strconv.Atoi(os.Getenv("DEFAULT_TASK_LIMIT"))
				user = domain.User{
					Email:     params.UserEmail,
					TaskLimit: defaultTaskLimit,
					CreatedAt: time.Time{},
				}

				if params.TaskLimit != 0 {
					user.TaskLimit = params.TaskLimit
				}
				user, err = service.userRepo.Create(user)
				if err != nil {
					return
				}
			}
			err = service.AssignToUser(createdTask, user)
			if err != nil {
				return
			}
		}
		return
	}()

	if dbErr != nil {
		tx.Rollback()
		return domain.Task{}, dbErr
	}
	tx.Commit()
	return createdTask, nil
}

func (service *TaskService) AssignToUser(task domain.Task, user domain.User) error {
	if user.TaskLimit <= 0 {
		return domain.ErrExceedTaskLimit
	}

	task.UserId = &user.ID
	err := service.repo.Update(task)
	if err != nil {
		return err
	}
	user.TaskLimit--
	err = service.userRepo.Save(user)
	if err != nil {
		return err
	}
	return nil
}
