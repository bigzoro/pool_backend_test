package models

type Addresses struct {
	BaseModel
	Address    string `json:"address"`
	UserId     uint   `json:"user_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}
