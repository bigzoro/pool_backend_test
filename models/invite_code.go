package models

import "time"

type InviteCode struct {
	BaseModel
	UserId    uint
	Code      string
	Hits      int // 点击次数
	RegNum    int // 注册人数
	CreatedAt time.Time
}
