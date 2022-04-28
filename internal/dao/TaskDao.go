package dao

import "gorm.io/gorm"

type TaskDao interface {
	CountDailyTask(id uint, db *gorm.DB) int64
}

// CountDaily returns the number of tasks created today by the user
func CountDailyTask(id uint, db *gorm.DB) int64 {

	var result int64
	db.Table("tasks").Where("user_id = ? AND created_at >= Date()", id).Count(&result)
	return result

}
