package response

type TaskResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint64 `json:"user_id"`
}

type TaskMetaData struct {
	Total       uint `json:"total_items"`
	PageNums    uint `json:"total_pages"`
	PageCurrent uint `json:"current_page"`
	Limit       uint `json:"item_per_page"`
}

type TaskArrayResponse struct {
	Items     []TaskResponse `json:"data"`
	MetatData TaskMetaData   `json:"meta_data"`
}
