package record

type Request struct {
	UserId string `json:"user_id" binding:"required"`
	Task   string `json:"task" binding:"required"`
}