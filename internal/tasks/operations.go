package tasks

import (
	"log"
	"time"

	"github.com/kozloz/togo"
	"github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/users"
)

type TaskStore interface {
	CreateTask(userID int64, task string) (*togo.Task, error)
}
type Operation struct {
	store         TaskStore
	userOperation *users.Operation
}

func NewOperation(store TaskStore, userOp *users.Operation) *Operation {
	return &Operation{
		store:         store,
		userOperation: userOp,
	}
}

// Create the task for the user
func (o *Operation) Create(userID int64, taskName string) (*togo.Task, error) {
	log.Printf("Creating task '%s' for user '%d'.", taskName, userID)
	// Get user object
	user, err := o.userOperation.Get(userID)
	if err != nil {
		return nil, err
	}

	// Create user if doesnt exist yet
	if user == nil {
		user, err = o.userOperation.Create(userID)
		if err != nil {
			return nil, err
		}
	}
	// Check if can still create tasks. Consider 0 as unlimited.
	counterYear, counterDay := 0, 0
	var counterMonth time.Month
	todayYear, todayMonth, todayDay := time.Now().Date()
	if user.DailyCounter != nil && user.DailyLimit > 0 {
		counterYear, counterMonth, counterDay = user.DailyCounter.LastUpdated.Date()
		if counterYear == todayYear && counterMonth == todayMonth && counterDay == todayDay &&
			user.DailyCounter.DailyCount >= user.DailyLimit {
			log.Printf("Max daily limit reached for today")
			return nil, errors.MaxLimit
		}
	}

	// Insert task
	task, err := o.store.CreateTask(user.ID, taskName)
	if err != nil {
		return nil, err
	}

	// Update user attributes
	if user.DailyLimit > 0 {
		if user.DailyCounter == nil {
			user.DailyCounter = &togo.DailyCounter{
				UserID: user.ID,
			}
		}

		// Reset the daily counter if last recorded was a different day
		if !(counterYear == todayYear && counterMonth == todayMonth && counterDay == todayDay) {
			user.DailyCounter.DailyCount = 0
			user.DailyCounter.LastUpdated = time.Now()
		}

		user.DailyCounter.DailyCount++
		log.Printf("Current counter: %v", user.DailyCounter)
	}
	// Create counter only if a limit exists for the user

	user.Tasks = append(user.Tasks, task)
	_, err = o.userOperation.Update(user)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully created task '%v'", task)

	return task, nil
}
