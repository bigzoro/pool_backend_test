package api

type UserResp struct {
	UserId   string `json:"user_id"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Sex      string `json:"sex"`
	NickName string `json:"nickname"`
}

type UserInfoResp struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Sex      string `json:"sex"`
	Mobile   string `json:"mobile"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Salt     string `json:"salt"`

	ReferrerID *uint  `gorm:"index"`
	InviteCode string `json:"invite_code"`
}

type UserPlansResp struct {
	PlanId   int    `json:"plan_id"`
	PlanName string `json:"plan_name"`
}

type UserPlanDetailsResp struct {
	PoolName string  `json:"pool_name"`
	Count    float64 `json:"count"`
}
