package test

import (
	"log"
	"testing"
	"time"

	"github.com/huynhhuuloc129/todo/models"
)

func TestCheckTaskInput(t *testing.T){
	task1 := models.NewTask{
		Content: "sadfsaf",
		Status: "pending",
		Time: time.Now(),
		TimeDone: time.Now(),
		UserId: 1,
	}
	task2 := models.NewTask{
		Content: "",
		Status: "pending",
		Time: time.Now(),
		TimeDone: time.Now(),
		UserId: 1,
	}
	task3 := models.NewTask{
		Content: "sadfsaf",
		Status: "pending",
		Time: time.Now(),
		TimeDone: time.Now(),
		UserId: 1,
	}
	task4 := models.NewTask{
		Content: "admin Task",
		Status: "pending",
		Time: time.Now(),
		TimeDone: time.Now(),
		UserId: 1,
	}
	result1 := models.CheckTaskInput(task1)
	result2 := models.CheckTaskInput(task2)
	result3 := models.CheckTaskInput(task3)
	result4 := models.CheckTaskInput(task4)

	if !result1 || result2 || !result3 || !result4 {
		log.Fatal("Check user input failed")
	}
}