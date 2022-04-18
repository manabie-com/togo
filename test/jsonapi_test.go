package test

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func TestMain(m *testing.M) {
	user := "togo"
	pass := "togo"
	host := "localhost"
	port := "3001"
	dbName := "togo"
	login := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)
	var err error
	db, err = sql.Open("mysql", login)
	if err != nil {
		log.Fatalf("Failed to connect to db with error %v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping db: %v", err)
		return
	}

	// Setup test data
	db.Exec("TRUNCATE TABLE users")
	db.Exec("TRUNCATE TABLE tasks")
	db.Exec("TRUNCATE TABLE user_daily_counters")
	db.Exec("INSERT INTO users (id, daily_limit) values(1,5)")
	db.Exec("INSERT INTO users (id, daily_limit) values(2,1)")
	db.Exec("INSERT INTO user_daily_counters (user_id, daily_count, last_updated) values(2,2,'2020-01-01 00:00:00')")
	m.Run()
}
func TestCreateSuccess(t *testing.T) {
	taskName := "TestCreateSuccess"
	userID := 1

	// Test JSON API
	statusCode, createTaskRes, err := CreateTask(userID, taskName)
	if err != nil {
		t.Fatalf("Failed to call CreateTask API with error %v", err)
	}
	if statusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, statusCode)
	}
	if createTaskRes.Error.ErrorCode != SuccessCode {
		t.Fatalf("Expected error code %d, got %d", SuccessCode, createTaskRes.Error.ErrorCode)
	}
	if createTaskRes.Task == nil {
		t.Fatalf("Expected task field, got nil")
	}
	if createTaskRes.Task.UserID != int64(userID) {
		t.Fatalf("Expected user ID %d, got %d", userID, createTaskRes.Task.UserID)
	}
	if createTaskRes.Task.Name != taskName {
		t.Fatalf("Expected user ID %d, got %d", userID, createTaskRes.Task.UserID)
	}

	// Verify database entries
	query := `
		SELECT name from tasks where user_id = ? and name = ?
	`
	row := db.QueryRow(query, userID, taskName)
	if err != nil {
		t.Fatalf("Failed to query user tasks with error: '%v'", err)
	}
	dbTaskName := ""
	err = row.Scan(&dbTaskName)
	if err == sql.ErrNoRows {
		t.Fatalf("Task %s not found in database", taskName)
	}
	if err != nil {
		t.Fatalf("Failed to query tasks with error %v", err)
	}
	if dbTaskName != taskName {
		t.Fatalf("Expected db task name %s, got %s", taskName, dbTaskName)
	}

	query2 := `
		SELECT daily_count from user_daily_counters where user_id = ? 
	`
	row = db.QueryRow(query2, userID)
	if err != nil {
		t.Fatalf("Failed to query user tasks with error: '%v'", err)
	}
	dailyCount := 0
	err = row.Scan(&dailyCount)
	if err == sql.ErrNoRows {
		t.Fatalf("Daily counter for %d not found in database", userID)
	}
	if err != nil {
		t.Fatalf("Failed to query tasks with error %v", err)
	}
	if dailyCount != 1 {
		t.Fatalf("Expected daily counter to be %d, got %d", 1, dailyCount)
	}
}

func TestCreateFailLimitReached(t *testing.T) {
	// Test JSON API
	taskName := "TestCreateFailLimitReached1"
	userID := 2
	statusCode, _, err := CreateTask(userID, taskName)
	if err != nil {
		t.Fatalf("Failed to call CreateTask API with error %v", err)
	}
	if statusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, statusCode)
	}

	// Test this will fail since user 2 has a daily limit of 1
	taskName = "TestCreateFailLimitReached2"
	statusCode, createTaskRes, err := CreateTask(userID, taskName)
	if err != nil {
		t.Fatalf("Failed to call CreateTask API with error %v", err)
	}
	if statusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, statusCode)
	}
	if createTaskRes.Error.ErrorCode != MaxLimitCode {
		t.Fatalf("Expected error code %d, got %d", MaxLimitCode, createTaskRes.Error.ErrorCode)
	}

	query := `
		SELECT daily_count from user_daily_counters where user_id = ? 
	`
	row := db.QueryRow(query, userID)
	if err != nil {
		t.Fatalf("Failed to query user tasks with error: '%v'", err)
	}
	dailyCount := 0
	err = row.Scan(&dailyCount)
	if err == sql.ErrNoRows {
		t.Fatalf("Daily counter for %d not found in database", userID)
	}
	if err != nil {
		t.Fatalf("Failed to query tasks with error %v", err)
	}
	if dailyCount != 1 {
		t.Fatalf("Expected daily counter to be %d, got %d", 1, dailyCount)
	}
}

func TestCreateFailBadRequest(t *testing.T) {
	taskName := ""
	userID := 3
	// Test this will fail since empty task name is not allowed
	statusCode, _, err := CreateTask(userID, taskName)
	if err != nil {
		t.Fatalf("Failed to call CreateTask API with error %v", err)
	}
	if statusCode != http.StatusBadRequest {
		t.Fatalf("Expected status code %d, got %d", http.StatusBadRequest, statusCode)
	}
}
