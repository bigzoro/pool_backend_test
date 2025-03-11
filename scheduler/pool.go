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
	poolApiURL = "https://explorer.coinex.com/res/btc/pools/distribution?page=1&limit=50&period=24h"
)

type HashRateResponse struct {
	Code int `json:"code"`
	Data struct {
		Code      int `json:"code"`
		Count     int `json:"count"`
		CurrPage  int `json:"curr_page"`
		Data      []*models.Pool
		HasNext   bool `json:"has_next"`
		Total     int  `json:"total"`
		TotalPage int  `json:"total_page"`
	} `json:"data"`
	Message string `json:"message"`
}

func fetchPoolInfo() ([]*models.Pool, error) {
	client := http.Client{Timeout: requestTimeout}

	var pools []*models.Pool

	for attempt := 1; attempt <= retryCount; attempt++ {
		resp, err := client.Get(poolApiURL)
		if err != nil {
			log.Printf("⚠️ [第 %d 次尝试] 获取矿池信息失败: %v", attempt, err)
			time.Sleep(retryDelay)
			continue
		}

		// 解析 JSON
		var result HashRateResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Printf("⚠️ [第 %d 次尝试] 解析矿池数据失败: %v", attempt, err)
			time.Sleep(retryDelay)
			continue
		}

		pools = result.Data.Data

		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}

		break
	}

	if len(pools) == 0 {
		return nil, fmt.Errorf("获取区块数据失败，重试 %d 次仍无结果", retryCount)
	}

	return pools, nil
}

func savePoolInfo(pools []*models.Pool) error {
	if len(pools) == 0 {
		return nil
	}

	//
	// todo: 过滤掉第一个
	// radio, 0.38
	//for _, pool := range pools {
	//	pool.Price = pool.Ratio / (1 - pool.Profit)
	//}

	// 批量插入
	result := global.GormDB.Model(&models.Pool{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "pool_name"}},
		DoUpdates: clause.AssignmentColumns([]string{"fee_reward_ratio", "hashps", "ratio"}),
	}).Create(&pools)
	if result.Error != nil {
		return fmt.Errorf("❌ 数据库写入失败: %v", result.Error)
	}

	log.Printf("✅ 成功插入 %d 条矿池数据", result.RowsAffected)

	return nil
}

func StartPoolScheduler(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(syncInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			poolInfo, err := fetchPoolInfo()
			if err != nil {
				log.Println(err)
				continue
			}

			err = savePoolInfo(poolInfo)
			if err != nil {
				log.Printf("❌ 保存矿池信息失败: %v", err)
			}
		case <-ctx.Done():
			log.Println("🛑 定时任务已停止")
			return
		}
	}
}
