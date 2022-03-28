package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	env := os.Getenv("TOGO_ENV")
	if env == "" || env == "dev" {
		godotenv.Load()
	}
}
