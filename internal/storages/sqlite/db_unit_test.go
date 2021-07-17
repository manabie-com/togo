package sqllite

import (
	"context"
	"database/sql"
	"github.com/bxcodec/faker/v3"
	"github.com/manabie-com/togo/internal/helpers"
	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func beforeTest() LiteDB {
	db, err := sql.Open("sqlite3", "../../../__fixtures__/data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	return LiteDB{DB: db}
}

func afterTest(l LiteDB) {
	l.DB.Exec(`DELETE FROM tasks;`)
	l.DB.Exec(`DELETE FROM users;`)
}

func TestLiteDB_ValidateUserSuccessBecauseUserExists(t *testing.T) {
	assert := assert.New(t)
	lite := beforeTest()

	email := faker.Email()
	userPassword := faker.Password()
	password, _ := helpers.HashPassword(userPassword)
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, email, password, 5)

	result := lite.ValidateUser(
		context.Background(),
		sql.NullString{String: email, Valid: true},
		sql.NullString{String: userPassword, Valid: true},
	)

	assert.True(result)
	afterTest(lite)
}

func TestLiteDB_ValidateUserFailedBecauseUserExistsButWrongPassword(t *testing.T) {
	assert := assert.New(t)
	lite := beforeTest()

	password, _ := helpers.HashPassword(faker.Password())
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, "user", password, 5)

	result := lite.ValidateUser(
		context.Background(),
		sql.NullString{String: "user", Valid: true},
		sql.NullString{String: faker.Password(), Valid: true},
	)

	assert.False(result)
	afterTest(lite)
}

func TestLiteDB_AddTaskSuccess(t *testing.T) {
	assert := assert.New(t)
	lite := beforeTest()

	password, _ := helpers.HashPassword(faker.Password())
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, "user", password, 5)
	result := lite.AddTask(
		context.Background(),
		&storages.Task{
			ID:          faker.UUIDHyphenated(),
			Content:     faker.Name(),
			UserID:      "user",
			CreatedDate: faker.Date(),
		},
	)

	assert.NoError(result)
	assert.Nil(result)

	afterTest(lite)
}

func TestLiteDB_RetrieveTasksEmptyBecauseNoTasks(t *testing.T) {
	assert := assert.New(t)
	lite := beforeTest()

	result, err := lite.RetrieveTasks(
		context.Background(),
		sql.NullString{String: faker.UUIDHyphenated(), Valid: true},
		sql.NullString{String: "", Valid: false},
	)

	assert.NoError(err)
	assert.Nil(err)
	assert.Nil(result)

	afterTest(lite)
}

func TestLiteDB_RetrieveTasksHaveTasksShowAllTaskOfUser(t *testing.T) {
	assert := assert.New(t)
	lite := beforeTest()
	userId := faker.Name()

	password, _ := helpers.HashPassword(faker.Password())
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, userId, password, 5)
	lite.AddTask(
		context.Background(),
		&storages.Task{
			ID:          faker.UUIDHyphenated(),
			Content:     faker.Name(),
			UserID:      userId,
			CreatedDate: faker.Date(),
		},
	)
	result, err := lite.RetrieveTasks(
		context.Background(),
		sql.NullString{String: userId, Valid: true},
		sql.NullString{String: "", Valid: false},
	)

	assert.NoError(err)
	assert.Nil(err)

	assert.NotNil(result)

	afterTest(lite)
}

func TestLiteDB_RetrieveTasksHaveTasksFilteredByCreatedDateOfUser(t *testing.T) {
	assert := assert.New(t)
	lite := beforeTest()
	userId := faker.Name()

	password, _ := helpers.HashPassword(faker.Password())
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, userId, password, 5)
	lite.AddTask(
		context.Background(),
		&storages.Task{
			ID:          faker.UUIDHyphenated(),
			Content:     faker.Name(),
			UserID:      userId,
			CreatedDate: "2021-07-15",
		},
	)
	lite.AddTask(
		context.Background(),
		&storages.Task{
			ID:          faker.UUIDHyphenated(),
			Content:     faker.Name(),
			UserID:      userId,
			CreatedDate: "2021-07-14",
		},
	)

	result, err := lite.RetrieveTasks(
		context.Background(),
		sql.NullString{String: userId, Valid: true},
		sql.NullString{String: "2021-07-15", Valid: true},
	)

	assert.NoError(err)
	assert.Nil(err)

	assert.NotNil(result)
	assert.Equal(1, len(result))
	assert.Equal("2021-07-15", result[0].CreatedDate)

	afterTest(lite)
}

func TestLiteDB_GetUserMaxToDo(t *testing.T) {
	assert := assert.New(t)
	lite := beforeTest()

	email := faker.Email()
	userPassword := faker.Password()
	password, _ := helpers.HashPassword(userPassword)
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, email, password, 5)

	result, _ := lite.GetUserMaxToDo(
		context.Background(),
		email,
	)

	assert.Equal(5, result)
	afterTest(lite)
}

func TestLiteDB_GetTotalAddedTasksOfUserByDate(t *testing.T) {
	assert := assert.New(t)
	lite := beforeTest()

	email := faker.Email()
	userPassword := faker.Password()
	password, _ := helpers.HashPassword(userPassword)
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, email, password, 5)

	result, _ := lite.GetTotalAddedTasksOfUserByDate(
		context.Background(),
		email,
		"2020-05-05",
	)

	assert.Equal(0, result)

	// add tasks
	lite.AddTask(
		context.Background(),
		&storages.Task{
			ID:          faker.UUIDHyphenated(),
			Content:     faker.Name(),
			UserID:      email,
			CreatedDate: "2020-05-05",
		},
	)

	result, err := lite.GetTotalAddedTasksOfUserByDate(
		context.Background(),
		email,
		"2020-05-05",
	)

	assert.Equal(1, result)
	assert.NoError(err)

	afterTest(lite)
}

func TestLiteDB_AddTaskFailedBecauseUserReachedMaxiumum(t *testing.T) {
	assert := assert.New(t)
	lite := beforeTest()

	email := faker.Email()
	userPassword := faker.Password()
	password, _ := helpers.HashPassword(userPassword)
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, email, password, 1)

	// add tasks
	result := lite.AddTask(
		context.Background(),
		&storages.Task{
			ID:          faker.UUIDHyphenated(),
			Content:     faker.Name(),
			UserID:      email,
			CreatedDate: "2020-05-05",
		},
	)

	assert.NoError(result)

	// add tasks
	result = lite.AddTask(
		context.Background(),
		&storages.Task{
			ID:          faker.UUIDHyphenated(),
			Content:     faker.Name(),
			UserID:      email,
			CreatedDate: "2020-05-05",
		},
	)

	assert.Error(result)
	assert.Contains(result.Error(), "reached the maximum tasks")

	afterTest(lite)
}
