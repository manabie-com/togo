package common

const (
	BadRequest          = 400
	Unauthorized        = 401
	InternalServerError = 500
	UnknowError         = 505
)

type Error struct {
	Code    ErrorCode
	Message string
}

type ErrorCode int

func (err *Error) withCode(code int) *Error {
	err.Code = ErrorCode(code)
	err.Message = errorMessage[err.Code]

	return err
}

var errorCodeName = map[ErrorCode]string{
	400: "bad_request",
	401: "unauthorized",
	500: "internal_server_error",
	505: "unknow_error",
}

var errorMessage = map[ErrorCode]string{
	400: "Bad Request",
	401: "Unauthorized",
	500: "Internal Server Error",
	505: "Unknow Error",
}
