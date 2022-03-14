package request

type CreateTodoReq struct {
	UserID  int    `json:"user_id"`
	Content string `json:"content" validate:"required,lte=255"`
}

type UpdateTodoReq struct {
	ID      int    `json:"id" validate:"required"`
	Content string `json:"content"`
	Status  int    `json:"status" validate:"gte=0,lte=1"`
}

type GetTodosReq struct {
	UserID int
	Size   int `json:"size"`
	Index  int `json:"index"`
}
