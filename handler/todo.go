package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/common/context"
	"github.com/manabie-com/togo/domain"
	"github.com/manabie-com/togo/domain/model"
	"github.com/manabie-com/togo/handler/dto"
	"log"
	"net/http"
	"time"
)

type TodoHandler interface {
	GetAccessToken(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	GetTasks(ctx *gin.Context)
}

func NewTodoHandler(todoService domain.TodoService) TodoHandler {
	return &todoHandler{
		todoService: todoService,
	}
}

type todoHandler struct {
	todoService domain.TodoService
}

func (t *todoHandler) GetTasks(ctx *gin.Context) {
	log.Println("get tasks")
	appCtx := context.FromContext(ctx)
	createDateString := ctx.Request.URL.Query().Get("created_date")
	createDate, err := time.Parse("2006-01-02", createDateString)
	if err != nil {
		log.Println("can't parse create date")
		dto.ResponseError(ctx, err, http.StatusBadRequest)
		return
	}

	tasks, err := t.todoService.GetTaskAtDate(appCtx, createDate)
	if err != nil {
		log.Println("can't get tasks")
		dto.ResponseError(ctx, err)
		return
	}
	tasksDto := mapTasksModelToDto(tasks)
	dto.ResponseSuccess(ctx, tasksDto)
}

func (t *todoHandler) CreateTask(ctx *gin.Context) {
	log.Println("create task")
	appCtx := context.FromContext(ctx)
	req := &dto.CreateTaskReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		log.Println("invalid request")
		dto.ResponseError(ctx, err)
		return
	}
	if req.Content == "" {
		log.Println("context is empty")
		dto.ResponseError(ctx, errors.New("content is empty"), http.StatusBadRequest)
		return
	}
	task, err := t.todoService.CreateTask(appCtx, req.Content)
	if err != nil {
		log.Println("can't create task: ", err)
		dto.ResponseError(ctx, err)
		return
	}
	taskDto := mapTaskModelToDto(task)
	dto.ResponseSuccess(ctx, taskDto)
}

func (t *todoHandler) GetAccessToken(ctx *gin.Context) {
	log.Println("get access token")
	appCtx := context.FromContext(ctx)
	username := ctx.Request.URL.Query().Get("user_id")
	password := ctx.Request.URL.Query().Get("password")
	token, err := t.todoService.GetAccessToken(appCtx, username, password)
	if err != nil {
		log.Println("can't get access token: ", err)
		dto.ResponseError(ctx, err)
		return
	}
	dto.ResponseSuccess(ctx, token)
}

func mapTaskModelToDto(taskModel *model.Task) *dto.Task {
	return &dto.Task{
		ID:          taskModel.ID,
		Content:     taskModel.Content,
		UserID:      taskModel.UserID,
		CreatedDate: taskModel.CreatedDate.Format("2006-01-02"),
	}
}

func mapTasksModelToDto(tasksModel []*model.Task) []*dto.Task {
	tasksDto := make([]*dto.Task, 0)
	for _, task := range tasksModel {
		tasksDto = append(tasksDto, mapTaskModelToDto(task))
	}
	return tasksDto
}
