package todo

import (
	"fmt"
	db "github.com/xrexonx/togo/cmd/app/config/database"
	"github.com/xrexonx/togo/internal/repository"
	userService "github.com/xrexonx/togo/internal/user"
	"log"
	"time"
)

// getBeginningOfDay
func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// Add service to add new todo
func Add(todo Todo) (Todo, error) {

	var pendingTask int64

	// Validate userID
	if len(todo.UserId) == 0 {
		return todo, fmt.Errorf("userID is required")
	}

	// Get user daily limit
	user := userService.FindByID(todo.UserId)

	// Get count of pending todo by user for the current day
	today := getBeginningOfDay(time.Now())
	condition := "completed = false AND user_id = ? AND created_at >= ?"
	results := db.Instance.Model(&Todo{}).
		Where(condition, todo.UserId, today).
		Count(&pendingTask).Error

	log.Println("tasks results:", results)

	// Validate
	if pendingTask >= int64(user.MaxDailyLimit) {
		return Todo{}, fmt.Errorf("user have reached the maximum daily limit")
	}

	// Create new todo
	var newTodo, _ = repository.Create[Todo](todo)
	return newTodo, nil
}
