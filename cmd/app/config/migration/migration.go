package migration

import (
	"github.com/xrexonx/togo/cmd/app/config/database"
	"github.com/xrexonx/togo/internal/todo"
	"github.com/xrexonx/togo/internal/user"
	"log"
)

const _dbErrCreateTable = "Could not create tables to database"

// Run Create database and tables then seed sample users
func Run() {
	errCreate := database.Instance.AutoMigrate(&user.User{}, &todo.Todo{})
	if errCreate != nil {
		log.Fatal(_dbErrCreateTable, errCreate)
	}

	// Seeds sample users
	result := database.Instance.Find(&user.User{})
	if result.RowsAffected == 0 {
		u1 := user.User{Name: "Rex", MaxDailyLimit: 5, Email: "rex@gmail.com.ph"}
		u2 := user.User{Name: "Riz", MaxDailyLimit: 4, Email: "roux@gmail.com.ph"}
		u3 := user.User{Name: "Roux", MaxDailyLimit: 3, Email: "roux@gmail.com.ph"}
		sampleUsers := []user.User{u1, u2, u3}
		for _, u := range sampleUsers {
			database.Instance.Create(&u)
			log.Println("User created: " + u.Name)
		}
	}
}
