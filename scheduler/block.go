package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"pool/global"
	"pool/models"
	"strconv"
	"sync"
	"time"
)

const (
	apiURL         = "https://explorer.coinex.com/res/btc/blocks?limit=10"
	requestTimeout = 10 * time.Second     // HTTP 超时
	retryCount     = 3                    // 失败重试次数
	retryDelay     = 3 * time.Second      // 重试间隔
	syncInterval   = 10 * time.Minute * 6 // 定时任务间隔
	//syncInterval = time.Second // 定时任务间隔
)

type MockBlock struct {
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
}

type BlockInfoResponse struct {
	Code int `json:"code"`
	Data struct {
		Code      int         `json:"code"`
		Count     int         `json:"count"`
		CurrPage  int         `json:"curr_page"`
		Data      []MockBlock `json:"data"`
		HasNext   bool        `json:"has_next"`
		Total     int         `json:"total"`
		TotalPage int         `json:"total_page"`
	} `json:"data"`
	Message string `json:"message"`
}

// ====================== 获取区块信息（带重试） ======================
func fetchBlockInfo() ([]MockBlock, error) {
	client := &http.Client{Timeout: requestTimeout}

	var blocks []MockBlock

	for attempt := 1; attempt <= retryCount; attempt++ {
		resp, err := client.Get(apiURL)
		if err != nil {
			log.Printf("⚠️ [第 %d 次尝试] 获取区块信息失败: %v", attempt, err)
			time.Sleep(retryDelay)
			continue
		}

		// 解析 JSON
		var result BlockInfoResponse

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Printf("⚠️ [第 %d 次尝试] 解析区块数据失败: %v", attempt, err)
			time.Sleep(retryDelay)
			continue
		}

		blocks = result.Data.Data

		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}

		break
	}

	if len(blocks) == 0 {
		return nil, fmt.Errorf("获取区块数据失败，重试 %d 次仍无结果", retryCount)
	}

	return blocks, nil
}

// ====================== 保存区块数据（避免重复） ======================
func saveBlockInfo(blocks []MockBlock) error {
	if len(blocks) == 0 {
		return nil
	}

	var savedBlocks []models.Block
	for _, block := range blocks {
		timestamp, err := strconv.ParseInt(block.Time, 10, 64) // 转换为 int64
		if err != nil {
			return err
		}
		saveTime := time.Unix(timestamp, 0)
		tmp := models.Block{
			BlockDifficulty: block.BlockDifficulty,
			BlockReward:     block.BlockReward,
			ChainTag:        block.ChainTag,
			Difficulty:      block.Difficulty,
			Hash:            block.Hash,
			Height:          block.Height,
			RelayedBy:       block.RelayedBy,
			Size:            block.Size,
			Time:            saveTime,
			TotalReward:     block.TotalReward,
			TxCount:         block.TxCount,
			TxFees:          block.TxFees,
		}

		savedBlocks = append(savedBlocks, tmp)
	}

	// 批量插入
	result := global.GormDB.Model(&models.Block{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "height"}},
		DoNothing: true,
		UpdateAll: true,
	}).Create(&savedBlocks)
	if result.Error != nil {
		return fmt.Errorf("❌ 数据库写入失败: %v", result.Error)
	}

	log.Printf("✅ 成功插入 %d 条区块数据", result.RowsAffected)

	return nil
}

func StartBlockScheduler(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	blockInfo, err := fetchBlockInfo()
	if err != nil {
		log.Println("❌ 获取区块信息失败:", err)
	} else {
		err = saveBlockInfo(blockInfo)
		if err != nil {
			log.Printf("❌ 保存区块信息失败: %v", err)
		}
	}

	ticker := time.NewTicker(syncInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			blockInfo, err := fetchBlockInfo()
			if err != nil {
				log.Println(err)
				continue
			}

			err = saveBlockInfo(blockInfo)
			if err != nil {
				log.Printf("❌ 保存区块信息失败: %v", err)
			}
		case <-ctx.Done():
			log.Println("🛑 定时任务已停止")
			return
		}
	}
}
