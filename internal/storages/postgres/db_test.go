package postgres

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cuongtop4598/togo-interview/togo/internal/helper"
	"github.com/cuongtop4598/togo-interview/togo/internal/storages"
	"github.com/google/uuid"
	"github.com/jinzhu/now"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *sql.DB
	err    error
	mock   sqlmock.Sqlmock
	gormDB *gorm.DB
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestValidateUser(t *testing.T) {
	db, mock = NewMock()
	defer db.Close()
	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	dbmanager := DBmanager{gormDB}

	query := `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."password" = $2 ORDER BY "users"."id" LIMIT 1`

	rows := sqlmock.NewRows([]string{"id", "password", "max_todo"}).AddRow("cuongnm", "123456", 5)

	mock.ExpectQuery(query).WithArgs("cuongnm", "123456").WillReturnRows(rows)

	result1 := dbmanager.ValidateUser(context.Background(), sql.NullString{String: "cuongnm", Valid: true}, sql.NullString{String: "123456", Valid: true})
	result2 := dbmanager.ValidateUser(context.Background(), sql.NullString{String: "xxxxxxx", Valid: true}, sql.NullString{String: "123456", Valid: true})
	assert.NotNil(t, result1)
	assert.Equal(t, true, result1)
	assert.Equal(t, false, result2)
}

func TestAddTask(t *testing.T) {
	db, mock = NewMock()
	defer db.Close()
	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	dbmanager := DBmanager{gormDB}
	tasks := []storages.Task{
		{
			ID:          uuid.MustParse("c9e4075a-0361-45c6-9191-bbd04860f58a"),
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   gorm.DeletedAt{},
			Content:     "hello1",
			UserID:      "cuongnm",
			CreatedDate: "2021-10-8",
		},
		{
			ID:          uuid.MustParse("c5e4075a-0361-45c6-9191-bbd04860f58a"),
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   gorm.DeletedAt{},
			Content:     "hello1",
			UserID:      "cuongnm",
			CreatedDate: "2021-10-8",
		},
		{
			ID:          uuid.MustParse("c1e4075a-0361-45c6-9191-bbd04860f58a"),
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   gorm.DeletedAt{},
			Content:     "hello1",
			UserID:      "cuongnm",
			CreatedDate: "2021-10-8",
		},
		{
			ID:          uuid.MustParse("c2e4075a-0361-45c6-9191-bbd04860f58a"),
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   gorm.DeletedAt{},
			Content:     "hello1",
			UserID:      "cuongnm",
			CreatedDate: "2021-10-8",
		},
		{
			ID:          uuid.MustParse("c3e4075a-0361-45c6-9191-bbd04860f58a"),
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   gorm.DeletedAt{},
			Content:     "hello5",
			UserID:      "cuongnm",
			CreatedDate: "2021-10-8",
		},
	}

	newTask := &storages.Task{
		ID:          uuid.MustParse("c8e4075a-0361-45c6-9191-bbd04860f58a"),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
		Content:     "hello6",
		UserID:      "cuongnm",
		CreatedDate: "2021-10-8",
	}

	rows := mock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "content", "user_id", "created_date"}).
		AddRow(tasks[0].ID, tasks[0].CreatedAt, tasks[0].UpdatedAt, tasks[0].DeletedAt, tasks[0].Content, tasks[0].UserID, tasks[0].CreatedDate).
		AddRow(tasks[1].ID, tasks[1].CreatedAt, tasks[1].UpdatedAt, tasks[1].DeletedAt, tasks[1].Content, tasks[1].UserID, tasks[1].CreatedDate).
		AddRow(tasks[2].ID, tasks[2].CreatedAt, tasks[2].UpdatedAt, tasks[2].DeletedAt, tasks[2].Content, tasks[2].UserID, tasks[2].CreatedDate).
		AddRow(tasks[3].ID, tasks[3].CreatedAt, tasks[3].UpdatedAt, tasks[3].DeletedAt, tasks[3].Content, tasks[3].UserID, tasks[3].CreatedDate).
		AddRow(tasks[4].ID, tasks[4].CreatedAt, tasks[4].UpdatedAt, tasks[4].DeletedAt, tasks[4].Content, tasks[4].UserID, tasks[4].CreatedDate)
	const query = `SELECT * FROM "tasks" WHERE created_at >= $1 AND "tasks"."deleted_at" IS NULL`
	mock.ExpectQuery(query).WithArgs(now.BeginningOfDay()).WillReturnRows(rows)

	const queryUser = `SELECT * FROM "users" WHERE "users"."id" = $1`
	rowsUser := sqlmock.NewRows([]string{"id", "password", "max_todo"}).AddRow("cuongnm", "123456", 5)
	mock.ExpectQuery(queryUser).WithArgs("cuongnm").WillReturnRows(rowsUser)

	err := dbmanager.AddTask(context.Background(), newTask)
	assert.Equal(t, helper.ErrExceedMaxTaskPerDay, err)
}
