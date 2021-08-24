package tools

import (
	"database/sql"
	"net/http"
)

type IRequestTool interface {
	Value(req *http.Request, p string) sql.NullString
}

type RequestTool struct{}

func (rt *RequestTool) Value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func NewRequestTool() IRequestTool {
	return &RequestTool{}
}
