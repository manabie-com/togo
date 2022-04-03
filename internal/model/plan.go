package model

const (
	FreemiumPlan = "Freemium"
	SilverPlan   = "Silver"
	GoldPlan     = "Gold"
)

// Plan represents the plan model
type Plan struct {
	Base
	Name     string `json:"name"`
	MaxTasks int    `json:"max_tasks"`
}
