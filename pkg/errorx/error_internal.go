package errorx

import (
	"fmt"
	"runtime"
	"strings"
)

// ErrFunc ...
type ErrFunc func(error) ErrorInterface

// err ...
type internalError struct {
	Code       string
	StatusCode int
	Err        error
	Title      string
	Message    string
	OrigFile   string
	OrigLine   int
	Meta       map[string]string
}

// Error return ErrFunc
func Error(errCode string, statusCode int, title string, msg string, args ...interface{}) ErrFunc {
	return func(err error) ErrorInterface {
		// Always include the original location
		_, file, line, _ := runtime.Caller(1)
		var xerr *internalError
		meta := map[string]string{}
		message := msg
		if err != nil {
			var ok bool
			if xerr, ok = err.(*internalError); ok {
				// Keep original message
				if xerr.Err != nil {
					meta["cause"] = xerr.Err.Error()
				}
			} else {
				meta["cause"] = err.Error()
			}
		}

		return &internalError{
			Code:       errCode,
			StatusCode: statusCode,
			Err:        err,
			Title:      title,
			Message:    message,
			OrigFile:   file,
			OrigLine:   line,
			Meta:       meta,
		}

	}
}

// Error return ErrorInterface with args...
//func Errorf(code string, err error, message string, args ...interface{}) ErrorInterface {
//	if len(args) > 0 {
//		message = fmt.Sprintf(message, args...)
//	}
//	return newError(code, message, err)
//}

func (t *internalError) GetCode() string {
	return t.Code
}

func (t *internalError) GetStatusCode() int {
	return t.StatusCode
}

func (t *internalError) Msg() string {
	return t.Message
}

func (t *internalError) GetTitle() string {
	return t.Title
}

func (t *internalError) GetMeta(key string) string {
	meta := t.Meta
	if meta != nil {
		return meta[key]
	}
	return ""
}

func (t *internalError) WithMeta(key string, val string) ErrorInterface {
	t.Meta[key] = val
	return t
}

func (t *internalError) MetaMap() map[string]string {
	return t.Meta
}

func (t *internalError) Error() string {
	var b strings.Builder
	b.WriteString(t.Message)

	for k, v := range t.Meta {
		b.WriteByte(' ')
		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(v)
	}
	return b.String()
}

func (t *internalError) Location() string {
	return fmt.Sprintf("%v:%v", t.OrigFile, t.OrigLine)
}

func ToErrorInterface(err error) ErrorInterface {
	xerr, ok := err.(*internalError)
	if !ok {
		return ErrInternal(err)
	}
	return xerr
}
