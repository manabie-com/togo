package error

import (
	"fmt"
)

type InternalServerError struct {
	status    int
	errorCode string
	codeType  string
	details   error
}

func (this *InternalServerError) Error() string {
	return fmt.Sprintf("status: %d, errorCode: %v, codeType: %v, details: %v\n", this.status, this.errorCode, this.codeType, this.details)
}

func NewInternalServerError(errorCode string, data error) error {
	err := &InternalServerError{
		status:    500,
		errorCode: errorCode,
		codeType:  "InternalServerError",
		details:   data,
	}

	fmt.Printf("[Error] %v", err.Error())
	return err
}
