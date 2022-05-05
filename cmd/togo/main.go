package main

import (
	"fmt"
	"github.com/jmsemira/togo/internal/api"
	"github.com/jmsemira/togo/internal/auth"
	"github.com/jmsemira/togo/internal/database"
	"github.com/jmsemira/togo/internal/models"
)

func main() {
	api.InitApi()
}

func init() {
	dbSettings := database.DBSettings{
		Type: "sqlite",
		Name: "togo.db",
	}

	// initialize and migrate schema
	database.InitializeDB(dbSettings, &models.User{}, &models.Todo{})

	// check if the system already had a user
	users := []models.User{}

	db := database.GetDB()
	db.Find(&users)

	if len(users) == 0 {
		// initialize system users
		for _, i := range []int{1, 2, 3} {
			user := models.User{}
			user.Username = fmt.Sprintf("user%v", i)
			user.Password = auth.HashPass(user.Username)

			user.RateLimitPerDay = i
			db.Create(&user)
		}
	}
}
