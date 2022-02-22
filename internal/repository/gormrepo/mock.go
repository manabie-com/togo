package gormrepo

import (
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newGormDBMock() (*sql.DB, sqlmock.Sqlmock, *gorm.DB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: true,
		Conn:                 db,
	}), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		return nil, nil, nil, fmt.Errorf("an error '%s' was not expected when create a gorm DB", err)
	}
	return db, mock, gdb, nil
}
