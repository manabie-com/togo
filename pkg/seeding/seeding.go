package seeding

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/pkg/database"
)

func Truncate() {
	database.DB().Exec("TRUNCATE TABLE users, tasks RESTART IDENTITY")
}

// SeedUsers seed users
func SeedUsers(number int16) {
	for i := int16(1); i <= number; i++ {
		user := model.NewUser()
		user.Name = fmt.Sprintf("User %d", i)
		user.LimitTaskPerDay = i
		err := user.Create()
		if err != nil {
			panic(err)
		}
	}
}

// SeedTasks seed user
func SeedTasks(number int16, userID sql.NullInt16) {
	now := time.Now()
	for i := int16(1); i <= number; i++ {
		task := model.NewTask()
		task.Content = fmt.Sprintf("Content %d", i)
		if userID.Valid {
			task.UserID = &userID.Int16
			task.DateAssign = &now
		}
		err := task.Create()
		if err != nil {
			panic(err)
		}
	}
}
