package utils

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func SafeString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func String(s string) *string {
	return &s
}

// --Remove Duplicate Data
func RemoveDuplicate(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

type Environment struct {
	Host        string `env:"HOST"`
	Port        string `env:"PORT"`
	JwtKey      string `env:"JWT_KEY"`
	JwtTimeout  string `env:"JWT_TIMEOUT"`
	DBDriver    string `env:"DB_DRIVER"`
	DBHost      string `env:"DB_HOST"`
	DBPort      string `env:"DB_PORT"`
	DBUser      string `env:"DB_USER"`
	DBPassword  string `env:"DB_PASSWORD"`
	DBName      string `env:"DB_NAME"`
	DBSslMode   string `env:"DB_SSL_MODE"`
	DSNPostgres string `env:"DSN_POSTGRES"`
}

var Env *Environment

func LoadEnv(file string) {
	err := godotenv.Load(file)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	if Env == nil {
		Env = new(Environment)
	}

	getEnvironmentFromEnv(Env)
}

func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Missing or invalid environment key: '%s'", key)
	}
	return value
}

func getEnvironmentFromEnv(object *Environment) {
	if object == nil {
		return
	}
	env.Parse(object)
}
