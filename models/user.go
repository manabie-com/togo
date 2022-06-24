package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId        uint32
	LimitDayTasks uint
	jwt.StandardClaims
}
type User struct {
	ID            uint32 `json:"id" validate:"omitempty"`
	Email         string `json:"email" validate:"required,email"`
	Name          string `json:"name" validate:"required,min=5,max=20"`
	Password      string `json:"password" validate:"required,min=6,max=20"`
	IsPayment     bool   `json:"isPayment" validate:"omitempty"`
	IsActive      bool   `json:"isActive"`
	LimitDayTasks uint   `json:"limitDayTasks" validate:"omitempty"`
}
