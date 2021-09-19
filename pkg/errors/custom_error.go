package errors

type CustomError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	TraceErr string `json:"-"`
}

func (err CustomError) Error() string {
	return err.Message
}

func (err CustomError) StatusCode() int {
	return err.Code
}

// NewCustomError return custom error
func NewCustomError(mess string, code int) CustomError {
	return CustomError{
		Code:    code,
		Message: mess,
	}
}
