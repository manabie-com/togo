package apierrors

import (
	"fmt"
)

// var (
// 	UserAlreadyExists = errors.New("User already exists")
// 	UserDoesNotExists = errors.New("User does not exists")
// 	MaxTasksReached   = errors.New("Max tasks created today has been reached")
// )

type BLError interface {
	Param() string
	Value() interface{}
	Error() string
}

type UserAlreadyExistsError struct {
	param string
	value interface{}
}

func NewUserAlreadyExistsError(param string, value interface{}) *UserAlreadyExistsError {
	return &UserAlreadyExistsError{
		param: param,
		value: value,
	}
}

func (err *UserAlreadyExistsError) Param() string {
	return err.param
}

func (err *UserAlreadyExistsError) Value() interface{} {
	return err.value
}

func (err *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("User \"%s\" already exists", err.value)
}

type UserDoesNotExistsError struct {
	param string
	value interface{}
}

func NewUserDoesNotExistsError(param string, value interface{}) *UserDoesNotExistsError {
	return &UserDoesNotExistsError{
		param: param,
		value: value,
	}
}

func (err *UserDoesNotExistsError) Param() string {
	return err.param
}

func (err *UserDoesNotExistsError) Value() interface{} {
	return err.value
}

func (err *UserDoesNotExistsError) Error() string {
	return fmt.Sprintf("User \"%s\" does not exists", err.value)
}

type MaxTasksReachedError struct {
	param string
	value interface{}
}

func NewMaxTasksReachedError(value interface{}) *MaxTasksReachedError {
	return &MaxTasksReachedError{
		value: value,
		param: "max_task",
	}
}

func (err *MaxTasksReachedError) Param() string {
	return err.param
}

func (err *MaxTasksReachedError) Value() interface{} {
	return err.value
}

func (err *MaxTasksReachedError) Error() string {
	return fmt.Sprintf("Max of %d tasks per day has been reached", err.value)
}
