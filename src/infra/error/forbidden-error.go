package error

import (
	"fmt"
)

type ForbiddenError struct {
	status    int
	errorCode string
	codeType  string
	details   error
}

func (this *ForbiddenError) Error() string {
	return fmt.Sprintf("status: %d, errorCode: %v, codeType: %v, details: %v\n", this.status, this.errorCode, this.codeType, this.details)
}

func NewForbiddenError(errorCode string, data error) error {
	err := &ForbiddenError{
		status:    503,
		errorCode: errorCode,
		codeType:  "ForbiddenError",
		details:   data,
	}

	fmt.Printf("[Error] %v", err.Error())
	return err
}
