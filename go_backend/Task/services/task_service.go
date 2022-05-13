package services

import (
	"backend_test/task/models"
	"backend_test/utils"
	"database/sql"
	"errors"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

/* private functions */
func checkerLimit(db sql.DB, userId string, taskPerDay int) bool {

	startDate := time.Now().Local()
	endDate := startDate.AddDate(0, 0, 1)

	/* Count current task assigned to user */
	var currentTaskCount int
	err := db.QueryRow("SELECT COUNT(*) FROM tasks WHERE assigned_to = $1 AND created_at >= TO_DATE($2, 'YYYY-MM-DD') AND created_at < TO_DATE($3, 'YYYY-MM-DD')", userId, startDate, endDate).Scan(&currentTaskCount)
	if err != nil {
		log.Printf("Error retrieving data to database, Reason: %v\n", err)
		return false
	}

	/* compare current task count to task per day allowed */
	return LimitValidator(currentTaskCount, taskPerDay)
}

func getUserDetails(db sql.DB, userId string) (int, error) {
	var taskPerDay int
	err := db.QueryRow("SELECT task_per_day FROM users WHERE user_id = $1 ", userId).Scan(&taskPerDay)
	if err != nil {
		log.Printf("Error retrieving user to database, Reason: %v\n", err)
	}
	return taskPerDay, err
}

func LimitValidator(currentTasksCount int, taskPerDay int) bool {
	if currentTasksCount >= taskPerDay {
		log.Printf("Numbers of limit reached, allowed limit: %d -- current tasks: %d", taskPerDay, currentTasksCount)
		return false
	}

	return true
}

/* business layer */
func ProcessTasks(tasks *models.Tasks) error {

	tasks.ID = uuid.NewV4().String()
	var timeNow time.Time = time.Now()
	tasks.CreatedAt = timeNow
	tasks.UpdatedAt = timeNow

	var db *sql.DB = utils.DBConnect()
	defer db.Close()

	/* get user details */
	taskPerDay, err := getUserDetails(*db, tasks.AssignedTo)
	if err != nil {
		log.Printf("Error saving tasks to database, Reason: %v\n", err)
		return err
	}

	/* check if task limit not yet reached */
	isAllowed := checkerLimit(*db, tasks.AssignedTo, taskPerDay)

	if isAllowed {
		/* insert task if less than limit */
		err := db.QueryRow(`INSERT INTO tasks (id, assigned_to, title, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, &tasks.ID, &tasks.AssignedTo, &tasks.Title, &tasks.Description, &tasks.CreatedAt, &tasks.UpdatedAt).Scan(&tasks.ID)
		if err != nil {
			log.Printf("Error saving task to database, Reason: %v\n", err)
		}
	} else {
		err = errors.New("Task limit has been reached for user: " + tasks.AssignedTo)
	}

	return err
}

func GetTasksDetails(transactionId string) (models.Tasks, error) {

	var db *sql.DB = utils.DBConnect()

	var task models.Tasks
	err := db.QueryRow("SELECT * FROM tasks WHERE id = $1", transactionId).Scan(&task.ID, &task.AssignedTo, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		log.Printf("Error retrieving tasks to database, Reason: %v\n", err)
	}
	db.Close()

	return task, err
}

func GetAllTasksDetails() ([]models.Tasks, error) {

	var db *sql.DB = utils.DBConnect()

	var tasksList []models.Tasks
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Printf("Error retrieving tasks to database, Reason: %v\n", err)
	}
	for rows.Next() {
		var task models.Tasks

		err := rows.Scan(&task.ID, &task.AssignedTo, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}

		tasksList = append(tasksList, task)
	}
	db.Close()

	return tasksList, err
}
