package task

import (
	db "github.com/dinhquockhanh/togo/internal/pkg/sql/sqlc"
)

func Convert(t *db.Task) *Task {
	return &Task{
		ID:          t.ID,
		Name:        t.Name,
		Assignee:    t.Assignee.String,
		AssignDate:  t.AssignDate,
		Description: t.Description.String,
		Status:      t.Status,
		Creator:     t.Creator,
		CreatedAt:   t.CreatedAt,
		StartDate:   t.StartDate,
		EndDate:     t.EndDate,
	}
}
