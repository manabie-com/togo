package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/manabie-com/togo/models"
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

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var data utils.LoginRequest

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.ValidateStruct(data); err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.MapErrors(err.(validator.ValidationErrors)))
		return
	}

	user, err := controller.UserService.GetUserService(data.Username)

	if err != nil {
		utils.JSON(w, http.StatusUnauthorized, map[string]string{"message": "Username/Password is invalid"})
		return
	}

	if err := utils.VerifyPassword(user.Password, data.Password); err != nil {
		utils.JSON(w, http.StatusUnauthorized, map[string]string{"message": "Username/Password is invalid"})
		return
	}

	token, err := utils.CreateToken(data.Username)

	if err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	utils.JSON(w, http.StatusCreated, map[string]string{"token": token})
}

func (controller *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	if !ok {
		utils.ERROR(w, http.StatusUnprocessableEntity, nil)
		return
	}

	if _, err := controller.UserService.GetUserService(username); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, map[string]string{"message": "Access denied"})
		return
	}

	queryParams := r.URL.Query()

	tasks, err := controller.TaskService.GetTasksByUserName(username, queryParams.Get("created_at"))

	if err != nil {
		utils.JSON(w, http.StatusOK, []string{})
		return
	}

	utils.JSON(w, http.StatusOK, tasks)
}

func (controller *TaskController) AddTask(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	if !ok {
		utils.ERROR(w, http.StatusUnprocessableEntity, nil)
		return
	}

	var data utils.AddTaskRequest

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.ValidateStruct(data); err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.MapErrors(err.(validator.ValidationErrors)))
		return
	}

	user, err := controller.UserService.GetUserService(username)

	if err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, map[string]string{"message": "Access denied"})
		return
	}

	count, err := controller.TaskService.Count(username)

	if err != nil {
		utils.JSON(w, http.StatusBadGateway, map[string]string{"message": "Please try again"})
		return
	}

	if count >= user.MaxTodo {
		utils.JSON(w, http.StatusUnprocessableEntity, map[string]string{"message": "You have reached a limit for adding task"})
		return
	}

	task, err := controller.TaskService.CreateTask(&models.Task{Username: username, Content: data.Content})

	if err != nil {
		utils.JSON(w, http.StatusBadGateway, map[string]string{"message": "Please try again"})
		return
	}

	utils.JSON(w, http.StatusCreated, *task)
}
