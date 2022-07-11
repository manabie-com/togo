package actions

type TodoTaskParam interface {
}

func SaveTodoTask(p TodoTaskParam) (r interface{}) {
	r = "Added"

	return
}
