package dto

type CreateTaskRequest struct {
	Title string `json:"title" binding:"required,max=255"`
}
