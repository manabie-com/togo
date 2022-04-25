package errors

import "fmt"

const (
	BadRequestContext    = 400
	InternalErrorContext = 503
)

const (
	BadRequestMessage   = "bad.request"
	InteralErrorMessage = "internal.error"
)

const (
	TodoTaskRequestInvalid  = "001"
	ExceedDailyLimitRecords = "002"
	UserIdNotFound          = "003"
	UnexpectedError         = "004"
)

type ErrorInfor struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

type CustomError struct {
	Code       int        `json:"code"`
	Message    string     `json:"message"`
	ErrorInfor ErrorInfor `json:"error"`
}

func (o CustomError) Error() string {
	return fmt.Sprintf("CustomError code = %v desc - %v errors = %v", o.Code, o.Message, o.ErrorInfor)
}

func GetError(errContext int, errMessage string, errCode string) *CustomError {
	errInfor := GenerateErrorInfor(errCode)
	customErr := &CustomError{
		Code:       errContext,
		Message:    errMessage,
		ErrorInfor: errInfor,
	}
	return customErr
}

func GenerateErrorInfor(errCode string) ErrorInfor {
	switch errCode {
	case TodoTaskRequestInvalid:
		return ErrorInfor{
			TodoTaskRequestInvalid, "To Do Task Request is invalid",
			"Invalid or missing requested fields on to do task request body! Please check and try again!",
		}
	case ExceedDailyLimitRecords:
		return ErrorInfor{
			ExceedDailyLimitRecords, "Exceed Daily Limit Records",
			"This user has exceeded daily limit of to do records",
		}
	case UserIdNotFound:
		return ErrorInfor{
			UserIdNotFound, "User Id is not existed",
			"Requested User Id is not existed! Please try again!",
		}
	}
	return ErrorInfor{
		UnexpectedError, "Un-expected error occurs",
		"An un-expected error has occurred",
	}
}
