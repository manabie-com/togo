package entity

import "time"

type SearchArgs struct {
	IsDone      *bool
	UserId      int32
	CreatedDate time.Time
}
