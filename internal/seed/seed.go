package main

import (
	"fmt"
	"log"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/module/task"
	"github.com/manabie-com/togo/internal/module/user"
	"github.com/manabie-com/togo/internal/util"

	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	// Drop tables

	db.Migrator().DropTable(&user.User{})
	db.Migrator().DropTable(&task.Task{})

	// Migrate
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&task.Task{})
}

func seed(db *gorm.DB) {
	password := util.HashPassword("12345678")

	var user = &user.User{
		Email:    "test@mail.com",
		Password: password,
		MaxTodo:  5,
	}
	db.Save(user)

	fmt.Println(user)
	var task1 = &task.Task{
		Content: "content 01",
		Status:  task.StatusActive,
		UserID:  user.ID,
	}
	db.Save(task1)

	var task2 = &task.Task{
		Content: "content 02",
		Status:  task.StatusActive,
		UserID:  user.ID,
	}
	db.Save(task2)

}

func main() {

	//Load env
	if err := config.Load(); err != nil {
		log.Fatalf("Error env, %v", err)
	}

	if db, err := util.CreateConnectionDB(config.Cfg); err != nil {
	} else {
		defer func() {
			dbSQL, ok := db.DB()
			if ok != nil {
				defer dbSQL.Close()
			}
		}()

		migrate(db)
		seed(db)

		fmt.Println("Run seed done!")
	}
}
