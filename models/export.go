package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huynhhuuloc129/todo/util"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

// create random user
func RandomUser() User {
	user := User{
		Id:        util.RandomId(),
		Username:  util.RandomUsername(),
		Password:  util.RandomPassword(),
		LimitTask: int(util.RandomLimittask()),
	}
	return user
}

// create random new user
func RandomNewUser() NewUser {
	newUser := NewUser{
		Username:  util.RandomUsername(),
		Password:  util.RandomPassword(),
		LimitTask: int(util.RandomLimittask()),
	}
	return newUser
}

// fucntion create a random task
func RandomTask() Task {
	task := Task{
		Id:      int(util.RandomId()),
		Content: util.RandomContent(),
		Status:  "pending",
		Time:    time.Now(),
		UserId:  int(util.RandomUserid()),
	}
	return task
}

//function create a random new task
func RandomNewTask() NewTask {
	task := NewTask{
		Content: util.RandomContent(),
		Status:  "pending",
		Time:    time.Now(),
		UserId:  int(util.RandomUserid()),
	}
	return task
}
