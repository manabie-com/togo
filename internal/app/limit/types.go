package limit

const (
	ReceiveTaskAction = "receive:task"
)

type (
	GetLimitReq struct {
		TierID int    `json:"tier_id"`
		Action string `json:"action"`
	}

	Limit struct {
		TierID int    `json:"tier_id"`
		Action string `json:"action"`
		Value  int    `json:"value"`
	}
)
