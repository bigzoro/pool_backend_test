package models

type Pool struct {
	BaseModel
	AvgFee         string  `json:"avg_fee"`                 // 平均矿工费
	AvgSize        string  `json:"avg_size"`                // 平均块大小(B)
	Count          string  `json:"count"`                   // 块数量
	FeeRewardRatio string  `json:"fee_reward_ratio"`        // 矿工费与块奖励占比
	Hashps         string  `json:"hashps"`                  // 算力
	OrphanCount    string  `json:"orphan_count"`            // 孤块数
	OrphanRatio    string  `json:"orphan_ratio"`            // 孤块率
	PoolName       string  `json:"pool_name" gorm:"unique"` // 矿池名
	Ratio          string  `json:"ratio"`                   // 算力占比
	Price          float64 `json:"price"`                   // 价格
}
