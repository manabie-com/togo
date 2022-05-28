package task

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dinhquockhanh/togo/internal/app/auth"
	"github.com/dinhquockhanh/togo/internal/app/limit"
	"github.com/dinhquockhanh/togo/internal/app/user"
	"github.com/dinhquockhanh/togo/internal/pkg/errors"
	"github.com/dinhquockhanh/togo/internal/pkg/http/response"
	"github.com/gin-gonic/gin"
)

type (
	Service interface {
		CreateTask(ctx context.Context, req *CreateTaskReq) (*Task, error)
		AssignTask(ctx context.Context, req *AssignTaskReq) (*Task, error)
		GetByID(ctx context.Context, req *GetTaskByIdReq) (*Task, error)
		ListTasks(ctx context.Context, req *ListTasksReq) ([]*Task, error)
		Delete(ctx context.Context, req *DeleteTaskByIdReq) error
		CountTasksOfUserToDay(ctx context.Context, username string) (int, error)
	}

	Handler struct {
		taskSvc  Service
		userSvc  user.Service
		limitSvc limit.Service
	}
)

func NewHandler(ts Service, us user.Service, ls limit.Service) *Handler {
	return &Handler{
		taskSvc:  ts,
		userSvc:  us,
		limitSvc: ls,
	}
}

func (h *Handler) GetByID(gc *gin.Context) {
	var req GetTaskByIdReq
	if err := gc.ShouldBindUri(&req); err != nil {
		response.Error(gc, err)
		return
	}
	task, err := h.taskSvc.GetByID(gc.Request.Context(), &req)
	if err != nil {
		response.Error(gc, err)
		return
	}

	gc.JSON(http.StatusOK, task)

}

func (h *Handler) CreateTask(gc *gin.Context) {
	var req CreateTaskReq
	if err := gc.ShouldBindJSON(&req); err != nil {
		response.Error(gc, err)
		return
	}

	ctx := gc.Request.Context()

	_, err := h.userSvc.GetByUserName(ctx, &user.GetUserByUserNameReq{UserName: req.Assignee})
	if err != nil {
		response.Error(gc, err)
		return
	}

	if usrName := auth.FromContext(ctx).Username; usrName == "" {
		response.Error(gc, &errors.Error{
			Code:    400,
			Message: "this user cannot create task",
		})
		return
	}

	req.Creator = auth.FromContext(ctx).Username

	task, err := h.taskSvc.CreateTask(ctx, &req)
	if err != nil {
		response.Error(gc, err)
		return
	}
	gc.JSON(http.StatusCreated, task)

}

func (h *Handler) AssignTask(gc *gin.Context) {
	var err error
	var task *Task
	defer func() {
		if err != nil {
			response.Error(gc, err)
			return
		}
		gc.JSON(http.StatusOK, task)
	}()

	ctx := gc.Request.Context()
	var req AssignTaskReq
	if err = gc.ShouldBindJSON(&req); err != nil {
		return
	}

	u, err := h.userSvc.GetByUserName(ctx, &user.GetUserByUserNameReq{UserName: req.Assignee})
	if err != nil {
		return
	}

	l, err := h.limitSvc.GetLimit(ctx, &limit.GetLimitReq{
		TierID: int(u.TierID),
		Action: limit.ReceiveTaskAction,
	})
	if err != nil {
		return
	}

	c, err := h.taskSvc.CountTasksOfUserToDay(ctx, u.Username)
	if err != nil {
		return
	}

	if c > l.Value {
		err = &errors.Error{
			Code:    http.StatusUnauthorized,
			Message: fmt.Sprintf("user %[1]s cannot receive tasks today anymore, the maximum task of user %[1]s is %d", u.Username, l.Value),
		}
		return
	}

	task, err = h.taskSvc.AssignTask(gc.Request.Context(), &req)
	if err != nil {
		return
	}

}
