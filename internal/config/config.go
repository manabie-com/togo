package config

import (
	"os"
	"strconv"
)

func ternary(value string, defaultValue string) string {
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// GET ENV VARIABLE FOR SECURITY CREDENTIALS
var (
	JWT_KEY                 = ternary(os.Getenv("JWTKey"), "wqGyEBBfPK9w3Lxw")
	HTTPPort, _             = strconv.Atoi(ternary(os.Getenv("HTTP_PORT"), "5050"))
	DataSource              = ternary(os.Getenv("DATA_SOURCE"), "host=localhost port=5440 user=manabie password=manabie dbname=togo sslmode=disable")

	LimitAllowTaskPerDay, _ = strconv.ParseUint(ternary(os.Getenv("LIMIT_ALLOW_TASK_BY_DAY"), "5"), 10, 32)
)
