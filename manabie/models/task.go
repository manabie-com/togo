package models

import (
	"manabie/manabie/services"
	"time"

	"gorm.io/gorm"
)

// Task struct
type Task struct {
	ID          string `gorm:"column:id" json:"id"`
	Content     string `gorm:"column:content" json:"content"`
	UserID      string `gorm:"column:user_id" json:"user_id"`
	CreatedDate string `gorm:"column:created_date" json:"created_date"`
}

// TableName func
func (Task) TableName() string {
	return "tasks"
}

// BeforeCreate func
func (object *Task) BeforeCreate(db *gorm.DB) (err error) {
	now := time.Now()
	object.CreatedDate = now.Format("2006-01-02")
	object.ID = services.GenerateUserID()
	return
}

// PassBodyJSONToModel func
func (u *Task) PassBodyJSONToModel(JSONObject map[string]interface{}) {
	var (
		res interface{}
		val string
	)
	val, res = services.ConvertJSONValueToVariable("Content", JSONObject)
	if res != nil {
		u.Content = val
	}
	return
}
