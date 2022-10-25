package models

import (
	"github.com/jinzhu/gorm"
	"togo/utils"
)

type User struct {
	gorm.Model

	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (obj *User) IsEmailExist() (bool, error) {
	var tmpObj User
	err := db.Where("email = ?", obj.Email).First(&tmpObj).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tmpObj.ID > 0 {
		return true, nil
	}
	return false, nil
}

func (obj *User) Add() (*User, error) {
	if err := db.Create(&obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (obj *User) Get(id uint) (*User, error) {
	err := db.Where("id = ? ", id).First(&obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (obj *User) GetTotal() (int, error) {
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (obj *User) Login() (*User, error) {
	var tmpObj User
	err := db.Where("email = ? ", obj.Email, true).First(&tmpObj).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	db.Model(&tmpObj).Update(obj)
	//response
	resObj, err := tmpObj.Get(tmpObj.ID)
	if err != nil {
		return nil, err
	}
	return resObj, nil
}

func (obj *User) IsExist(id uint) (bool, error) {
	err := db.Where("id = ?", id).Find(&obj).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if obj.ID > 0 {
		return true, nil
	}
	return false, nil
}

func (obj *User) MigrateData() {
	if total, _ := obj.GetTotal(); total == 0 {
		passHash, _ := utils.HashPassword("123456")

		obj = &User{
			FullName: "Admin",
			Email:    "admin@gmail.com",
			Password: passHash,
		}
		_, _ = obj.Add()
	}
}
