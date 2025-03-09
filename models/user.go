package models

type User struct {
	BaseModel
	Username         string `json:"username"`
	Password         string `json:"password"`
	Avatar           string `json:"avatar"`
	Sex              string `json:"sex"`
	Mobile           string `json:"mobile"`
	NickName         string `json:"nickname"`
	Email            string `json:"email"`
	Salt             string `json:"salt"`
	WithdrawPassword string `json:"withdraw_password"`
	WithdrawSalt     string `json:"withdraw_salt"`

	ReferrerID *uint  `gorm:"index"`
	InviteCode string `json:"invite_code"`
}
