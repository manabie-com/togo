package errorx

import (
	"strings"
)

// ErrorJSON sẽ là cấu trúc Error API trả về cho client
type ErrorJSON struct {
	StatusCode int               `json:"status_code"`
	Code       string            `json:"code"`
	Title      string            `json:"title"`
	Msg        string            `json:"msg"`
	Meta       map[string]string `json:"meta,omitempty"`
}

func ToErrorJSON(errInterface ErrorInterface) *ErrorJSON {
	return &ErrorJSON{
		StatusCode: errInterface.GetStatusCode(),
		Code:       errInterface.GetCode(),
		Msg:        errInterface.Msg(),
		Title:      errInterface.GetTitle(),
		Meta:       errInterface.MetaMap(),
	}
}

func (e *ErrorJSON) Error() (s string) {
	if len(e.Meta) == 0 {
		return e.Msg
	}
	b := strings.Builder{}
	b.WriteString(e.Msg)
	b.WriteString(" (")
	for _, v := range e.Meta {
		b.WriteString(v)
		break
	}
	b.WriteString(")")
	return b.String()
}
