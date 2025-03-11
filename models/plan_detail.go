package models

type PlanDetail struct {
	BaseModel
	PlanId   int     `json:"plan_id"`
	PoolName string  `json:"pool_name"`
	Count    float64 `json:"count"`
}
