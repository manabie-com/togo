package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/models"
	u "github.com/manabie-com/togo/utils"
)

var GetTasks = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// get user id here
	decoded := r.Context().Value("user").(*models.Token)
	rows, err := db.Query(`SELECT * FROM tasks WHERE user_id = $1`, decoded.UserId)
	if err != nil {
		u.Respond(w, http.StatusNotFound, "Failure", err.Error(), nil)
		return
	}

	var tasks []*models.Task

	for rows.Next() {
		var task = &models.Task{}
		rows.Scan(&task.ID, &task.Name, &task.Content, &task.CreatedAt, &task.UserId)
		tasks = append(tasks, task)
	}

	u.Respond(w, http.StatusOK, "Success", "Success", tasks)
}

var GetTask = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// decode token from middleware
	decoded := r.Context().Value("user").(*models.Token)
	task := &models.Task{}

	id := mux.Vars(r)["id"] // get id from url param

	err := db.QueryRow(`SELECT name, content, created_at FROM tasks WHERE id = $1 AND user_id = $2`, id, decoded.UserId).Scan(&task.Name, &task.Content, &task.CreatedAt)
	if err != nil {
		u.Respond(w, http.StatusNotFound, "Failure", "Not found task", nil)
		return
	}

	u.Respond(w, http.StatusOK, "Success", "Success", map[string]interface{}{
		"name":       task.Name,
		"content":    task.Content,
		"created_at": task.CreatedAt,
	})
}

var Add = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// decode token from middleware
	decoded := r.Context().Value("user").(*models.Token)
	task := &models.Task{
		UserId: decoded.UserId,
	}
	// json body -> task object
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "invalid request", nil)
		return
	}
	// validate task object
	validate := validator.New()
	if err = validate.Struct(task); err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", err.Error(), nil)
		return
	}
	// check limit day tasks
	var recordsLength uint
	_ = db.QueryRow(`SELECT COUNT(id)
	FROM tasks 
	WHERE created_at >= NOW() - INTERVAL '24 HOURS' AND user_id = $1`, decoded.UserId).Scan(&recordsLength)
	if recordsLength == decoded.LimitDayTasks {
		u.Respond(w, http.StatusBadRequest, "Failure", "Today tasks had limited, Please Comeback tomorrow.", nil)
		return
	}
	// insert database
	err = db.QueryRow(`INSERT INTO tasks(name, content, user_id) VALUES($3, $2, $1) RETURNING name, content, created_at`, task.UserId, task.Content, task.Name).Scan(&task.Name, &task.Content, &task.CreatedAt)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", err.Error(), nil)
		return
	}
	u.Respond(w, http.StatusCreated, "Success", "Success create task", map[string]interface{}{
		"name":       task.Name,
		"content":    task.Content,
		"created_at": task.CreatedAt,
	})
}

var Edit = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	decoded := r.Context().Value("user").(*models.Token)
	task := &models.Task{}
	id := mux.Vars(r)["id"] // get id from url params
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", err.Error(), nil)
		return
	}
	err = db.QueryRow(`UPDATE tasks SET name = $3 WHERE id = $1 AND user_id = $2 RETURNING name, content, created_at`, id, decoded.UserId, task.Name).Scan(&task.Name, &task.Content, &task.CreatedAt)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", err.Error(), nil)
		return
	}

	u.Respond(w, http.StatusCreated, "Success", "Success create task", map[string]interface{}{
		"name":       task.Name,
		"content":    task.Content,
		"created_at": task.CreatedAt,
	})
}

var Delete = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	decoded := r.Context().Value("user").(*models.Token)
	id := mux.Vars(r)["id"]
	_, err := db.Exec(`DELETE FROM tasks WHERE id = $1 AND user_id = $2`, id, decoded.UserId)
	if err != nil {
		u.Respond(w, http.StatusNotFound, "Failure", err.Error(), nil)
		return
	}
	u.Respond(w, http.StatusCreated, "Success", "Success delete task", nil)
}
