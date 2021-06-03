package tasks

import (
	"net/http"

	"github.com/manabie-com/togo/internal/consts"
	requestUtils "github.com/manabie-com/togo/internal/utils/request"
)

type IGetRequest interface {
	Bind(*http.Request) error
	Validate() error
	GetCreateDate() string
}
type GetRequest struct {
	CreateDate string
}

var RequestParams = struct {
	CreatedDate string
}{
	CreatedDate: "created_date",
}

// Bind
func (h *GetRequest) Bind(req *http.Request) error {
	h.CreateDate = requestUtils.QueryParam(req, RequestParams.CreatedDate)
	if h.CreateDate == "" {
		return consts.ErrInvalidParamWithName(RequestParams.CreatedDate)
	}
	return nil
}

// Validate read request
func (h GetRequest) Validate() error {
	// TODO: validate all parameters parsed from request
	return nil
}

func (h GetRequest) GetCreateDate() string {
	return h.CreateDate
}
