package env

import (
	"ManabieProject/helper"
	"os"
)

func Init() {
	os.Setenv("DB_HOST", "0.0.0.0")
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_USERNAME", "")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_NAME", "mongodb")

	// db strucure
	os.Setenv("DB", "Manabie")
	os.Setenv("COLLECTION", "Account")

	// JWTMaker key init
	secretKey, err := helper.GenKeyAEStoBase64()
	if secretKey != "" && err == nil {
		os.Setenv("SECRETKEY", secretKey)
		os.Setenv("DURATION", "10")
	}
}

func Value(key string) string {
	return os.Getenv(key)
}
