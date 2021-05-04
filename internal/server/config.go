package server

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var Config appConfig

type appConfig struct {
	Port            int
	DBUrl           string
	DBMaxIdleTime   int
	DBMaxLifeTime   int
	DBMaxConnection int
	DBMinConnection int
}

func getEnvString(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
func getEnvInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		i1, err := strconv.Atoi(value)
		if err == nil {
			return i1
		}
	}
	return defaultVal
}

func loadConfig() {
	Config = appConfig{
		Port:            getEnvInt("PORT", 5050),
		DBUrl:           getEnvString("DATABASE_URL", ""),
		DBMaxIdleTime:   getEnvInt("DATABASE_MAX_IDLE_TIME", 3600),
		DBMaxLifeTime:   getEnvInt("DATABASE_MAX_LIFE_TIME", 1800),
		DBMaxConnection: getEnvInt("DATABASE_MIN_CON", 0),
		DBMinConnection: getEnvInt("DATABASE_MAX_CON", 5),
	}
}

func InitServerConfig(path string) {
	// load application configurations
	if path == "" {
		path = ".env"
	}
	if err := godotenv.Load(path); err != nil {
		log.Println("No .env file found")
	}
	loadConfig()
}
