package storages_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/stretchr/testify/require"
)

var stores = []storages.Store{}

func init() {
	createSQLLiteStore()
	createPostgresStore()
}

func createSQLLiteStore() {
	// Setup SQLITE
	err := os.Remove("./test-data.db")
	if err != nil {
		log.Println("Could not remove ./test-data.db:", err)
	}

	db, err := sql.Open("sqlite3", "./test-data.db")
	if err != nil {
		log.Fatal("error opening db ", err)
	}

	stores = append(stores, &sqlite.LiteDB{
		DB: db,
	})
}

func createPostgresStore() {
	// Connect to Test Postgres
	const (
		host     = "localhost"
		port     = 65432
		user     = "postgres"
		password = "postgres"
		dbname   = "postgres-test"
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("error opening db ", err)
	}

	_, err = db.Exec("drop table tasks; drop table users; ")
	if err != nil {
		log.Fatal("error clearing test-postgres", err)
	}

	stores = append(stores, &postgres.PostgresDB{
		DB: db,
	})
}
func storeName(store storages.Store) string {
	return reflect.TypeOf(store).String()
}

func TestInitTables(t *testing.T) {
	for _, store := range stores {
		fmt.Println("Testing store", storeName(store))
		err := store.InitTables()
		require.Nil(t, err, "init tables error")
	}
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

	for _, store := range stores {
		for _, tt := range tests {
			t.Run(storeName(store)+"/"+tt.testName, func(t *testing.T) {
				succeed := true
				err := store.AddUser(context.Background(), tt.user)
				if err != nil {
					succeed = false
				}
				require.Equal(t, tt.succeed, succeed, err)
			})
		}
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
	for _, store := range stores {
		for _, tt := range tests {
			t.Run(storeName(store)+"/"+tt.testName, func(t *testing.T) {
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
}

func TestCreateTask(t *testing.T) {
	tests := []struct {
		testName string
		task     *storages.Task
		canAdd   bool
		noError  bool
	}{
		{"create todo #1 (u1, max 1)", &storages.Task{UserID: user1.ID, CreatedDate: "2021-1-1"}, true, true},
		{"create todo #2 (u1, max 1)", &storages.Task{UserID: user1.ID, CreatedDate: "2021-1-1"}, false, true},
		{"create todo #1 (u2, max 2)", &storages.Task{UserID: user2.ID, CreatedDate: "2021-1-1"}, true, true},
		{"create todo #2 (u2, max 2)", &storages.Task{UserID: user2.ID, CreatedDate: "2021-1-1"}, true, true},
		{"create todo #3 (u2, max 2)", &storages.Task{UserID: user2.ID, CreatedDate: "2021-1-1"}, false, true},
		{"create todo #1 (u1, max 1, day 2)", &storages.Task{UserID: user1.ID, CreatedDate: "2021-1-2"}, true, true},
		{"create todo #2 (u1, max 1, day 2)", &storages.Task{UserID: user1.ID, CreatedDate: "2021-1-2"}, false, true},
	}

	for _, store := range stores {
		for k, tt := range tests {
			t.Run(storeName(store)+"/"+tt.testName, func(t *testing.T) {
				noError := true
				tt.task.ID = fmt.Sprintf("task-%d", k)
				tt.task.Content = fmt.Sprintf("content-%d", k)
				canAdd, err := store.AddTask(context.Background(), tt.task)
				if err != nil {
					noError = false
				}
				require.Equal(t, tt.canAdd, canAdd, "canAdd not equal")
				require.Equal(t, tt.noError, noError)
			})
		}
	}

}

func TestValidateUser(t *testing.T) {
	tests := []struct {
		testName string
		id       string
		password string
		succeed  bool
	}{
		{"validate user1", user1.ID, user1.Password, true},
		{"validate user2", user2.ID, user2.Password, true},
		{"validate user3", user3.ID, user3.Password, true},
		{"validate user1 with wrong password", user1.ID, user1.Password + "asdf", false},
		{"validate user2 with wrong accounts password", user2.ID, user1.Password, false},
		{"validate user3 with no password", user3.ID, "", false},
	}
	for _, store := range stores {
		for _, tt := range tests {
			t.Run(storeName(store)+"/"+tt.testName, func(t *testing.T) {
				succeed := store.ValidateUser(context.Background(), tt.id, tt.password)
				require.Equal(t, tt.succeed, succeed)
			})
		}
	}
}

func TestUpdatePassword(t *testing.T) {
	for _, store := range stores {
		fmt.Println("Testing store", storeName(store))

		err := store.SetUserPassword(context.Background(), user1.ID, user1.Password+"-updated")
		require.Nil(t, err)

		ok := store.ValidateUser(context.Background(), user1.ID, user1.Password)
		require.False(t, ok)

		ok = store.ValidateUser(context.Background(), user1.ID, user1.Password+"-updated")
		require.True(t, ok)
		user1.Password = user1.Password + "-updated"
	}
}
