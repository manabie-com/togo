package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserId uint
	jwt.StandardClaims
}
type User struct {
	ID            uint32 `json:"id"`
	Email         string `json:"email" validate:"required, email"`
	Name          string `json:"name" validate:"required"`
	Password      string `json:"password" validate:"required"`
	IsPayment     bool   `json:"isPayment"`
	LimitDayTasks uint   `json:"limitDayTasks"`
}
