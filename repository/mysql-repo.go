package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/manabie-com/backend/entity"
	"github.com/manabie-com/backend/utils"
)

const (
	queryGetTaskAll             = ("SELECT * FROM task;")
	queryGetTask                = "SELECT * FROM task WHERE id = ?"
	queryFindTaskByUserIdAndDay = ("SELECT * FROM task WHERE user_id=? AND created_date LIKE ?;")
	queryFindTaskByContent      = ("SELECT * FROM task WHERE content=?;")
	queryInsertTask             = ("INSERT INTO task(id, user_id, content, status) VALUE(?, ?, ?, ?);")
	queryUpdate                 = ("UPDATE task SET status=?, user_id=?, content=?, created_date=? WHERE id=?;")
	queryDelete                 = ("DELETE FROM task WHERE id=?;")
)

type mysqlRepository struct {
	mysqlDb *sql.DB
}

func (r *mysqlRepository) FindTaskByContent(content string) (*entity.Task, *utils.ErrorRest) {
	var task entity.Task

	stml, err := r.mysqlDb.Prepare(queryFindTaskByContent)
	if err != nil {
		return nil, utils.ErrInternal("Error when trying to prepare tasks")
	}

	defer stml.Close()

	result := stml.QueryRow(content)

	if err := result.Scan(&task.ID, &task.UserID, &task.Content, &task.Status, &task.CreatedDate); err != nil {
		if strings.Contains(err.Error(), "no rows in result") {
			return nil, nil
		}
		return nil, utils.ErrInternal("Error when trying to scan name")
	}

	return &task, nil
}

func (r *mysqlRepository) GetTaskAll() ([]entity.Task, *utils.ErrorRest) {
	stml, err := r.mysqlDb.Prepare(queryGetTaskAll)
	if err != nil {
		return nil, utils.ErrInternal("Error when trying to prepare tasks")
	}
	defer stml.Close()

	rows, err := stml.Query()
	if err != nil {
		return nil, utils.ErrInternal("Error when trying to get tasks")
	}

	defer rows.Close()

	results := make([]entity.Task, 0)
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Content, &task.Status, &task.CreatedDate); err != nil {
			return nil, utils.ErrInternal("Error when trying to scan")
		}

		results = append(results, task)
	}

	return results, nil
}

func (r *mysqlRepository) CreateTask(task *entity.Task) *utils.ErrorRest {
	stml, err := r.mysqlDb.Prepare(queryInsertTask)
	if err != nil {
		return utils.ErrInternal("Error when trying to prepare task")

	}

	defer stml.Close()

	result, err := stml.Exec(&task.ID, &task.UserID, &task.Content, &task.Status)

	if err != nil {
		return utils.ErrInternal("Error when trying add task")
	}

	_, errInsert := result.LastInsertId()
	if errInsert != nil {
		return utils.ErrInternal("Error when trying add task")

	}

	return nil
}

func (r *mysqlRepository) UpdateTask(task *entity.Task) *utils.ErrorRest {
	stmt, err := r.mysqlDb.Prepare(queryUpdate)
	if err != nil {
		fmt.Printf("Err: %v", err)
		return utils.ErrInternal("Error when trying update task")
	}
	defer stmt.Close()

	_, err = stmt.Exec(&task.Status, &task.UserID, &task.Content, &task.CreatedDate, &task.ID)
	if err != nil {
		return utils.ErrInternal("Error when trying update task repo")
	}

	return nil
}

func (r *mysqlRepository) DeleteTask(id string) *utils.ErrorRest {
	stmt, err := r.mysqlDb.Prepare(queryDelete)
	if err != nil {
		fmt.Printf("Err: %v", err)
		return utils.ErrInternal("Error when trying delete task ")
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return utils.ErrInternal("Error when trying delete task")
	}

	return nil
}

func (r *mysqlRepository) FindTaskByUserIdAndDay(userId string, day string) ([]entity.Task, *utils.ErrorRest) {
	stml, err := r.mysqlDb.Prepare(queryFindTaskByUserIdAndDay)
	if err != nil {
		return nil, utils.ErrInternal("Error when trying to prepare tasks")
	}
	defer stml.Close()

	today := day + "%"

	rows, err := stml.Query(userId, today)
	if err != nil {
		return nil, utils.ErrInternal("Error when trying to get tasks")
	}

	defer rows.Close()

	results := make([]entity.Task, 0)
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Content, &task.Status, &task.CreatedDate); err != nil {
			return nil, utils.ErrInternal("Error when trying to scan")
		}

		results = append(results, task)
	}

	return results, nil
}

func (r *mysqlRepository) GetTask(id string) (*entity.Task, *utils.ErrorRest) {
	var task entity.Task

	stmt, err := r.mysqlDb.Prepare(queryGetTask)
	if err != nil {
		return nil, utils.ErrInternal("Error when trying to prepare tasks")
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)

	if err := result.Scan(&task.ID, &task.UserID, &task.Content, &task.Status, &task.CreatedDate); err != nil {
		if strings.Contains(err.Error(), "no rows in result") {
			return nil, nil
		}
		return nil, utils.ErrInternal("Error when trying to scan name")
	}

	return &task, nil
}
