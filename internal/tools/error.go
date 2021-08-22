package tools

type TodoError struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
}

func (err TodoError) Error() string {
	return err.ErrorMessage
}

func (err TodoError) ToRes() interface{} {
	return err
}

func NewTodoError(code int, errMessage string) *TodoError {
	return &TodoError{
		Code:         code,
		ErrorMessage: errMessage,
	}
}
