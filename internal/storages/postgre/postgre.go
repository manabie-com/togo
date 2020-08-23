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
