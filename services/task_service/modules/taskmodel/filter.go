package taskmodel

type Filter struct {
	UserId int   `json:"user_id" form:"-"`
	IsDone *bool `json:"is_done,omitempty" form:"is_done"`
}
