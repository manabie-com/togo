package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/pkg/database"
)

// task table tasks
type task struct {
	ID         int16      `gorm:"column:id;primaryKey"`
	Content    string     `gorm:"column:content"`
	UserID     *int16     `gorm:"column:user_id"`
	DateAssign *time.Time `gorm:"column:date_assign"`
	IsDeleted  bool       `gorm:"column:is_deleted"`
}

// NewTask task constructor
func NewTask() *task {
	return &task{}
}

// filter filter task is not deleted
func (*task) filter() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("tasks.is_deleted IS FALSE")
	}
}

// assignToDay filter task has been assigned to day
func (*task) assignToDay() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("date_assign::DATE = now()::DATE")
	}
}

// noAssign filter task is not assign yet
func (*task) noAssign() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id IS NULL")
	}
}

// IsNotAssign check taskIDs is not assign yet
func (t *task) IsNotAssign(taskIDs []int16) (bool, error) {
	tasks := []task{}
	err := database.DB().
		Scopes(
			where("id IN (?)", taskIDs),
			t.filter(),
			t.noAssign(),
		).
		Find(&tasks).Error
	return len(tasks) == len(taskIDs), err
}

// Assign assign taskIDs to userID
func (t *task) Assign(userID int16, taskIDs []int16) error {
	now := time.Now()
	return database.DB().
		Model(task{}).
		Where("id IN (?)", taskIDs).
		Updates(
			task{
				UserID:     &userID,
				DateAssign: &now,
			},
		).Error
}

// Create create task
func (t *task) Create() error {
	return database.DB().
		Create(t).Error
}
