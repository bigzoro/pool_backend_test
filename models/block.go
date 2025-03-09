package models

type Block struct {
	BaseModel
	BlockDifficulty string `json:"block_difficulty"`
	BlockReward     string `json:"block_reward"`
	ChainTag        string `json:"chain_tag"`
	Difficulty      string `json:"difficulty"`
	Hash            string `json:"hash"`
	Height          string `json:"height" gorm:"unique"`
	RelayedBy       string `json:"relayed_by"`
	Size            string `json:"size"`
	Time            string `json:"time"`
	TotalReward     string `json:"total_reward"`
	TxCount         string `json:"tx_count"`
	TxFees          string `json:"tx_fees"`
	Web             bool   `json:"web"`
}
