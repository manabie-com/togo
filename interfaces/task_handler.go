package interfaces

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jfzam/togo/application"
	"github.com/jfzam/togo/domain/entity"
	"github.com/jfzam/togo/infrastructure/auth"
)

type Task struct {
	TaskApp application.TaskAppInterface
	userApp application.UserAppInterface
	tk      auth.TokenInterface
	rd      auth.AuthInterface
}

// Task constructor
func NewTask(fApp application.TaskAppInterface, uApp application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Task {
	return &Task{
		TaskApp: fApp,
		userApp: uApp,
		rd:      rd,
		tk:      tk,
	}
}

func (fo *Task) SaveTask(c *gin.Context) {
	// check is the user is authenticated first
	metadata, err := fo.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	// lookup the metadata in redis:
	userId, err := fo.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var saveTaskError = make(map[string]string)

	title := c.PostForm("title")
	description := c.PostForm("description")
	if fmt.Sprintf("%T", title) != "string" || fmt.Sprintf("%T", description) != "string" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "Invalid json",
		})
		return
	}

	emptyTask := entity.Task{}
	emptyTask.Title = title
	emptyTask.Description = description
	saveTaskError = emptyTask.Validate("")
	if len(saveTaskError) > 0 {
		c.JSON(http.StatusUnprocessableEntity, saveTaskError)
		return
	}

	//check if the user exist
	_, err = fo.userApp.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, "user not found, unauthorized")
		return
	}

	var Task = entity.Task{}
	Task.UserID = userId
	Task.Title = title
	Task.Description = description
	savedTask, saveErr := fo.TaskApp.SaveTask(&Task)
	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, saveErr)
		return
	}
	c.JSON(http.StatusCreated, savedTask)
}

func (fo *Task) GetAllTask(c *gin.Context) {
	allTask, err := fo.TaskApp.GetAllTask()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, allTask)
}

func (fo *Task) GetTaskAndCreator(c *gin.Context) {
	TaskId, err := strconv.ParseUint(c.Param("Task_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}
	Task, err := fo.TaskApp.GetTask(TaskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	user, err := fo.userApp.GetUser(Task.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	TaskAndUser := map[string]interface{}{
		"Task":    Task,
		"creator": user.PublicUser(),
	}
	c.JSON(http.StatusOK, TaskAndUser)
}
