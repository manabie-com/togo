package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/manabie-com/togo/models"
	u "github.com/manabie-com/togo/utils"
)

var GetTasks = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// get user id here
	var userId = 1
	rows, err := db.Query(`SELECT name, content, created_at FROM tasks WHERE user_id = $1`, userId)
	if err != nil {
		u.Respond(w, http.StatusNotFound, "Failure", err.Error(), nil)
		return
	}
	var tasks []*models.Task

	for rows.Next() {
		var task = &models.Task{}
		rows.Scan(&task.Name, &task.Content, &task.CreatedAt)
		tasks = append(tasks, task)
	}
	u.Respond(w, http.StatusOK, "Success", "Success", tasks)
}

// var GetTask = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	// check taskID exist
// 	// if not exists
// 	u.Respond(w, http.StatusNotFound, map[string]interface{}{})
// 	// else
// 	u.Respond(w, http.StatusOK, map[string]interface{}{})
// }

var Add = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// decode userid from jwt => userId
	task := &models.Task{
		CreatedAt: time.Now().UTC(),
		UserId:    1,
	}
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "invalid request", nil)
		return
	}
	validate := validator.New()

	if err = validate.Struct(task); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", err.Error(), nil)
		return
	}
	// insert database
	err = db.QueryRow(`INSERT INTO tasks(name, content, created_at, user_id) VALUES($1, $2, $3, $4, $5) RETURNING name, email`, task.Name, task.Content, task.CreatedAt, task.UserId).Scan(&task.Name, &task.Content)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", err.Error(), nil)
		return
	}
	u.Respond(w, http.StatusCreated, "Success", "Success creates task", map[string]interface{}{
		"name":       task.Name,
		"content":    task.Content,
		"created_at": task.CreatedAt,
	})

	// check task number today greater than or equal to current task limitDayTasks
	// if true
	// u.Respond(w, http.StatusNotAcceptable, map[string]interface{}{})
	// // else
	// // update task number today += 1
	// u.Respond(w, http.StatusCreated, map[string]interface{}{})
}

// var Edit = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// id := r.URL.Query().Get("id")
// err := db.QueryRow(`UPDATE tasks SET name = $1 WHERE id = 3; RETURNING name, email`, id).Scan(&task.Name, &task.Content)
// if err != nil {
// 	u.Respond(w, http.StatusNotFound, "Failure", err.Error(), nil)
// 	return
// }
// check taskID exist
// if not exists
// u.Respond(w, http.StatusNotFound, map[string]interface{}{})
// // else
// u.Respond(w, http.StatusOK, map[string]interface{}{})
// }
