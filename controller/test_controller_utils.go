package controller

import (
	"net/http"
	"net/http/httptest"
	"time"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/qgdomingo/todo-app/model"
	"github.com/qgdomingo/todo-app/mock"
)

// These are functions used by the unit testing for Task and User controllers

func createRegularContext() (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header),
    }

	return w, ctx
}

func createMockTaskData(length int) []model.Task {
	taskList := []model.Task{}
	if length > 0 {
		for i := 1; i <= length; i++ {
			task := model.Task {
				ID: i,
				Title: "Sample Task Title",
				Description: "Sample Task Description",
				Username: "todo_test_user",
				CreateDate: time.Now() }
			taskList = append(taskList, task)
		}
	}
	return taskList
}

func createMockUserDetails(isEmpty bool) []model.UserDetails {
	userList := []model.UserDetails{}
	if !isEmpty {
		user := model.UserDetails {
			Username: "todo_test_user",
			Name: "Todo Test User",
			Email: "todotestuser@sample.com",
			TaskLimit: 10,
		}
		userList = append(userList, user)
	}
	return userList
}

func createMockErrorMessage(message string, errMsg string) map[string]string {
	errMessage := make(map[string]string)
	errMessage["message"] = message
	errMessage["error"] = errMsg
	return errMessage
}

func createMockTaskJSON(c *gin.Context, title string, desc string, username string) ([]byte, error) {
	c.Request.Method = "POST"
    c.Request.Header.Set("Content-Type", "application/json")

	taskDetails := model.TaskUserEnteredDetails {
		Title: title,
		Description: desc,
		Username: username,
	}
	jsonbytes, err := json.Marshal(taskDetails)

	return jsonbytes, err
}

func createMockUserLoginJSON(c *gin.Context, username string, password string) ([]byte, error) {
	c.Request.Method = "POST"
    c.Request.Header.Set("Content-Type", "application/json")

	userLogin := model.UserLogin {
		Username: username,
		Password: password,
	}
	jsonbytes, err := json.Marshal(userLogin)

	return jsonbytes, err
}

func createMockUserRegisterJSON(c *gin.Context, username string, name string, email string, password string, repeatPassword string, taskLimit int) ([]byte, error) {
	c.Request.Method = "POST"
    c.Request.Header.Set("Content-Type", "application/json")

	newUser := model.NewUser {
		Username: username,
		Name: name,
		Email: email,
		Password: password,
		RepeatPassword: repeatPassword,
		TaskLimit: taskLimit,
	}
	jsonbytes, err := json.Marshal(newUser)

	return jsonbytes, err
}

func createMockUserUpdateJSON(c *gin.Context, username string, name string, email string, taskLimit int) ([]byte, error) {
	c.Request.Method = "POST"
    c.Request.Header.Set("Content-Type", "application/json")

	userDetails := model.UserDetails {
		Username: username,
		Name: name,
		Email: email,
		TaskLimit: taskLimit,
	}
	jsonbytes, err := json.Marshal(userDetails)

	return jsonbytes, err
}

func createMockUserPwdChangeJSON(c *gin.Context, currentPassword string, newPassword string, repeatPassword string) ([]byte, error) {
	c.Request.Method = "POST"
    c.Request.Header.Set("Content-Type", "application/json")

	userPwdDetails := model.UserNewPassword {
		CurrentPassword: currentPassword,
		NewPassword: newPassword,
		RepeatPassword: repeatPassword,
	}
	jsonbytes, err := json.Marshal(userPwdDetails)

	return jsonbytes, err
}

func getTaskController(taskList []model.Task, isTaskSuccessful bool, errorMessage map[string]string) (*TaskController) {
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: taskList,
		IsTaskSuccessful: isTaskSuccessful,
		ErrorMessage: errorMessage }

	return &TaskController{ TaskRepo : &taskRepoMock }
}

func getUserController(userList []model.UserDetails, isTaskSuccessful bool, errorMessage map[string]string) (*UserController) {
	userRepoMock := mock.UserRepositoryMock {
		UserList: userList,
		IsTaskSuccessful: isTaskSuccessful,
		ErrorMessage: errorMessage }

	return &UserController{ UserRepo : &userRepoMock }
}
