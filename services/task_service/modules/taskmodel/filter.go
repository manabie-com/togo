package taskmodel

import (
	"time"
)

type Filter struct {
	UserId      int        `json:"user_id" form:"-"`
	IsDone      *bool      `json:"is_done,omitempty" form:"is_done"`
	CreatedDate *time.Time `json:"created_date" form:"created_date"  time_format:"2006-01-02"`
}
