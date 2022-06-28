package domain

import (
	e "lntvan166/togo/internal/entities"
	"net/http"
)

type TaskRepository interface {
	CreateTask(t *e.Task) error
	GetAllTask() (*[]e.Task, error)
	GetTaskByID(id int) (*e.Task, error)
	GetTasksByUserID(id int) (*[]e.Task, error)
	GetNumberOfTaskTodayByUserID(id int) (int, error)
	GetMaxTaskByUserID(id int) (int, error)
	UpdateTask(t *e.Task) error
	CompleteTask(id int) error
	DeleteTask(id int) error
	// DeleteAllTask() error
	DeleteAllTaskOfUser(userID int) error
}

type TaskUsecase interface {
	CreateTask(task *e.Task, username string) (int, error)
	GetAllTask() (*[]e.Task, error)
	GetTaskByID(id int, username string) (*e.Task, error)
	GetTasksByUsername(username string) (*[]e.Task, error)
	CheckLimitTaskToday(id int) (bool, error)
	UpdateTask(id int, username string, r *http.Request) error
	CompleteTask(id int, username string) error
	DeleteTask(id int, username string) error
}

type UserRepository interface {
	AddUser(u *e.User) error
	GetAllUsers() ([]*e.User, error)
	GetUserByName(username string) (*e.User, error)
	GetUserByID(id int) (*e.User, error)
	GetUserIDByUsername(username string) (int, error)
	GetPlanByID(id int) (string, error)
	UpdateUser(u *e.User) error
	// UpgradePlan(id int, plan string, limit int) error
	DeleteUserByID(id int) error
}

type UserUsecase interface {
	Register(user *e.User) error
	Login(user *e.User) (string, error)
	GetAllUsers() ([]*e.User, error)
	GetUserByID(id int) (*e.User, error)
	GetUserIDByUsername(username string) (int, error)
	GetMaxTaskByUserID(id int) (int, error)
	GetPlan(username string) (string, error)
	UpdateUser(u *e.User) error
	UpgradePlan(userID int, plan string, maxTodo int) error
	DeleteUserByID(id int) error
	CheckUserExist(username string) bool
}
