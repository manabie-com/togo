package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/qgdomingo/todo-app/model"
)

// Struct where the database connection memory address is stored
type TaskRepository struct {
	DBPoolConn *pgxpool.Pool
}

// Function that fetches tasks and will add a filter in the query if there is an integer (for task id) or string (for tasks under a username) search parameter passed
func (t *TaskRepository) GetTasksDB (searchParam any) ([]model.Task, map[string]string) {
	var allTasks []model.Task
	var rows pgx.Rows
	var err error
	message := make(map[string]string)
	sql := "SELECT id, title, description, username, create_date FROM tasks "

	if searchParam != nil {
		switch searchParam.(type) {
			case int:
				sql += "WHERE id = $1 "
			case string:
				sql += "WHERE username = $1 "
			default:
				message["message"] = "Search task parameter entered is invalid"
				message["error"] = ""
				return nil, message
		}
		sql += "ORDER BY id asc"
		rows, err = t.DBPoolConn.Query(context.Background(), sql, searchParam)
	} else {
		sql += "ORDER BY id asc"
		rows, err = t.DBPoolConn.Query(context.Background(), sql)
	}

	if err != nil {
		message["message"] = "Unable to fetch data from tasks table"
		message["error"] = err.Error()
		return nil, message
	}

	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Username, &task.CreateDate)
		if err != nil {
			message["message"] = "Error encountered when row data is being fetched"
			message["error"] = err.Error()
			return nil, message
		}
		allTasks = append(allTasks, task)
	}

	if rows.Err() != nil {
		message["message"] = "Error encountered when accesssing query results"
		message["error"] = rows.Err().Error()
		return nil, message
	}

	return allTasks, nil
}

// Function that inserts a task into the tasks table, the limit of each user is computed on the sql query
func (t *TaskRepository) InsertTaskDB (task *model.TaskUserEnteredDetails) (bool, map[string]string) {
	message := make(map[string]string)
	sql := "INSERT INTO tasks (title, description, username) SELECT $1, $2, $3::VARCHAR FROM task_limit_config WHERE username = $3 AND task_limit > (SELECT COUNT(id) FROM tasks WHERE username = $3 AND create_date = current_date) RETURNING id"

	row, err := t.DBPoolConn.Query(context.Background(), sql, task.Title, task.Description, task.Username)
	if err != nil {
		message["message"] = "Unable to insert data into the tasks table"
		message["error"] = err.Error()
		return false, message
	}

	defer row.Close()

	return row.Next(), nil
}

// Function that updates a task in the tasks table
func (t *TaskRepository) UpdateTaskDB (task *model.TaskUserEnteredDetails, id int) (bool, map[string]string) {
	message := make(map[string]string)
	sql := "UPDATE tasks SET title = $1, description = $2 WHERE username = $3 AND id = $4 RETURNING id"

	row, err := t.DBPoolConn.Query(context.Background(), sql, task.Title, task.Description, task.Username, id)
	if err != nil {
		message["message"] = "Unable to update data into the tasks table"
		message["error"] = err.Error()
		return false, message
	}

	defer row.Close()

	return row.Next(), nil
}

// Function that deletes a task in the tasks table
func (t *TaskRepository) DeleteTaskDB (id int) (bool, map[string]string) {
	message := make(map[string]string)
	sql := "DELETE FROM tasks WHERE id = $1 RETURNING id"

	row, err := t.DBPoolConn.Query(context.Background(), sql, id)
	if err != nil {
		message["message"] = "Unable to delete data from the tasks table"
		message["error"] = err.Error()
		return false, message
	}

	defer row.Close()

	return row.Next(), nil
}