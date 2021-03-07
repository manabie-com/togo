package error

import "fmt"

type NotFoundError struct {
	status    int
	errorCode string
	codeType  string
	details   error
}

func (this *NotFoundError) Error() string {
	return fmt.Sprintf("status: %d, errorCode: %v, codeType: %v, details: %v\n", this.status, this.errorCode, this.codeType, this.details)
}

func NewNotFoundError(errorCode string, data error) error {
	err := &NotFoundError{
		status:    404,
		errorCode: errorCode,
		codeType:  "NotFoundError",
		details:   data,
	}

	fmt.Printf("[Error] %v", err.Error())
	return err
}
