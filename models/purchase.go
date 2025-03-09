package models

type Purchase struct {
	BaseModel
	// 用户 ID
	UserId uint `json:"user_id"`
	// 矿池名
	PoolName string `json:"pool_name"`
	// 购买数量
	Count float64 `json:"count"`
	// 购买的价格
	Price float64 `json:"price"`
	// 区块号
	BlockNumber uint64 `json:"block_number"`
}
