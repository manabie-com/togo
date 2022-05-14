package taskmodel

type Filter struct {
	CreatedBy      int
	ParentId       int
	AssigneeId     int
	Status         string `json:"status" form:"status"`
	CreatedAt      string `json:"created_at" form:"created_at"`
	FakeCreatedBy  string `json:"created_by" form:"created_by"`
	FakeParentId   string `json:"parent_id" form:"parent_id"`
	FakeAssigneeId string `json:"assignee_id" form:"assignee_id"`
}
