package models

import (
	"log"
	"testing"
	"time"
)

func TestCheckUserInput(t *testing.T) {
	user1 := NewUser{
		Username: "",
		Password: "",
	}
	user2 := NewUser{
		Username: "",
		Password: "asdfasf",
	}
	user3 := NewUser{
		Username: "asdfasf",
		Password: "",
	}
	user4 := NewUser{
		Username: "asdfasfa",
		Password: "asdfsafsf",
	}
	result1 := CheckUserInput(user1)
	result2 := CheckUserInput(user2)
	result3 := CheckUserInput(user3)
	result4 := CheckUserInput(user4)

	if result1 || result2 || result3 || !result4 {
		log.Fatal("Check user input failed")
	}

}

func TestCheckTaskInput(t *testing.T) {
	task1 := NewTask{
		Content:  "sadfsaf",
		Status:   "pending",
		Time:     time.Now(),
		TimeDone: time.Now(),
		UserId:   1,
	}
	task2 := NewTask{
		Content:  "",
		Status:   "pending",
		Time:     time.Now(),
		TimeDone: time.Now(),
		UserId:   1,
	}
	task3 := NewTask{
		Content:  "sadfsaf",
		Status:   "pending",
		Time:     time.Now(),
		TimeDone: time.Now(),
		UserId:   1,
	}
	task4 := NewTask{
		Content:  "admin Task",
		Status:   "pending",
		Time:     time.Now(),
		TimeDone: time.Now(),
		UserId:   1,
	}
	result1 := CheckTaskInput(task1)
	result2 := CheckTaskInput(task2)
	result3 := CheckTaskInput(task3)
	result4 := CheckTaskInput(task4)

	if !result1 || result2 || !result3 || !result4 {
		t.Fatal("Check user input failed")
	}
}
