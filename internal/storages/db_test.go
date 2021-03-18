package storages

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"testing"
)

func TestRetrieveTasks(t *testing.T) {
	db, err := getConnection("postgres", "localhost", 5432, "postgres", "postgres", "todo")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	userID := value("firstUser")
	createdDate := value("2021-03-17")
	var tasks []*Task
	tasks, err = (&LiteDB{DB: db}).RetrieveTasks(userID, createdDate)

	if tasks == nil {
		t.Error("TestRetrieveTasks() did not return tasks list")
	}
}

func TestAddTask(t *testing.T) {
	db, err := getConnection("postgres", "localhost", 5432, "postgres", "postgres", "todo")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	task := Task{
		ID:          "TestAddTask_2",
		Content:     "Second insert test",
		UserID:      "firstUser",
		CreatedDate: "2021-03-17",
	}

	err = (&LiteDB{DB: db}).AddTask(&task)
	if err != nil {
		t.Error("TestAddTask() did not return err != nil")
	}
}

func TestCheckNumTasksInDay(t *testing.T) {
	db, err := getConnection("postgres", "localhost", 5432, "postgres", "postgres", "todo")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	userID := value("firstUser")
	createdDate := value("2021-03-17")
	bol := (&LiteDB{DB: db}).CheckNumTasksInDay(userID, createdDate)

	if bol {
		t.Error("TestCheckNumTasksInDay() did not return num of tasks in day")
	}
}

func TestValidateUser(t *testing.T) {
	db, err := getConnection("postgres", "localhost", 5432, "postgres", "postgres", "todo")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	userID := value("firstUser")
	password := value("example")
	bol := (&LiteDB{DB: db}).ValidateUser(userID, password)

	if !bol {
		t.Error("TestValidateUser() did not return true")
	}
}

func getConnection(driverName string, host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	datasource := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	return sql.Open(driverName, datasource)
}

func value(v string) sql.NullString {
	return sql.NullString{
		String: v,
		Valid:  true,
	}
}
