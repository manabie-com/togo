package test

import (
	"log"
	"testing"

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
