package taskmodel

import (
	"github.com/japananh/togo/common"
	"strings"
)

type TaskCreate struct {
	common.SQLModel `json:",inline"`
	CreatedBy       int    `json:"-" gorm:"column:created_by;"`
	AssigneeId      int    `json:"-" gorm:"column:assignee;"`
	ParentId        int    `json:"-" gorm:"column:parent_id;"`
	Title           string `json:"title" form:"title" binding:"required" gorm:"column:title;"`
	Description     string `json:"description" form:"description" binding:"required" gorm:"column:description;"`
	FakeCreatedBy   string `json:"created_by" form:"created_by" gorm:"-"`
	FakeAssigneeId  string `json:"assignee,omitempty" form:"assignee_id" gorm:"-"`
	FakeParentId    string `json:"parent_id,omitempty" form:"parent_id" gorm:"-"`
}

func (TaskCreate) TableName() string {
	return Task{}.TableName()
}

func (tc *TaskCreate) Mask() {
	tc.GenUID(common.DbTypeTask)
}

func (tc *TaskCreate) Validate() error {
	tc.Title = strings.TrimSpace(tc.Title)
	tc.Description = strings.TrimSpace(tc.Description)

	if tc.FakeCreatedBy != "" {
		tc.FakeCreatedBy = strings.TrimSpace(tc.FakeCreatedBy)

		createdBy, err := common.FromBase58(strings.TrimSpace(tc.FakeCreatedBy))
		if err != nil {
			return err
		}

		tc.CreatedBy = int(createdBy.GetLocalID())
	}

	if tc.FakeAssigneeId != "" {
		tc.FakeAssigneeId = strings.TrimSpace(tc.FakeAssigneeId)

		parentId, err := common.FromBase58(strings.TrimSpace(tc.FakeAssigneeId))
		if err != nil {
			return err
		}

		tc.AssigneeId = int(parentId.GetLocalID())
	}

	if tc.FakeParentId != "" {
		tc.FakeParentId = strings.TrimSpace(tc.FakeParentId)

		parentId, err := common.FromBase58(strings.TrimSpace(tc.FakeParentId))
		if err != nil {
			return err
		}

		tc.ParentId = int(parentId.GetLocalID())
	}

	return nil
}
