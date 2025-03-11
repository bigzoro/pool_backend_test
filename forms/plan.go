package forms

type PlanDetail struct {
	PoolName string  `json:"pool_name"`
	Count    float64 `json:"count"`
}

type AddPlanForm struct {
	UserId      int          `json:"user_id"`
	PlanName    string       `json:"plan_name"`
	PlanDetails []PlanDetail `json:"plan_details"`
}

type GetUserPlanForm struct {
	UserId int `json:"user_id"`
}

type GetUserPlanDetailsForm struct {
	PlanId int `json:"plan_id"`
}

type DeletePlansForm struct {
	PlansId []int `json:"plans_id"`
}
