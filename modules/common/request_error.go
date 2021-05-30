package common

// RequestError is implementation of default error which give information about error status code
type RequestError struct {
	StatusCode  int
	MessageCode string
}

func (r RequestError) Error() string {
	return r.MessageCode
}

// NewInternalError create new RequestError instance with status code 500
func NewInternalError() *RequestError {
	return &RequestError{
		StatusCode:  500,
		MessageCode: "error_default",
	}
}

// NewCustomInternalError create new NewCustomInternalError instance with status code 500
func NewCustomInternalError(messageCode string) *RequestError {
	if messageCode == "" {
		messageCode = "error_default"
	}
	return &RequestError{
		StatusCode:  500,
		MessageCode: messageCode,
	}
}

// NewBadRequestError create new RequestError instance with status code 400 and input message code
func NewBadRequestError(messageCode string) *RequestError {
	return &RequestError{
		StatusCode:  400,
		MessageCode: messageCode,
	}
}

// NewNotFoundError create new RequestError instance with status code 404 and input message code
func NewNotFoundError(messageCode string) *RequestError {
	if messageCode == "" {
		messageCode = "error_not_found"
	}

	return &RequestError{
		StatusCode:  404,
		MessageCode: messageCode,
	}
}

// NewPermissionError create new RequestError instance with status code 403 and input message code
func NewPermissionError(messageCode string) *RequestError {
	if messageCode == "" {
		messageCode = "error_no_permission"
	}

	return &RequestError{
		StatusCode:  403,
		MessageCode: messageCode,
	}
}

// NewUnAuthorizeError create new NewUnAuthorizeError instance with status code 401
func NewUnAuthorizeError(messageCode string) *RequestError {
	if messageCode == "" {
		messageCode = "error_unauthorize"
	}
	return &RequestError{
		StatusCode:  401,
		MessageCode: messageCode,
	}
}
