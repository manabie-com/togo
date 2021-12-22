package postgresql

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/storages/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgreSQLWrapper struct {
	Database *gorm.DB
}

func Connect(dsn string) *PostgreSQLWrapper {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Printf(config.CONNECTED_TO, "PostgreSQL Server")
	return &PostgreSQLWrapper{db}
}

func GetDb(ctx *gin.Context) (*PostgreSQLWrapper, bool) {
	poster, ok := ctx.Get(config.POSTGRESQL_DB)
	return poster.(*PostgreSQLWrapper), ok
}

func MakeMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})
}
