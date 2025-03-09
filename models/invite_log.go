package models

type InviteLog struct {
	BaseModel
	InviterId  uint   `json:"inviter_id"`
	UserId     uint   `json:"user_id"`
	Username   string `json:"username"`
	InviteCode string `json:"invite_code"`
}
