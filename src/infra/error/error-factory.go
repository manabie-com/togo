package error

import "togo/src"

type ErrorFactory struct {
}

func (this *ErrorFactory) UnauthorizedError(errorCode string, data error) error {
	return NewUnauthorizedError(errorCode, data)
}

func (this *ErrorFactory) NotFoundError(errorCode string, data error) error {
	return NewNotFoundError(errorCode, data)
}

func (this *ErrorFactory) InternalServerError(errorCode string, data error) error {
	return NewInternalServerError(errorCode, data)
}

func (this *ErrorFactory) ForbiddenError(errorCode string, data error) error {
	return NewForbiddenError(errorCode, data)
}

func (this *ErrorFactory) BadRequestError(errorCode string, data error) error {
	return NewBadRequestError(errorCode, data)
}

func NewErrorFactory() src.IErrorFactory {
	return &ErrorFactory{}
}
