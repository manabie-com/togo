package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          string `json:"id"           mapstructure:"id"           gorm:"primary_key:true;column:id;not null"`
	Content     string `json:"content"      mapstructure:"content"      gorm:"column:content;not null"`
	UserID      string `json:"user_id"      mapstructure:"user_id"      gorm:"column:user_id;not null"`
	CreatedDate string `json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	User		*User  `gorm:"foreignKey:UserID"`
}

func(Task) TableName() string {
	return "tasks"
}

func(Task) BeforeCreate(tx *gorm.DB) error {
	return nil
}
