package common

type AppError struct {
	Code int
	Err error
}

func (e *AppError) Error() string {
	return e.Err.Error()
}