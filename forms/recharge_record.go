package forms

type UserRechargeForm struct {
	UserId int `json:"user_id"`
	Page   int `json:"page"`
	Size   int `json:"size"`
}
