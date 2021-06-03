package tasks

import (
	"database/sql"
	"net/http"

	"github.com/manabie-com/togo/internal/consts"
	taskService "github.com/manabie-com/togo/internal/services/tasks"
	userService "github.com/manabie-com/togo/internal/services/users"
	requestUtils "github.com/manabie-com/togo/internal/utils/request"
	"github.com/manabie-com/togo/internal/utils/response"
)

func NewHandler(database *sql.DB) TaskHandler {
	return TaskHandler{
		NewCreateRequest: func() ICreateRequest { return &CreateRequest{} },
		NewGetRequest:    func() IGetRequest { return &GetRequest{} },
		TaskService:      taskService.NewService(database),
		UserService:      userService.NewService(database),
	}
}

type TaskHandler struct {
	NewCreateRequest func() ICreateRequest
	NewGetRequest    func() IGetRequest
	TaskService      taskService.ITaskService
	UserService      userService.IUserService
}

// Create new task
func (t TaskHandler) Post(resp http.ResponseWriter, req *http.Request) error {
	request := t.NewCreateRequest()
	if err := request.Bind(req); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return err
	}
	userId, ok := requestUtils.GetUserID(req)
	if !ok {
		return consts.ErrInvalidRequest
	}
	ctx := req.Context()
	count, err := t.TaskService.TaskCountByUser(ctx, userId)
	if err != nil {
		return consts.ErrInvalidRequest
	}
	user, err := t.UserService.GetUserByID(ctx, userId)
	if err != nil {
		return consts.ErrInvalidRequest
	}
	if int(count) >= user.MaxTodo {
		return consts.ErrMaxTodoReached
	}
	model := request.ToModel(userId)
	if err := t.TaskService.CreateNew(req.Context(), model); err != nil {
		return err
	}
	return response.JSON(resp, model)
}

// Get task list
func (t TaskHandler) Get(resp http.ResponseWriter, req *http.Request) error {
	request := t.NewGetRequest()
	if err := request.Bind(req); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return err
	}
	data := t.TaskService.GetTasksCreatedOn(req.Context(), request.GetCreateDate())
	return response.JSON(resp, data)
}
