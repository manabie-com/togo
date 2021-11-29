package services

import (
	"errors"
	"time"
	"togo/models"
	"togo/utils"
)

type UserReq struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRes struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

func toUserRes(user *models.User) *UserRes {
	var userRes = &UserRes{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
	}
	return userRes
}

func (obj *UserReq) Login() (*UserRes, error) {
	model := models.User{
		Email: obj.Email,
	}
	user, err := model.Login()
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(obj.Password, user.Password) {
		return nil, errors.New("Login failed")
	}
	return toUserRes(user), nil
}
