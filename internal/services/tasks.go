package services

import (
	"log"

	"github.com/manabie-com/togo/internal/storages"
)

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
	Jwt struct {
		Secret string
	}
}

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Config Config
	//pointer to logger standard library
	Logger *log.Logger
	Models storages.Models
}
