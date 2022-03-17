package apierrors

import (
	"fmt"
)

// BLError (BusinessLogicError) is the interface for the custom errors
type BLError interface {
	Param() string
	Value() interface{}
	Error() string
}

// UserAlreadyExistsError is the error used when the given user already exists
type UserAlreadyExistsError struct {
	param string
	value interface{}
}

// NewUserAlreadyExistsError is the constructor for UserAlreadyExistsError
func NewUserAlreadyExistsError(param string, value interface{}) *UserAlreadyExistsError {
	return &UserAlreadyExistsError{
		param: param,
		value: value,
	}
}

// Param returns the error parameter
func (err *UserAlreadyExistsError) Param() string {
	return err.param
}

// Value returns the error value
func (err *UserAlreadyExistsError) Value() interface{} {
	return err.value
}

// Error returns the error in string
func (err *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("User \"%s\" already exists", err.value)
}

// UserDoesNotExistsError is the error used when the given user does not exists
type UserDoesNotExistsError struct {
	param string
	value interface{}
}

// NewUserDoesNotExistsError is the constructor for UserDoesNotExistsError
func NewUserDoesNotExistsError(param string, value interface{}) *UserDoesNotExistsError {
	return &UserDoesNotExistsError{
		param: param,
		value: value,
	}
}

// Param returns the error parameter
func (err *UserDoesNotExistsError) Param() string {
	return err.param
}

// Value returns the error value
func (err *UserDoesNotExistsError) Value() interface{} {
	return err.value
}

// Error returns the error in string
func (err *UserDoesNotExistsError) Error() string {
	return fmt.Sprintf("User \"%s\" does not exists", err.value)
}

// MaxTasksReachedError is the error used when the user already reached its maximum number of tasks
type MaxTasksReachedError struct {
	param string
	value interface{}
}

// NewMaxTasksReachedError is the constructor for MaxTasksReachedError
func NewMaxTasksReachedError(value interface{}) *MaxTasksReachedError {
	return &MaxTasksReachedError{
		value: value,
		param: "max_task",
	}
}

// Param returns the error parameter
func (err *MaxTasksReachedError) Param() string {
	return err.param
}

// Value returns the error value
func (err *MaxTasksReachedError) Value() interface{} {
	return err.value
}

// Error returns the error in string
func (err *MaxTasksReachedError) Error() string {
	return fmt.Sprintf("Max of %d tasks per day has been reached", err.value)
}
