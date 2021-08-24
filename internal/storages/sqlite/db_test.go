package sqllite_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

var store *sqllite.LiteDB

func init() {
	// reset the database
	err := os.Remove("./test-data.db")
	if err != nil {
		log.Println("Could not remove ./test-data.db:", err)
	}

	db, err := sql.Open("sqlite3", "./test-data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	store = &sqllite.LiteDB{
		DB: db,
	}
}

func TestInitTables(t *testing.T) {
	err := store.InitTables()
	require.Nil(t, err, "init tables error")
}

// Test Users
var (
	user1 = storages.User{ID: "user1", Password: "password1", MaxTodo: 1}
	user2 = storages.User{ID: "user2", Password: "password2", MaxTodo: 2}
	user3 = storages.User{ID: "user3", Password: "password3", MaxTodo: 4}
)

func TestCreateUsers(t *testing.T) {
	tests := []struct {
		testName string
		user     *storages.User
		succeed  bool
	}{
		{"create user1", &user1, true},
		{"create user2", &user2, true},
		{"create user3", &user3, true},
		{"duplicate user", &user3, false},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			succeed := true
			err := store.AddUser(context.Background(), tt.user)
			if err != nil {
				succeed = false
			}
			require.Equal(t, tt.succeed, succeed, err)
		})
	}

}
func TestMaxTasks(t *testing.T) {
	tests := []struct {
		testName string
		userID   string
		maxTodo  int
		succeed  bool
	}{
		{"test user1", user1.ID, user1.MaxTodo, true},
		{"test user2", user2.ID, user2.MaxTodo, true},
		{"test user3", user3.ID, user3.MaxTodo, true},
		{"test not-exist user", "user4", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			succeed := true
			maxTodo, err := store.MaxTodo(context.Background(), tt.userID)
			if err != nil {
				succeed = false
			}
			require.Equal(t, tt.succeed, succeed, err)
			require.Equal(t, tt.maxTodo, maxTodo, "incorrect number of tasks")
		})
	}

}

func TestCreateTask(t *testing.T) {
	tests := []struct {
		testName string
		task     *storages.Task
		succeed  bool
	}{
		{"create todo #1 (u1, max 1)", &storages.Task{UserID: user1.ID, CreatedDate: "2021-1-1"}, true},
		{"create todo #2 (u1, max 1)", &storages.Task{UserID: user1.ID, CreatedDate: "2021-1-1"}, false},
		{"create todo #1 (u2, max 2)", &storages.Task{UserID: user2.ID, CreatedDate: "2021-1-1"}, true},
		{"create todo #2 (u2, max 2)", &storages.Task{UserID: user2.ID, CreatedDate: "2021-1-1"}, true},
		{"create todo #3 (u2, max 2)", &storages.Task{UserID: user2.ID, CreatedDate: "2021-1-1"}, false},
		{"create todo #1 (u1, max 1, day 2)", &storages.Task{UserID: user1.ID, CreatedDate: "2021-1-2"}, true},
		{"create todo #2 (u1, max 1, day 2)", &storages.Task{UserID: user1.ID, CreatedDate: "2021-1-2"}, false},
	}

	for k, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			succeed := true
			tt.task.ID = fmt.Sprintf("task-%d", k)
			tt.task.Content = fmt.Sprintf("content-%d", k)
			err := store.AddTask(context.Background(), tt.task)
			if err != nil {
				succeed = false
			}
			require.Equal(t, tt.succeed, succeed, err)
		})
	}

}
