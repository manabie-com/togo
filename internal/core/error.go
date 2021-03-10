package core

const ERROR_CODE_EXCEED_TASK_LIMITS = 1

var ERROR_EXCEED_TASK_LIMITS = &InternalError{
	ErrCode: ERROR_CODE_EXCEED_TASK_LIMITS,
	Detail:  "Exceed maximum task per day",
}

type InternalError struct {
	ErrCode uint8
	Detail  string
}

func (err *InternalError) Error() string {
	return err.Detail
}
