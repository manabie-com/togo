package model

import (
	"time"
)

// UserStatus contain user's status
type UserStatus string

// UserStatus define
const (
	UserStatusActive   UserStatus = "active"
	UserStatusInActive UserStatus = "in_active"
	UserStatusDelete   UserStatus = "delete"
)

// User contain info of user
type User struct {
	ID             int        `json:"_id,omitempty" bson:"_id,omitempty"`
	Username       string     `json:"username,omitempty" bson:"username,omitempty"`
	HashedPassword string     `json:"-" bson:"hashed_password,omitempty"`
	Status         UserStatus `json:"status,omitempty" bson:"status,omitempty"`
	// maxTasks max num task of user
	// note add "omitempty" in bson tag for supprot update to 0
	MaxTasks int `json:"max_tasks" bson:"max_tasks"`
	// currentTasks num of current task of user
	CurrentTasks      int        `json:"current_tasks" bson:"current_tasks"`
	ChangedPasswordAt *time.Time `json:"changed_password_at,omitempty" bson:"changed_password_at,omitempty"`
	// Tracing
	UpdatedIP   string     `json:"updated_ip,omitempty" bson:"updated_ip,omitempty"`
	UpdatedDate *time.Time `json:"updated_date,omitempty" bson:"updated_date,omitempty"`
	//
	CreatedIP   string     `json:"created_ip,omitempty" bson:"created_ip,omitempty"`
	CreatedDate *time.Time `json:"created_date,omitempty" bson:"created_date,omitempty"`
}

// IsExists check user's existed
func (m *User) IsExists() bool {
	return m.ID > 0
}

// IsActive check user's active or not
func (m *User) IsActive() bool {
	return m.Status == UserStatusActive
}

// CanCreateNewTask check user can created new task or not ?
func (m *User) CanCreateNewTask() bool {
	return m.CurrentTasks < m.MaxTasks
}
