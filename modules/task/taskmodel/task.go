package taskmodel

import "github.com/japananh/togo/common"

const EntityName = "Task"

type TaskState int

const (
	Open TaskState = iota
	InProgress
	Done
	Holding
	Canceled
)

func (s TaskState) String() string {
	return toString[s]
}

var toString = map[TaskState]string{
	Open:       "open",
	InProgress: "inprogress",
	Done:       "done",
	Holding:    "holding",
	Canceled:   "canceled",
}

type Task struct {
	common.SQLModel `json:",inline"`
	CreatedBy       int         `json:"-" gorm:"column:created_by;"`
	AssigneeId      int         `json:"-" gorm:"column:assignee;"`
	ParentId        int         `json:"-" gorm:"column:parent_id;"`
	Title           string      `json:"title"  gorm:"column:title;"`
	Status          string      `json:"status" gorm:"status;default:'open';"`
	Description     string      `json:"description"  gorm:"column:description;"`
	FakeCreatedBy   *common.UID `json:"created_by" binding:"required" gorm:"-"`
	FakeAssigneeId  *common.UID `json:"assignee" gorm:"-"`
	FakeParentId    *common.UID `json:"parent_id" gorm:"-"`
}

func (Task) TableName() string {
	return "tasks"
}

func (t *Task) Mask() {
	t.GenUID(common.DbTypeTask)

	fakeCreatedBy := common.NewUID(uint32(t.CreatedBy), common.DbTypeTask, 1)
	t.FakeCreatedBy = &fakeCreatedBy

	fakeAssigneeId := common.NewUID(uint32(t.AssigneeId), common.DbTypeTask, 1)
	t.FakeAssigneeId = &fakeAssigneeId

	fakeParentId := common.NewUID(uint32(t.ParentId), common.DbTypeTask, 1)
	t.FakeParentId = &fakeParentId
}
