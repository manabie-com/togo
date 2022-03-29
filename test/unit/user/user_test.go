package user_test

import (
	"os"
	"testing"
	"togo/globals/database"
	"togo/migration"
	"togo/models"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestShouldFindUser(t *testing.T) {
	database.SQL.Model(models.User{}).Create(&models.User{ID: 7, TasksPerDay: 4})
	user, err := models.UserById(7)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, uint(7), user.ID)
	assert.Equal(t, uint8(4), user.TasksPerDay)
}

func TestShouldCreateUserWithTask(t *testing.T) {
	var userId = 10
	userCreated, err := models.CreateUserWithTask(uint(userId), "test user with id 10")
	if err != nil {
		t.Fatal(err)
	}

	userCheck := models.User{}
	err = database.SQL.Take(&userCheck, userId).Error
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, userCreated.ID, userCheck.ID)
}

func TestShouldReachDailyLimit(t *testing.T) {
	database.SQL.Model(models.User{}).Create(&models.User{ID: 3, TasksPerDay: 3})
	for i:= 0; i < 3; i++ {
		database.SQL.Model(models.Task{}).Create(&models.Task{UserID: 3, Detail: "testing daily limit"})
	}
	res, _ := models.UserAtDailyLimit(uint(3))
	assert.Equal(t, true, res)
}

func TestShouldNotReachDailyLimit(t *testing.T) {
	database.SQL.Model(models.User{}).Create(&models.User{ID: 4, TasksPerDay: 8})
	for i:= 0; i < 3; i++ {
		database.SQL.Model(models.Task{}).Create(&models.Task{UserID: 4, Detail: "testing daily limit"})
	}
	res, _ := models.UserAtDailyLimit(uint(4))
	assert.Equal(t, false, res)
}

func setup() {
	database.InitDBConnection()
	migration.Migrate(database.SQL)
}

func teardown() {
	migration.Rollback(database.SQL)
}

func TestMain(m *testing.M){
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	setup()
	test := m.Run()
	teardown()
	os.Exit(test)
}