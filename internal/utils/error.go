package utils


type NotFoundError struct {
}

func (e *NotFoundError) Error() string{
	return "NotFoundID"
}