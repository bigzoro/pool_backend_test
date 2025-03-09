package models

import "math/big"

type Transfer struct {
	BaseModel
	UserId  uint      `json:"user_id"`
	Amount  big.Float `json:"amount"`
	Address string    `json:"address"`
}
