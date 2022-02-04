package errors

// Error is a custom error with more info compared to the default. It also implements golang's error interface
type Error struct {
	ErrorCode int
	ErrorDesc string
}

func (e Error) Error() string {
	return e.ErrorDesc
}

var Success Error = Error{ErrorCode: 0, ErrorDesc: "Success"}
var MaxLimit Error = Error{ErrorCode: 1, ErrorDesc: "Max daily limit reached"}
