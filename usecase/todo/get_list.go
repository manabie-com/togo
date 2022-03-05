package todo

import (
	"context"
	"fmt"

	"github.com/khangjig/togo/codetype"
	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/util/jwt"
	"github.com/khangjig/togo/util/myerror"
)

type GetListRequest struct {
	codetype.GetListRequest
	Status *codetype.TodoStatus `json:"status,omitempty" query:"status"`
}

func (u *UseCase) GetList(ctx context.Context, req *GetListRequest) (*ResponseListWrapper, error) {
	myUser, _ := ctx.Value(jwt.MyUserClaim).(*model.User)

	req.Format()

	var (
		order      = fmt.Sprintf("created_at %s", req.SortBy)
		conditions = map[string]interface{}{}
	)

	if req.OrderBy == "title" || req.OrderBy == "status" || req.OrderBy == "edited_at" {
		order = fmt.Sprintf("%s %s", req.OrderBy, req.SortBy)
	}

	if req.Status != nil && req.Status.IsValid() {
		conditions["status"] = *req.Status
	}

	myTodos, total, err := u.TodoRepo.GetList(ctx, myUser.ID, conditions, req.Search, order, req.Paginator)
	if err != nil {
		return nil, myerror.ErrGet(err)
	}

	return &ResponseListWrapper{
		Todos: myTodos,
		Page:  req.Paginator.Page,
		Limit: req.Paginator.Limit,
		Total: total,
	}, nil
}
