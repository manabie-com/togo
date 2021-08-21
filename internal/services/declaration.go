package services

import "github.com/jinzhu/gorm"

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  *gorm.DB
}

type userAuthKey int8
