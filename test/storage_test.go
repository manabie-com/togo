package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/driver"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
)

func Test_Storages_GetTasks(t *testing.T) {

	var (
		uname = "notail"
		pwd   = "1234567"
	)

	db, err := driver.GetConnection()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	repository := postgres.InitRepository(db.Conn)

	arrTasks, err := repository.GetTasks(uname, pwd)

	fmt.Println("arrTasks:", arrTasks)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func Test_Storages_InsertTask(t *testing.T) {

	var (
		err   error
		now   = time.Now()
		task  = storages.Task{}
		uname = "notail"
	)

	task.ID = uuid.New().String()
	task.UserID = uname
	task.Content = "task 3"
	task.CreatedDate = now.Format("2006-01-02")

	db, err := driver.GetConnection()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	repository := postgres.InitRepository(db.Conn)

	err = repository.InsertTask(&task)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func Test_Storages_Login(t *testing.T) {

	var (
		uname = "notail"
		pwd   = "1234567"
	)

	db, err := driver.GetConnection()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	repository := postgres.InitRepository(db.Conn)

	err = repository.Login(uname, pwd)

	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_Storages_GetUserByID(t *testing.T) {

	var (
		uname = "notail"
	)

	db, err := driver.GetConnection()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	repository := postgres.InitRepository(db.Conn)

	_, err = repository.GetUserByID(uname)

	if err != nil {
		t.Errorf(err.Error())
	}
}
