package task

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"

	gErrcode "github.com/manabie-com/togo/app/common/gconstant/errcode"
	gHandler "github.com/manabie-com/togo/app/common/gstuff/handler"
	"github.com/manabie-com/togo/app/model"
	"github.com/manabie-com/togo/app/utils/token"

	taskRepo "github.com/manabie-com/togo/app/repo/mongo/task"
	userRepo "github.com/manabie-com/togo/app/repo/mongo/user"
)

// Create created new task
func (s *service) Create(c echo.Context) error {
	httpCtx := c.Request().Context()

	type myResponse struct {
		Task            model.Task `json:"task"`
		CurrentUserTask int        `json:"current_user_task"`
	}

	type myRequest struct {
		Name        string `json:"name" query:"name" validate:"required,max=500"`
		Description string `json:"description" query:"description" validate:"required,max=500"`
	}
	request := new(myRequest)
	if err := c.Bind(request); err != nil {
		return gHandler.NewHTTPError(http.StatusBadRequest, err.Error(), gErrcode.UserErrCommon)
	}
	if err := c.Validate(request); err != nil {
		return gHandler.NewHTTPError(http.StatusBadRequest, err.Error(), gErrcode.UserErrCommon)
	}
	// get info enforcer
	if c.Get("authorization_payload") == nil {
		return gHandler.NewHTTPError(http.StatusInternalServerError, "empty info enforcer, check token validate", gErrcode.ServerErrorCommon)
	}

	authPayload := c.Get("authorization_payload").(*token.Payload)

	// 1. check user existed by userid
	existedUser, err := s.userRepo.GetOneByID(httpCtx, authPayload.UserID)
	if err != nil {
		return gHandler.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("get user [%v]: %s", authPayload.UserID, err), gErrcode.ServerErrorCommon)
	}
	if !existedUser.IsExists() {
		return gHandler.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user [%v] is not existed", authPayload.Username), gErrcode.UserErrCommon)
	}
	if !existedUser.IsActive() {
		return gHandler.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user [%v] is not active", authPayload.Username), gErrcode.UserErrCommon)
	}
	if !existedUser.CanCreateNewTask() {
		return gHandler.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user [%v] reach limit tasks [%v]", authPayload.Username, existedUser.MaxTasks), gErrcode.UserErrCommon)
	}

	// ! Create task with transaction lock
	// ! Should use instead with event driven architecture, for optimize performance
	session, err := s.mongoClient.StartSession()
	if err != nil {
		return gHandler.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("get mongo session: %s", err), gErrcode.ServerErrorCommon)
	}
	// Start transaction
	err = session.StartTransaction()
	if err != nil {
		return gHandler.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("start mongo transaction: %s", err), gErrcode.ServerErrorCommon)
	}
	defer session.EndSession(httpCtx)

	// create response
	response := myResponse{}

	err = mongo.WithSession(httpCtx, session, func(sessCtx mongo.SessionContext) (err error) {
		// 1. Create new task
		createTaskReq := taskRepo.CreateReq{
			UserID:      authPayload.UserID,
			Name:        request.Name,
			Description: request.Description,
			// tracing
			CreatedIP: c.RealIP(),
		}
		response.Task, err = s.taskRepo.Create(sessCtx, createTaskReq)
		if err != nil {
			session.AbortTransaction(sessCtx)
			return fmt.Errorf("create task: %s", err)
		}
		// 2. Increase num task of user
		incUserTaskReq := userRepo.IncNumTaskReq{
			UserID:   authPayload.UserID,
			MaxTasks: existedUser.MaxTasks,
			// tracing
			UpdatedIP: c.RealIP(),
		}

		updatedUser, err := s.userRepo.IncNumTask(sessCtx, incUserTaskReq)
		if err != nil {
			session.AbortTransaction(sessCtx)
			return fmt.Errorf("inc user current task: %s", err)
		}
		if !updatedUser.IsExists() {
			session.AbortTransaction(sessCtx)
			return fmt.Errorf("not adapt current task < mask task")
		}

		response.CurrentUserTask = updatedUser.CurrentTasks

		session.CommitTransaction(sessCtx)
		return
	})

	if err != nil {
		return gHandler.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("fail transaction: %s", err), gErrcode.ServerErrorCommon)
	}

	return c.JSON(gHandler.Success(response))
}
