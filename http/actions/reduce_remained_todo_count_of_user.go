package actions

type ReduceRemainedTodoCountOfUserParam interface {
	GetAssigneeEmail() string // Used to get user
	GetTaskSavedCount() int   // Used to reduce remained todo task per day
}

func ReduceRemainedTodoCountOfUser(p ReduceRemainedTodoCountOfUserParam) (ok bool) {
	ok = true

	return
}
