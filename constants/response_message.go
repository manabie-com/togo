package constants

import "fmt"

const (
	RequestBodyEmpty  = "Request body is empty."

	SuccessMessage = "Success"
	GeneralFailureMessage = "General Failure"
	BadRequestMessage = "Bad Request"

	// Bad request message

	InvalidEmail = "invalid email"
	MissingEmail = "missing email"
	MissingTask = "missing task"
	MissingTaskLimit = "missing limit"

	InvalidDate = "date must be in format yyyy-mm-dd"

	ExceedTaskPerDayLimit = "your todo tasks exceed limit per day"
)

var (
	LimitTaskOverMaxValue = fmt.Sprintf("Max limit task per day is %v.", MaxTaskPerDay)
)
