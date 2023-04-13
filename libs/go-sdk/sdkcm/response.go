package sdkcm

import (
	"net/http"
)

// Response helpers
var (
	SimpleSuccessResponse = func(data interface{}) Response {
		return newResponse(http.StatusOK, data, nil, nil)
	}

	ResponseWithPaging = func(data, param interface{}, other interface{}) Response {
		if v, ok := other.(Paging); ok {
			return newResponse(http.StatusOK, data, param, v)
		}
		return newResponse(http.StatusOK, data, param, other)
	}
)

type Response struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Param  interface{} `json:"param,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
}

func newResponse(code int, data, param, other interface{}) Response {
	return Response{
		Code:   code,
		Data:   data,
		Param:  param,
		Paging: other,
	}
}
