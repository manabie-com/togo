package repository

import (
	"errors"
	"fmt"
	"github.com/manabie/project/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CheckAccount(user model.User) (string, error)
	CheckAccountExists(user model.User) error
	CreateAccount(user model.User) error
	CountTask(id int) int
	CreateTask(task model.Task) error
	UpdateTask(id int, task model.Task) error
	DeleteTask(id int) error
	TaskAll() ([]model.Task, error)
	TaskById(id int) (model.Task, error)
	TodoUser(id int) int
	UpdateTodoUser(id int, maxTodo int) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db:db,
	}
}

func(r *repository) CheckAccount(user model.User) (string, error) {
	var result model.User
	query := r.db.Where("username = ?", user.Username).Limit(1).Find(&result)

	if query.Error != nil{
		return "", errors.New("query error")
	}

	if result.Id == 0 {
		return "", errors.New("user does not exist")
	}
	return result.Password, nil
}

func(r *repository) CheckAccountExists(user model.User) error {
	var result model.User
	query := r.db.Where("username = ?", user.Username).Limit(1).Find(&result)
	if query.Error != nil{
		return errors.New("query error")
	}

	if result.Id != 0 {
		return errors.New("user does exist")
	}
	return nil
}

func(r *repository) CreateAccount(user model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func(r *repository) CountTask(id int) int {
	var result model.User
	err := r.db.Model(&model.User{}).Select("max_todo").Where("id = ?", id).First(&result)
	if err != nil {
		return 0
	}
	return result.MaxTodo
}

func(r *repository) CreateTask(task model.Task) error {
	if err := r.db.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func(r *repository) UpdateTask(id int, task model.Task) error {
	err := r.db.Model(model.Task{}).Where("id = ?",id).Updates(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func(r *repository) DeleteTask(id int) error {
	query := r.db.Delete(model.Task{},"id = ?", id)
	if query.RowsAffected == 0 {
		return errors.New("delete task failed")
	}
	return nil
}

func(r *repository) TaskAll() ([]model.Task, error) {
	var tags []model.Task
	err := r.db.Find(&tags).Error
	if err != nil{
		return nil, err
	}

	return tags, nil
}

func(r *repository) TaskById(id int) (model.Task, error) {
	var tag model.Task
	result := r.db.Where("id = ?",id).Find(&tag)
	if result.Error != nil{
		return model.Task{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Task{}, errors.New("todo task not found")
	}
	return tag, nil
}

func(r *repository) TodoUser(id int) int {
	var user model.User
	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return 0
	}
	fmt.Println(user.MaxTodo , ">>>>>>>")
	return user.MaxTodo
}

func(r *repository) UpdateTodoUser(id int, maxTodo int) error {
	err := r.db.Model(&model.User{}).Where("id = ?", id).Update("max_todo", maxTodo).Error
	if err != nil {
		return err
	}
	return nil
}