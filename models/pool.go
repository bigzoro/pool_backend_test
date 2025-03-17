package models

type Pool struct {
	BaseModel
	Hashps   string  `json:"hashps"`                      // 算力
	PoolName string  `json:"pool_name" gorm:"unique"`     // 矿池名
	Ratio    string  `json:"ratio"`                       // 算力占比
	Price    float64 `json:"price"`                       // 价格
	Profit   float64 `json:"profit" gorm:"default: 0.09"` // 利润率
}
