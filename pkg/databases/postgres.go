package databases

import (
	"fmt"
	"togo/configs"
	"togo/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=%s",
		configs.C.PostgreSQL.Host,
		configs.C.PostgreSQL.User,
		configs.C.PostgreSQL.Name,
		configs.C.PostgreSQL.Password,
		configs.C.PostgreSQL.Port,
		configs.C.PostgreSQL.SslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.L.Sugar().Fatal(err)
	}
	return db
}
