package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	testUserModel *UserModel
	testTaskModel *TaskModel

	/* Testing User variables */
	testUserVar []User

	/* Testing Task variables */
	testTaskVar []Task
)

func TestMain(m *testing.M) {
	/* Testing models */
	var db *gorm.DB
	/* First load config from .env.test */
	/* Change workspace dir to read .env */
	_, filename, _, _ := runtime.Caller(0)
	os.Chdir(path.Join(path.Dir(filename), ".."))

	wd, _ := os.Getwd()
	err := godotenv.Load(filepath.Join(wd, ".env-test"))
	if err != nil {
		log.Fatal("Cannot load .env-test file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	if db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal("[MODEL TESTING] Cannot connect to database")
	}

	/* Database migration */
	// Drop all tables before testing
	db.Migrator().DropTable(&Task{}, &User{})

	// Migrating database
	if err = db.AutoMigrate(
		&User{},
		&Task{},
	); err != nil {
		log.Fatal("[MODEL TESTING] Failed when migrate database")
	}

	/* Init models */
	testUserModel = NewUserModel(db)
	testTaskModel = NewTaskModel(db)

	/* Init testing variables */
	testUserVar = append(testUserVar,
		User{UserID: "User_ID#1", DailyTasksLimit: 8, MaxDailyTasks: 8},   // Test create and get
		User{UserID: "User_ID#2", DailyTasksLimit: 10, MaxDailyTasks: 10}, // Test update
		User{UserID: "User_ID#3", DailyTasksLimit: 16, MaxDailyTasks: 16}, // Test non exist userID
	)
	testTaskVar = append(testTaskVar,
		Task{UserID: testUserVar[0].UserID, TaskDetail: "Todo task content 1"}, // With UserId exist
		Task{UserID: testUserVar[2].UserID, TaskDetail: "Todo task content 2"}, // With UserId not exist
	)
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	testContext := context.Background()
	newTestUser, err := testUserModel.CreateUser(testContext, testUserVar[0].UserID, testUserVar[0].MaxDailyTasks)
	require.NoError(t, err)
	require.Equal(t, testUserVar[0].UserID, newTestUser.UserID)
	require.Equal(t, testUserVar[0].DailyTasksLimit, newTestUser.DailyTasksLimit)
	require.Equal(t, testUserVar[0].MaxDailyTasks, newTestUser.MaxDailyTasks)

	// Test column UserID constraint
	newTestUser, err = testUserModel.CreateUser(testContext, testUserVar[0].UserID, testUserVar[0].MaxDailyTasks)
	require.Error(t, err)
}

func TestUpdateUser(t *testing.T) {
	testContext := context.Background()
	newTestUser, err := testUserModel.CreateUser(testContext, testUserVar[1].UserID, testUserVar[1].MaxDailyTasks)
	require.NoError(t, err)
	require.Equal(t, testUserVar[1].UserID, newTestUser.UserID)
	require.Equal(t, testUserVar[1].DailyTasksLimit, newTestUser.DailyTasksLimit)
	require.Equal(t, testUserVar[1].MaxDailyTasks, newTestUser.MaxDailyTasks)

	newTestUser.DailyTasksLimit = 5
	isUpdated, err := testUserModel.UpdateUser(testContext, newTestUser)
	require.NoError(t, err)
	require.Equal(t, isUpdated, true)
}

func TestGetUserByUserId(t *testing.T) {
	/* Perform 2 test case:
	 * 1. Test get non exist user and it must return nil
	 * 2. Test get exist user by userId
	 */
	// Case 1
	testContext := context.Background()
	testUser, err := testUserModel.GetUserByUserId(testContext, testUserVar[2].UserID)
	require.NoError(t, err)
	require.Nil(t, testUser)

	// Case 2
	testUser, err = testUserModel.GetUserByUserId(testContext, testUserVar[0].UserID)
	require.NoError(t, err)
	require.Equal(t, testUserVar[0].UserID, testUser.UserID)
	require.Equal(t, testUserVar[0].DailyTasksLimit, testUser.DailyTasksLimit)
	require.Equal(t, testUserVar[0].MaxDailyTasks, testUser.MaxDailyTasks)
}

func TestCreateTask(t *testing.T) {
	testContext := context.Background()
	// Case 1: With USERID exist
	newTask, err := testTaskModel.CreateTask(testContext, testTaskVar[0].UserID, testTaskVar[0].TaskDetail)
	require.NoError(t, err)
	require.Equal(t, testTaskVar[0].TaskDetail, newTask.TaskDetail)
	require.Equal(t, testTaskVar[0].UserID, newTask.UserID)

	// Case 2: With USERID not exist
	newTask, err = testTaskModel.CreateTask(testContext, testTaskVar[1].UserID, testTaskVar[1].TaskDetail)
	require.Error(t, err)
	require.Nil(t, newTask)
}
