package postgre

import (
	"github.com/go-pg/pg/v10"

	entity "github.com/manabie-com/togo/internal/entities"
)

// Storage working with postgre drive
type Storage struct {
	DB *pg.DB
}

// Add adds a new task to DB
func (pgres *Storage) Add(entity *entity.Task) (*entity.Task, error) {

	_, err := pgres.DB.Model(entity).Insert()

	return entity, err
}

// GetAll retrives tasks
func (pgres *Storage) GetAll(createdDate string) ([]entity.Task, error) {

	var tasks []entity.Task

	err := pgres.DB.Model(&tasks).Where("created_date = (?)", createdDate).Select()

	return tasks, err
}

// GetByID get task by id
func (pgres *Storage) GetByID(taskID string) (*entity.Task, error) {

	task := &entity.Task{}

	err := pgres.DB.Model(task).Where("id = (?)", taskID).Select()

	return task, err
}

// ValidateUser validate users info
func (pgres *Storage) ValidateUser(username string, password string) (*entity.User, error) {

	user := &entity.User{}

	count, err := pgres.DB.Model(user).Where("username = (?)", username).Where("password = (?)", password).SelectAndCount()

	if count > 0 {
		defer pgres.DB.Model(user).Where("username=(?)", username).Select()
	}

	return user, err
}

// GetByUserID get task by userID
func (pgres *Storage) GetByUserID(userID string, createdDate string) ([]entity.Task, error) {

	var tasks []entity.Task

	err := pgres.DB.Model(&tasks).Where("user_id = (?)", userID).Where("created_date = (?)", createdDate).Select()

	return tasks, err
}
