package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/manabie-com/togo/services"
	"github.com/manabie-com/togo/utils"
	"net/http"
)

type AuthController struct {
	UserService services.IUserService
}

type TaskController struct {
	UserService services.IUserService
	TaskService services.ITaskService
}

func NewAuthController(userService *services.IUserService) AuthController {
	return AuthController{UserService: *userService}
}

func NewTaskController(userService *services.IUserService, taskService *services.ITaskService) TaskController {
	return TaskController{UserService: *userService, TaskService: *taskService}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func NotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (controller *TaskController) AA(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var data utils.LoginRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&data)

	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = utils.ValidateStruct(data)

	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.MapErrors(err.(validator.ValidationErrors)))
		return
	}

	//services.UserService{}

	utils.JSON(w, http.StatusCreated, data)
}
