package error

import "fmt"

type UnauthorizedError struct {
	status    int
	errorCode string
	codeType  string
	details   error
}

func (this *UnauthorizedError) Error() string {
	return fmt.Sprintf("status: %d, errorCode: %v, codeType: %v, details: %v\n", this.status, this.errorCode, this.codeType, this.details)
}

func NewUnauthorizedError(errorCode string, data error) error {
	err := &UnauthorizedError{
		status:    400,
		errorCode: errorCode,
		codeType:  "UnauthorizedError",
		details:   data,
	}

	fmt.Printf("[Error] %v", err.Error())
	return err
}
