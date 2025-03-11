package models

type Plan struct {
	BaseModel
	UserId   int    `json:"user_id"`
	PlanName string `json:"plan_name"`
}
