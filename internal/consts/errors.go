package consts

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrNotFound       Error = "data not found"
	ErrInternal       Error = "internal server error"
	ErrInvalidAuth    Error = "incorrect user_id/pwd"
	ErrUnauthorized   Error = "unathorized"
	ErrInvalidParam   Error = "invalid param: "
	ErrInvalidRequest Error = "invalid request body"
	ErrMaxTodoReached Error = "max todo reached"
)

func ErrInvalidParamWithName(paramName string) Error {
	return Error(ErrInvalidParam.Error() + paramName)
}
