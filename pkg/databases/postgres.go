package databases

import (
	"fmt"
	"togo/configs"
	"togo/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(postgreConfig *configs.PostgreSQLConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=%s",
		postgreConfig.Host,
		postgreConfig.User,
		postgreConfig.Name,
		postgreConfig.Password,
		postgreConfig.Port,
		postgreConfig.SslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.L.Sugar().Fatal(err)
	}
	return db
}
