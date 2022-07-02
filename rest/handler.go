package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
	"togo/common/response"
	"togo/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type TogoRestService struct {
	DB *gorm.DB
}

func Handler(db *gorm.DB) TogoRestService {
	return TogoRestService{
		DB: db,
	}
}

// CreateUser is the handler for POST /api/user
// This is neccessary to record username and it's daily task limit
func (restService TogoRestService) CreateUser(w http.ResponseWriter, r *http.Request) {

	userRequest := models.CreateUserRequest{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userRequest)

	// I add validation for username since it's a required string input and based on the database schema
	errInput := validateUsername(userRequest.Username)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	// I add validation for taskDailyLimit because it's a number input
	errInput = validateTaskDailyLimit(userRequest.TaskDailyLimit)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	user := &models.User{}
	newID := uuid.NewV1()
	username := trimLowerUsername(userRequest.Username)
	now := time.Now().UTC()

	err := restService.DB.First(&user, "is_active = true AND LOWER(username) = ?", username).Error

	// I add validation if the user is already existing even the users.username has a unique constraint
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		response.HandleStatusBadRequest(w, "user already exist")
		return
	}

	user = &models.User{
		UserID:         newID,
		Username:       username,
		TaskDailyLimit: userRequest.TaskDailyLimit,
		IsActive:       true,
		CreatedWhen:    now,
		CreatedBy:      newID,
	}

	// Insert the user to the `users` table
	err = restService.DB.Create(&user).Error

	// I return error if the database failed to insert the user
	if err != nil {
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	// I return status created when the insert is success
	response.HandleStatusCreated(w, "user has been successfully created", user)
}

// UpdateUser is the handler for PATCH /api/user (optional to call)
// This enable the user to update the task daily limit differently
func (restService TogoRestService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userRequest := models.UpdateUserRequest{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userRequest)

	// I add validation for username since it's a required string input and based on the database schema
	errInput := validateUsername(userRequest.Username)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	// I add validation for taskDailyLimit because it's a number input
	errInput = validateTaskDailyLimit(userRequest.TaskDailyLimit)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	username := trimLowerUsername(userRequest.Username)

	user := &models.User{}
	now := time.Now().UTC()

	err := restService.DB.First(&user, "is_active = true AND LOWER(username) = ?", username).Error

	// I return error if the user is not existing in users table
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.HandleStatusNotFound(w, "user not found")
		return
	} else if err != nil {
		// I return error if the database failed to find the user
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	user.TaskDailyLimit = userRequest.TaskDailyLimit
	user.UpdatedWhen = &now
	user.UpdatedBy = &user.UserID

	// I update the user to the `users` table
	err = restService.DB.Save(&user).Error

	// I return error if the database failed to update the user
	if err != nil {
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	// I return status ok when the update is success
	response.HandleStatusOK(w, "user has been successfully updated", user)
}

// CreateTask is the handler for POST /api/task
// This enable the user to create task
func (restService TogoRestService) CreateTask(w http.ResponseWriter, r *http.Request) {
	taskRequest := models.CreateTaskRequest{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &taskRequest)

	var count int64
	now := time.Now().UTC()

	// I add validation for username since it's a required string input and based on the database schema
	errInput := validateUsername(taskRequest.Username)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	// I add validation for title since it's a required string input and based on the database schema
	errInput = validateTaskTitle(taskRequest.Title)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	// There's no validation for description since it's nullable and no character limit in the database schema

	username := trimLowerUsername(taskRequest.Username)
	user := &models.User{}
	err := restService.DB.First(&user, "is_active = true AND LOWER(username) = ?", username).Error

	// I return error when the user is not existing in `users` table because the `tasks`` table has relationship to `users` table
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.HandleStatusNotFound(w, "user not found")
		return
	} else if err != nil {
		// I return error if the database fail to find in users table
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	// I get the datetime of the begining of the day in Local time intentionally and use it to get the count of the tasks
	today := getBeginningOfDay(time.Now())
	err = restService.DB.Model(&models.Task{}).
		Where("is_active = true AND user_id = ? AND created_when >= ?", user.UserID, today).
		Count(&count).Error

	// I return error if the user meet the maximum daily limit
	if count >= int64(user.TaskDailyLimit) {
		response.HandleStatusBadRequest(w, "user meet the maximum daily limit")
		return
	}

	newID := uuid.NewV1()
	task := models.Task{}
	task.TaskID = newID
	task.Title = taskRequest.Title
	task.Description = taskRequest.Description

	task.UserID = user.UserID
	task.IsActive = true
	task.CreatedBy = user.UserID
	task.CreatedWhen = now

	// Else, I insert the task in the `tasks` table
	err = restService.DB.Create(&task).Error

	// I return error if the database fail to insert the task
	if err != nil {
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	// I return success when the task has been successfully inserted
	response.HandleStatusCreated(w, "task has been successfully created", task)
}

// DeleteUser is the handler for DELETE /api/user (optional to call)
// I created this for the clean up of integration testing to delete the test user and it's created task
func (restService TogoRestService) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userRequest := models.DeleteUserRequest{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userRequest)

	// I add validation for username since it's a required string input and based on the database schema
	errInput := validateUsername(userRequest.Username)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	username := trimLowerUsername(userRequest.Username)

	user := &models.User{}
	tasks := &[]models.Task{}

	db := restService.DB.First(&user, "is_active = true AND LOWER(username) = ?", username)

	// I return error if the user is not existing in `users` table
	if db.Error != nil {
		response.HandleStatusBadRequest(w, db.Error.Error())
		return
	}

	// I delete the tasks created by the user
	err := restService.DB.Where("is_active = true AND user_id = ?", user.UserID).Delete(&tasks).Error
	// I delete the user in the `users` table
	err = db.Delete(&user, "is_active = true AND user_id = ?", user.UserID).Error

	// I return error if the database failed to delete user or tasks
	if err != nil {
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	// return status ok when the tasks and user has been delete in the table
	response.HandleStatusOK(w, "user and it's tasks has been successfully deleted", user)
}
