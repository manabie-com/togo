package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model

	Content     string `json:"content"`
	UserID      uint `json:"user_id"`
	CreatedTask string `json:"created_task"`
}

func (obj *Task) Get(id uint) (*Task, error) {
	err := db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}

	return obj, err
}

func (obj *Task) GetAll(pageNum, pageSize int, maps interface{}) ([]*Task, error) {
	var list []*Task
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return list, nil
}

func (obj *Task) GetTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Task{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (obj *Task) GetCountTaskByUser(userId uint, createAt string) (int, error) {
	var count int
	condition := map[string]interface{}{
		"user_id": userId,
		"created_task": createAt,
	}
	if err := db.Model(&Task{}).Where(condition).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (obj *Task) Add() (*Task, error) {
	if err := db.Create(&obj).Error; err != nil {
		return nil, err
	}
	obj, err = obj.Get(obj.ID)
	if err != nil {
		return nil, err
	}

	return obj, err
}

func (obj *Task) Update(id uint) (*Task, error) {
	var tmpObj Task
	err := db.Where("id = ?", id).First(&tmpObj).Error
	if err != nil {
		return nil, err
	}
	db.Model(&tmpObj).Update(obj)
	//Response
	resObj, err := tmpObj.Get(id)
	if err != nil {
		return nil, err
	}

	return resObj, err
}

func (obj *Task) Delete(id uint) (*Task, error) {
	resObj, err := obj.Get(id)
	if err != nil {
		return nil, err
	}
	if resObj.ID > 0 {
		if err := db.Where("id = ?", resObj.ID).Delete(&resObj).Error; err != nil {
			return nil, err
		}
		return resObj, nil
	}

	return nil, errors.New("Task does not exist!")
}
