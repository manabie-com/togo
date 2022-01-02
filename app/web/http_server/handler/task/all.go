package task

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	gErrcode "github.com/manabie-com/togo/app/common/gconstant/errcode"
	gHandler "github.com/manabie-com/togo/app/common/gstuff/handler"
	"github.com/manabie-com/togo/app/model"
	"github.com/manabie-com/togo/app/utils/token"

	taskRepo "github.com/manabie-com/togo/app/repo/mongo/task"
)

// All get all task by user
func (s *service) All(c echo.Context) error {
	httpCtx := c.Request().Context()

	type myResponse struct {
		LastOffset string       `json:"last_offset" query:"last_offset"`
		Data       []model.Task `json:"data" query:"data"`
	}

	type myRequest struct {
		Status *model.TaskStatus `json:"status" query:"status" validate:"omitempty,oneof=active in_active"`
		Offset string            `json:"offset" query:"offset"`
		Limit  int               `json:"limit" query:"limit"`
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

	// create get all task req
	allTaskReq := taskRepo.AllReq{
		UserID: &authPayload.UserID,
		Status: request.Status,
		Offset: primitive.NilObjectID,
		Limit:  request.Limit,
	}
	if request.Offset != "" {
		offset, err := primitive.ObjectIDFromHex(request.Offset)
		if err != nil {
			return gHandler.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("offset hex [%v] is invalid: %v", request.Offset, err), gErrcode.UserErrCommon)

		}
		allTaskReq.Offset = offset
	}

	result, err := s.taskRepo.All(httpCtx, allTaskReq)
	if err != nil {
		return gHandler.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("get all task: %s", err), gErrcode.ServerErrorCommon)
	}

	response := myResponse{
		Data: make([]model.Task, 0),
	}
	if len(result) > 0 {
		response.Data = result
		response.LastOffset = result[len(result)-1].ID.Hex()
	}
	return c.JSON(gHandler.Success(response))
}
