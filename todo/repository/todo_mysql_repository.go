package repository

import (
	"github.com/triet-truong/todo/todo/model"
	"github.com/triet-truong/todo/todo/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoMysqlRepository struct {
	db *gorm.DB
}

func NewTodoMysqlRepository(dsn string) *TodoMysqlRepository {
	mysqlGorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.ErrorLog(err)
	return &TodoMysqlRepository{
		db: mysqlGorm,
	}
}

func (r *TodoMysqlRepository) SelectUser(id uint) (model.UserModel, error) {
	var user model.UserModel
	err := r.db.First(&user, id).Error
	utils.ErrorLog(err)
	return user, err
}

func (r *TodoMysqlRepository) InsertItem(item model.TodoItemModel) error {
	err := r.db.Create(&item).Error
	utils.ErrorLog(err)
	return err
}
