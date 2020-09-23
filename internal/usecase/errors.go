package usecase

var (
	ErrMaxTaskReached = &Error{Code: -2000, Message: "User reached maximum number of created task per day"}
)

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

