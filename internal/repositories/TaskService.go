package repo

import (
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

type TaskService interface {
	RetrieveTasks(userID, createdDate sql.NullString) ([]*storages.Task, error)
	AddTask(t *storages.Task, email string) error
	DeleteTask(id int) error
	UpdateTask(task *storages.Task) error
	ValidateUser(email, pwd sql.NullString) bool
	GetUserFromEmail(email string) (*storages.User, error)
}
