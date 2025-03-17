package models

type RechargeRecord struct {
	BaseModel
	UserId    int    `json:"user_id"`
	Amount    string `json:"amount"`
	Address   string `json:"address"`
	Signature string `json:"signature" gorm:"unique"`
}
