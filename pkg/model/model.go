package model

import "github.com/golang-jwt/jwt"

type LoginRequest struct {
	UserName string `json:"username" query:"username"`
	Password string `json:"password" query:"password"`
}

type GetTaskRequest struct {
	CreateDate string `json:"create_date" query:"create_date"`
}

type AddTaskRequest struct {
	Content string `json:"content" query:"content"`
}
type AddTaskParams struct {
	UserId     string
	CreateDate string
	Content    string
}

type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

type AddTaskResponse struct {
	UserId     string `json:"user_id"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type User struct {
	Id      string `json:"id"`
	MaxToDo int64  `json:"max_to_do"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
