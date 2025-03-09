package forms

type Purchase struct {
	// 用户 ID
	UserId int `json:"user_id"`
	// 矿池 ID
	PoolId uint `json:"pool_id"`
	// 购买数量
	Count       float32 `json:"num"`
	PoolName    string  `json:"pool_name"`
	Price       float32 `json:"price"`
	BlockNumber uint    `json:"block_number"`
}

type PurchasesForm struct {
	UserId    string     `json:"user_id"`
	Purchases []Purchase `json:"purchases"`
}
