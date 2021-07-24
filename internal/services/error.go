package services


const (
	limitErrorMsg = "users are limited to create only 5 tasks only per day"
)
type serviceError struct {
	Message string
	Status  int
}

func (s *serviceError) Error() string {
	return s.Message
}
func (s *serviceError) StatusCode() int {
	return s.Status
}
func newError(statusCode int, errMsg string) *serviceError {
	return &serviceError{
		Message: errMsg,
		Status:  statusCode,
	}
}
