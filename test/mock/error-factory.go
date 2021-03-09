package mock

import (
	"errors"
)

type ErrorFactoryMock struct {
	Get400ErrorFunc func(errorCode string, data error) error
	Get404ErrorFunc func(errorCode string, data error) error
	Get500ErrorFunc func(errorCode string, data error) error
	Get503ErrorFunc func(errorCode string, data error) error
}

func (this *ErrorFactoryMock) UnauthorizedError(errorCode string, data error) error {
	return this.Get400ErrorFunc(errorCode, data)
}

func (this *ErrorFactoryMock) NotFoundError(errorCode string, data error) error {
	return this.Get404ErrorFunc(errorCode, data)
}

func (this *ErrorFactoryMock) InternalServerError(errorCode string, data error) error {
	return this.Get500ErrorFunc(errorCode, data)
}

func (this *ErrorFactoryMock) ForbiddenError(errorCode string, data error) error {
	return this.Get503ErrorFunc(errorCode, data)
}

func (this *ErrorFactoryMock) BadRequestError(errorCode string, data error) error {
	return this.Get400ErrorFunc(errorCode, data)
}

var (
	ERROR_400 = errors.New("ERROR_400")
	ERROR_404 = errors.New("ERROR_404")
	ERROR_500 = errors.New("ERROR_500")
	ERROR_503 = errors.New("ERROR_503")
)

func NewErrorFactoryMock() *ErrorFactoryMock {
	return &ErrorFactoryMock{
		Get400ErrorFunc: func(errorCode string, data error) error {
			return ERROR_400
		},
		Get404ErrorFunc: func(errorCode string, data error) error {
			return ERROR_404
		},
		Get500ErrorFunc: func(errorCode string, data error) error {
			return ERROR_500
		},
		Get503ErrorFunc: func(errorCode string, data error) error {
			return ERROR_503
		},
	}
}
