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

func (restService TogoRestService) CreateUser(w http.ResponseWriter, r *http.Request) {

	userRequest := models.CreateUserRequest{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userRequest)

	errInput := validateUsername(userRequest.Username)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	errInput = validateTaskDailyLimit(userRequest.TaskDailyLimit)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	user := &models.User{}
	newID := uuid.NewV1()
	username := trimLowerUsername(userRequest.Username)
	now := time.Now().UTC()

	err := restService.DB.First(&user, "is_active = true AND username = ?", username).Error

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

	err = restService.DB.Create(&user).Error

	if err != nil {
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	response.HandleStatusCreated(w, "user has been successfully created", user)
}

func (restService TogoRestService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userRequest := models.UpdateUserRequest{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userRequest)

	errInput := validateUsername(userRequest.Username)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	errInput = validateTaskDailyLimit(userRequest.TaskDailyLimit)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	username := trimLowerUsername(userRequest.Username)

	user := &models.User{}
	now := time.Now().UTC()

	err := restService.DB.First(&user, "is_active = true AND username = ?", username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.HandleStatusNotFound(w, "user not found")
		return
	} else if err != nil {
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	user.TaskDailyLimit = userRequest.TaskDailyLimit
	user.UpdatedWhen = &now
	user.UpdatedBy = &user.UserID

	err = restService.DB.Save(&user).Error

	if err != nil {
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	response.HandleStatusOK(w, "user has been successfully updated", user)
}

func (restService TogoRestService) CreateTask(w http.ResponseWriter, r *http.Request) {
	taskRequest := models.CreateTaskRequest{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &taskRequest)
	var count int64
	now := time.Now().UTC()

	errInput := validateUsername(taskRequest.Username)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	errInput = validateTaskTitle(taskRequest.Title)
	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	username := trimLowerUsername(taskRequest.Username)
	user := &models.User{}
	err := restService.DB.First(&user, "is_active = true AND username = ?", username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.HandleStatusNotFound(w, "user not found")
		return
	} else if err != nil {
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	today := getBeginningOfDay(time.Now())
	err = restService.DB.Model(&models.Task{}).
		Where("is_active = true AND user_id = ? AND created_when > ?", user.UserID, today).
		Count(&count).Error

	newID := uuid.NewV1()

	if count >= int64(user.TaskDailyLimit) {
		response.HandleStatusBadRequest(w, "user meet the maximum daily limit")
		return
	}

	task := models.Task{}
	task.TaskID = newID
	task.Title = taskRequest.Title
	task.Description = taskRequest.Description

	task.UserID = user.UserID
	task.IsActive = true
	task.CreatedBy = user.UserID
	task.CreatedWhen = now

	err = restService.DB.Create(&task).Error

	if err != nil {
		response.HandleStatusInternalServerError(w, err.Error())
		return
	}

	response.HandleStatusCreated(w, "task has been successfully created", task)
}

func (restService TogoRestService) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userRequest := models.DeleteUserRequest{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userRequest)

	errInput := validateUsername(userRequest.Username)

	if errInput != nil {
		response.HandleStatusBadRequest(w, errInput.Error())
		return
	}

	username := trimLowerUsername(userRequest.Username)

	user := &models.User{}
	tasks := &[]models.Task{}

	db := restService.DB.First(&user, "is_active = true AND LOWER(username) = ?", username)

	if db.Error != nil {
		response.HandleStatusBadRequest(w, db.Error.Error())
		return
	}

	err := restService.DB.Where("is_active = true AND user_id = ?", user.UserID).Delete(&tasks).Error
	err = db.Delete(&user, "is_active = true AND user_id = ?", user.UserID).Error

	if err != nil {
		response.HandleStatusBadRequest(w, err.Error())
		return
	}

	response.HandleStatusOK(w, "user and it's tasks has been successfully deleted", user)
}
