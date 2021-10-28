package handler

type (
	reqLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	reqAddTask struct {
		Content string `json:"content"`
	}
)
