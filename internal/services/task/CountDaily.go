package task

import (
	"togo/internal/connect"
)

func CountDaily(id uint) int64 {

	var result int64
	connect.DB.Table("tasks").Where("user_id = ? AND created_at >= Date()", id).Count(&result)

	return result

}
