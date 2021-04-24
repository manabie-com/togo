package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}
type Jwt struct {
	SecretKey []byte
}

type Config struct {
	Database Database
	Jwt      Jwt
}

func LoadConfigs() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		Database{os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		}, Jwt{[]byte(os.Getenv("SECRET_KEY"))}}
}
