package codetype

type TodoStatus uint8

const (
	TodoStatusOpen       TodoStatus = 0
	TodoStatusInProgress TodoStatus = 1
	TodoStatusResolved   TodoStatus = 2
)

func (ts *TodoStatus) IsValid() bool {
	switch *ts {
	case TodoStatusOpen, TodoStatusInProgress, TodoStatusResolved:
		return true
	default:
		return false
	}
}
