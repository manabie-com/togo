package services

import (
	"github.com/manabie-com/togo/internal/repositories"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  *repositories.LiteDB
}

