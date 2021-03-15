package tasks

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"togo/src/common/types"
	"togo/src/modules/users"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint       `json:"id" gorm:"col:id,primary_key"`
	Content   string     `json:"content" gorm:"col:content"`
	UserID    uint       `json:"user_id" gorm:"col:content"`
	CreatedBy users.User `gorm:"foreignkey:UserID"`
	CreatedAt time.Time  `json:"created_at" gorm:"col:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"col:updated_at"`
}

func (t *Task) ToJSON() types.JSON {
	return types.JSON{
		"id":         t.ID,
		"content":    t.Content,
		"user_id":    t.UserID,
		"createdBy":  t.CreatedBy.ToJSON(),
		"created_at": t.CreatedAt,
		"updated_at": t.UpdatedAt,
	}
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	var count int64
	userId := t.CreatedBy.ID
	current_time := time.Now()
	year := current_time.Year()
	month := current_time.Month()
	date := current_time.Day()
	startDateStr := strings.Join([]string{
		strconv.Itoa(year),
		"-",
		strconv.Itoa(int(month)),
		"-", strconv.Itoa(date),
		"T00:00:00.000Z",
	}, "")
	endDateStr := strings.Join([]string{
		strconv.Itoa(year),
		"-",
		strconv.Itoa(int(month)),
		"-", strconv.Itoa(date),
		"T23:59:59.999Z",
	}, "")
	tx.Model(&Task{}).Where(
		"user_id = ? AND created_at > ? AND created_at < ?",
		userId,
		startDateStr,
		endDateStr,
	).Count(&count)

	fmt.Printf(strconv.Itoa(int(count)))

	if count > 4 {
		return errors.New("Maximun tasks in day is 5")
	}
	return
}
