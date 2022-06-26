package task

import (
	"fmt"

	"github.com/manabie-com/togo/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func setupMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New() // mock sql.DB
	if err != nil {
		return nil, nil, errors.Wrap(err, "Fail connection to database")
	}

	db, err := gorm.Open("postgres", sqlDB) // open gorm db
	if err != nil {
		return nil, nil, err
	}

	if db == nil {
		return nil, nil, errors.New("Fail connection to database")
	}

	return db, mock, nil
}

func recordStats(db *gorm.DB, task models.Task) (err error) {
	_, err = db.CommonDB().Exec("INSERT INTO tasks (id, content, create_date, user_id) VALUES (?, ?,?,?,?)", 1, task.Content, task.CreateDate, task.UserID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error '%s' was not expected, while inserting a row", err))
	}
	return
}
