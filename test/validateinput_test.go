package test

import (
	"log"
	"testing"
	"time"

	"github.com/huynhhuuloc129/todo/models"
)

func TestCheckUserInput(t *testing.T) {
	user1 := models.NewUser{
		Username: "",
		Password: "",
	}
	user2 := models.NewUser{
		Username: "",
		Password: "asdfasf",
	}
	user3 := models.NewUser{
		Username: "asdfasf",
		Password: "",
	}
	user4 := models.NewUser{
		Username: "asdfasfa",
		Password: "asdfsafsf",
	}
	result1 := models.CheckUserInput(user1)
	result2 := models.CheckUserInput(user2)
	result3 := models.CheckUserInput(user3)
	result4 := models.CheckUserInput(user4)

	if result1 || result2 || result3 || !result4 {
		log.Fatal("Check user input failed")
	}

}

func TestCheckTaskInput(t *testing.T) {
	task1 := models.NewTask{
		Content:  "sadfsaf",
		Status:   "pending",
		Time:     time.Now(),
		TimeDone: time.Now(),
		UserId:   1,
	}
	task2 := models.NewTask{
		Content:  "",
		Status:   "pending",
		Time:     time.Now(),
		TimeDone: time.Now(),
		UserId:   1,
	}
	task3 := models.NewTask{
		Content:  "sadfsaf",
		Status:   "pending",
		Time:     time.Now(),
		TimeDone: time.Now(),
		UserId:   1,
	}
	task4 := models.NewTask{
		Content:  "admin Task",
		Status:   "pending",
		Time:     time.Now(),
		TimeDone: time.Now(),
		UserId:   1,
	}
	result1 := models.CheckTaskInput(task1)
	result2 := models.CheckTaskInput(task2)
	result3 := models.CheckTaskInput(task3)
	result4 := models.CheckTaskInput(task4)

	if !result1 || result2 || !result3 || !result4 {
		log.Fatal("Check user input failed")
	}
}
