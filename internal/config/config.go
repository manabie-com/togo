package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config struct
type Config struct {
	Env            string
	Port           string
	DbDriver       string
	DbHost         string
	DbUser         string
	DbPassword     string
	DbName         string
	DbPort         string
	JwtKey         string
	JwtExp         int
	MaxTodoDefault int
}

// Cfg config
var Cfg = Config{}

// Load func
func Load() error {
	err := godotenv.Load()
	if err == nil {
		jwtExp, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
		MaxTodoDefault, _ := strconv.Atoi(os.Getenv("MAX_TODO_DEFAULT"))
		Cfg = Config{
			Env:            os.Getenv("ENV"),
			Port:           os.Getenv("PORT"),
			DbDriver:       os.Getenv("DB_DRIVER"),
			DbHost:         os.Getenv("DB_HOST"),
			DbUser:         os.Getenv("DB_USER"),
			DbPassword:     os.Getenv("DB_PASSWORD"),
			DbName:         os.Getenv("DB_NAME"),
			DbPort:         os.Getenv("DB_PORT"),
			JwtKey:         os.Getenv("JWT_KEY"),
			JwtExp:         jwtExp,
			MaxTodoDefault: MaxTodoDefault,
		}
	}
	return err
}
