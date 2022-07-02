package todo

import (
	"fmt"
	"github.com/xrexonx/togo/internal/repository"
	userService "github.com/xrexonx/togo/internal/user"
	"gorm.io/gorm"
	"log"
	"time"
)

func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func Add(db *gorm.DB, todo Todo) (Todo, error) {

	var pendingTask int64

	// Get user daily limit
	user := userService.GetById(db, todo.UserId)

	// Validate
	today := getBeginningOfDay(time.Now())
	condition := "completed = false AND user_id = ? AND created_at >= ?"
	results := db.Model(&Todo{}).
		Where(condition, todo.UserId, today).
		Count(&pendingTask).Error

	log.Println("tasks results:", results)

	if pendingTask >= int64(user.MaxDailyLimit) {
		return Todo{}, fmt.Errorf("user have reached the maximum daily limit")
	}

	var newTodo, _ = repository.Create[Todo](db, todo)
	return newTodo, nil
}
