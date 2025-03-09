package models

type PoolRecord struct {
	BaseModel
	PoolName string `json:"pool_name"`
	UserId   uint   `json:"user_id"`
}
