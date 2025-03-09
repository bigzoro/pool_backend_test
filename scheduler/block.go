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
	"sync"
	"time"
)

const (
	apiURL         = "https://explorer.coinex.com/res/btc/blocks?limit=10"
	requestTimeout = 10 * time.Second // HTTP 超时
	retryCount     = 3                // 失败重试次数
	retryDelay     = 3 * time.Second  // 重试间隔
	syncInterval   = 10 * time.Minute // 定时任务间隔
)

type BlockInfoResponse struct {
	Code int `json:"code"`
	Data struct {
		Code      int            `json:"code"`
		Count     int            `json:"count"`
		CurrPage  int            `json:"curr_page"`
		Data      []models.Block `json:"data"`
		HasNext   bool           `json:"has_next"`
		Total     int            `json:"total"`
		TotalPage int            `json:"total_page"`
	} `json:"data"`
	Message string `json:"message"`
}

// ====================== 获取区块信息（带重试） ======================
func fetchBlockInfo() ([]models.Block, error) {
	client := &http.Client{Timeout: requestTimeout}

	var blocks []models.Block

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
func saveBlockInfo(blocks []models.Block) error {
	if len(blocks) == 0 {
		return nil
	}

	// 批量插入
	result := global.GormDB.Model(&models.Block{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "height"}},
		DoNothing: true,
		UpdateAll: true,
	}).Create(&blocks)
	if result.Error != nil {
		return fmt.Errorf("❌ 数据库写入失败: %v", result.Error)
	}

	log.Printf("✅ 成功插入 %d 条区块数据", result.RowsAffected)

	return nil
}

func StartBlockScheduler(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

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
