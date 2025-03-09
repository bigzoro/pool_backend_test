package forms

type HashRateResp struct {
	AvgFee         string `json:"avg_fee"`
	AvgSize        string `json:"avg_size"`
	Count          string `json:"count"`
	FeeRewardRatio string `json:"fee_reward_ratio"`
	Hashps         string `json:"hashps"`
	OrphanCount    string `json:"orphan_count"`
	OrphanRatio    string `json:"orphan_ratio"`
	PoolName       string `json:"pool_name"`
	Ratio          string `json:"ratio"`
	Web            bool   `json:"web"`
}
