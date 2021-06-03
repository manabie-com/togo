package model

import (
	"togo/validation"
)

type User struct {
	ID       string `pg:"id,type:text,pk" json:"id"`
	Password string `pg:"password,type:text" json:"password"`
	MaxTodo  int    `pg:"max_todo,type:text" json:"max_todo"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserById(user_id string) (user User, err error) {
	err = Db.Model(&user).Where(`id = ?`, user_id).First()
	if err != nil {
		return user, err
	}
	return user, nil
}

func Login(username, password string) (interface{}, error) {
	user := new(User)
	err := Db.Model(user).Where("id = ? AND password = ?", username, password).First()
	if err != nil {
		return nil, err
	}
	token, err := validation.GenerateToken(username)
	if err != nil {
		return nil, err
	}
	return token, nil
}
