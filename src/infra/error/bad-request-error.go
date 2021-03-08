package error

import (
	"fmt"
)

type BadRequestError struct {
	status    int
	errorCode string
	codeType  string
	details   error
}

func (this *BadRequestError) Error() string {
	return fmt.Sprintf("status: %d, errorCode: %v, codeType: %v, details: %v\n", this.status, this.errorCode, this.codeType, this.details)
}

func NewBadRequestError(errorCode string, data error) error {
	err := &BadRequestError{
		status:    400,
		errorCode: errorCode,
		codeType:  "BadRequestError",
		details:   data,
	}

	fmt.Printf("[Error] %v", err.Error())
	return err
}
