package repository

import (
	"testing"
	"github.com/qgdomingo/todo-app/database"
	"github.com/qgdomingo/todo-app/model"
)

func TestFetchAllTasks(t *testing.T) {
	dbConn, err := database.CreateConnection()

	if err != nil {
		t.Fatalf("Unable to connect to the database: %v", err.Error())
	}

	allTasks, errMessage := GetTasksDB(dbConn, nil)

	if errMessage != nil {
		t.Errorf("Failed to pull tasks from database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	if len(allTasks) == 0 {
		t.Log("There are no tasks available from database")
	}

	dbConn.Close()
}

func TestFetchTaskByID(t *testing.T) {
	dbConn, err := database.CreateConnection()

	if err != nil {
		t.Fatalf("Unable to connect to the database: %v", err.Error())
	}

	// Fetch all tasks and use first one as sample
	allTasks, errMessage := GetTasksDB(dbConn, nil)

	if errMessage != nil {
		t.Errorf("Failed to pull tasks from database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	if len(allTasks) == 0 {
		t.Error("There are no tasks available from database")
	} else {
		sampleID := allTasks[0].ID

		// Fetch the task using the sample ID
		taskFetched, errMessage := GetTasksDB(dbConn, sampleID)

		if errMessage != nil {
			t.Errorf("Failed to pull task by ID from database, message: %v | error: %v", errMessage["message"], errMessage["error"])
		}

		if len(taskFetched) == 0 {
			t.Errorf("Task with ID %v from database is not found", sampleID)
		} 

		if sampleID != taskFetched[0].ID {
			t.Errorf("Expected task ID %v is different from the fetched task ID %v", sampleID, taskFetched[0].ID)
		}
	}

	dbConn.Close()
}

func TestFetchTaskByUsername(t *testing.T) {
	dbConn, err := database.CreateConnection()

	if err != nil {
		t.Fatalf("Unable to connect to the database: %v", err.Error())
	}

	// Fetch all tasks 
	allTasks, errMessage := GetTasksDB(dbConn, nil)

	if errMessage != nil {
		t.Errorf("Failed to pull tasks from database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	if len(allTasks) == 0 {
		t.Error("There are no tasks available from database")
	} else {

		// Get the count of tasks owned using the first task user name
		expectedTaskCount := 0
		sampleUsername := allTasks[0].Username

		for _, task := range allTasks {
			if task.Username == sampleUsername {
				expectedTaskCount += 1
			}
		}

		// Fetch the all tasks under the sample username
		taskFetched, errMessage := GetTasksDB(dbConn, sampleUsername)

		if errMessage != nil {
			t.Errorf("Failed to pull task by username from database, message: %v | error: %v", errMessage["message"], errMessage["error"])
		}

		fetchedTaskCount := len(taskFetched)

		if fetchedTaskCount == 0 {
			t.Errorf("Tasks owned by %v are not found in the database", sampleUsername)
		} 

		if fetchedTaskCount != expectedTaskCount {
			t.Errorf("Expected task count %v is different from the fetched task count %v", expectedTaskCount, fetchedTaskCount)
		}
	}

	dbConn.Close()
}

func TestFetchTaskEmptyStringParameter(t *testing.T) {
	dbConn, err := database.CreateConnection()

	if err != nil {
		t.Fatalf("Unable to connect to the database: %v", err.Error())
	}

	// Fetch task using empty string
	taskFetched, errMessage := GetTasksDB(dbConn, "")

	if errMessage != nil {
		t.Errorf("Failed to pull task by ID from database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	fetchedTaskCount := len(taskFetched)

	if fetchedTaskCount > 0 {
		t.Errorf("There were tasks fetched using an empty username string")
	} 

	dbConn.Close()
}

func TestCreateTask(t *testing.T) {
	dbConn, err := database.CreateConnection()

	if err != nil {
		t.Fatalf("Unable to connect to the database: %v", err.Error())
	}

	taskData := model.TaskUserEnteredDetails{
		Title: "Unit Test Create",
		Description: "Unit Test Create",
		Username: "todo_test_user",
	}

	isTaskCreated, errMessage := InsertTaskDB(dbConn, &taskData)

	if errMessage != nil {
		t.Errorf("Failed to insert data into the database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	if !isTaskCreated {
		t.Error("Failed to insert data into the database")
	}

	dbConn.Close()
}

func TestCreateTaskInvalidUser(t *testing.T) {
	dbConn, err := database.CreateConnection()

	if err != nil {
		t.Fatalf("Unable to connect to the database: %v", err.Error())
	}

	taskData := model.TaskUserEnteredDetails{
		Title: "Unit Test Create",
		Description: "Unit Test Create",
		Username: "todo_test_user_fake",
	}

	isTaskCreated, errMessage := InsertTaskDB(dbConn, &taskData)

	if errMessage != nil {
		t.Errorf("Failed to attempt to insert data into the database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	if isTaskCreated {
		t.Error("Failed validation, data was inserted")
	}

	dbConn.Close()
}

func TestUpdateTask(t *testing.T) {
	dbConn, err := database.CreateConnection()

	if err != nil {
		t.Fatalf("Unable to connect to the database: %v", err.Error())
	}

	// Fetch sample tasks created by todo_test_user
	taskList, errMessage := GetTasksDB(dbConn, "todo_test_user")

	if errMessage != nil {
		t.Errorf("Failed to pull task by username from database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	if len(taskList) == 0 {
		t.Error("There are no tasks available by test user from database")
	}

	taskData := model.TaskUserEnteredDetails{
		Title: "Unit Test Update",
		Description: "Unit Test Update",
		Username: "todo_test_user",
	}

	isTaskUpdated, errMessage := UpdateTaskDB(dbConn, &taskData, taskList[0].ID)

	if errMessage != nil {
		t.Errorf("Failed to update data into the database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	if !isTaskUpdated {
		t.Errorf("Failed to update data %v into the database", taskList[0].ID)
	}

	dbConn.Close()
}

func TestUpdateTaskInvalidUser(t *testing.T) {
	dbConn, err := database.CreateConnection()

	if err != nil {
		t.Fatalf("Unable to connect to the database: %v", err.Error())
	}

	// Fetch sample tasks created by todo_test_user
	taskList, errMessage := GetTasksDB(dbConn, "todo_test_user")

	if errMessage != nil {
		t.Errorf("Failed to pull task by username from database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	if len(taskList) == 0 {
		t.Error("There are no tasks available by test user from database")
	}

	taskData := model.TaskUserEnteredDetails{
		Title: "Unit Test Update",
		Description: "Unit Test Update",
		Username: "todo_test_user_fake",
	}

	isTaskUpdated, errMessage := UpdateTaskDB(dbConn, &taskData, taskList[0].ID)

	if errMessage != nil {
		t.Errorf("Failed to update data into the database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	if isTaskUpdated {
		t.Errorf("Failed validation, data with ID %v was updated", taskList[0].ID)
	}

	dbConn.Close()
}

func TestDeleteTask(t *testing.T) {
	dbConn, err := database.CreateConnection()

	if err != nil {
		t.Fatalf("Unable to connect to the database: %v", err.Error())
	}

	// Fetch sample tasks created by test user
	taskList, errMessage := GetTasksDB(dbConn, "todo_test_user")

	if errMessage != nil {
		t.Errorf("Failed to pull task by username from database, message: %v | error: %v", errMessage["message"], errMessage["error"])
	}

	if len(taskList) == 0 {
		t.Error("There are no tasks available by test user from database")
	}

	for _, task := range taskList {
		isTaskDeleted, errMessage := DeleteTaskDB(dbConn, task.ID)

		if errMessage != nil {
			t.Errorf("Failed to update data into the database, message: %v | error: %v", errMessage["message"], errMessage["error"])
		}
	
		if !isTaskDeleted {
			t.Errorf("Failed to delete task %v into the database", task.ID)
		}
	}

	dbConn.Close()
}