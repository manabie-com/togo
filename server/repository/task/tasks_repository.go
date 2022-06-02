package repository

import (
	"time"
	"togo/models"
)

// Define `Task` Repository Interface with the following
// Methods which will be utilized by the `UserService`
type TaskRepository interface {
	// Add a new task in the database
	CreateTask(task *models.Task) (*models.Task, error)

	// Get the number of `Tasks` for current date
	CountTask(token string, now time.Time) (int, error)
}
