package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"math/big"
	"net/http"
	"pool/global"
	"pool/models"
	"strconv"
	"strings"
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

	// 过滤掉 pool_name 为 "NETWORK" 的矿池
	var filteredPools []*models.Pool
	for _, pool := range pools {
		if !strings.EqualFold(pool.PoolName, "NETWORK") { // 忽略大小写比较
			filteredPools = append(filteredPools, pool)
		}
	}

	// 如果过滤后没有剩余数据，直接返回
	if len(filteredPools) == 0 {
		log.Println("⚠️  没有需要插入的矿池数据 (所有数据被过滤)")
		return nil
	}

	var savedPools []*models.Pool
	// 修改数据，保存前七个，剩余的合并到一起
	var frontRadio float64
	if len(filteredPools) >= 7 {
		for i := 0; i < len(filteredPools); i++ {
			if i <= 6 {
				savedPools = append(savedPools, filteredPools[i])
				parseFloat, err := strconv.ParseFloat(filteredPools[i].Ratio, 64)
				if err != nil {
					continue
				}
				frontRadio += parseFloat
			}
		}
	}

	otherRadio := 1 - frontRadio

	bigHashFloat := new(big.Float)
	bigHashFloat.SetString(filteredPools[0].Hashps)
	bigRadioFloat := new(big.Float)
	bigRadioFloat.SetString(filteredPools[0].Ratio)
	allHashBigFloat := new(big.Float).Quo(bigHashFloat, bigRadioFloat)
	otherRadioStr := strconv.FormatFloat(otherRadio, 'f', 4, 64)
	bigOtherRadioFloat := new(big.Float)
	bigOtherRadioFloat.SetString(otherRadioStr)
	bigOtherHashFloat := new(big.Float).Mul(allHashBigFloat, bigOtherRadioFloat)

	// 计算全部的 hash
	combinedPools := &models.Pool{
		Hashps:   bigOtherHashFloat.Text('f', -1),
		PoolName: "其它",
		Ratio:    strconv.FormatFloat(otherRadio, 'f', 4, 64),
	}

	savedPools = append(savedPools, combinedPools)

	for _, pool := range savedPools {
		// 检查数据库中是否已存在该矿池
		var existingPool *models.Pool
		if err := global.GormDB.Model(&models.Pool{}).Where("pool_name = ?", pool.PoolName).First(&existingPool).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 不存在则插入
				if err := global.GormDB.Create(pool).Error; err != nil {
					log.Println("❌ 插入失败:", err)
				} else {
					//log.Println("✅ 插入成功:", pool.PoolName)
				}
			} else {
				log.Println("❌ 查询出错:", err)
			}
		} else {
			// 存在则更新
			existingPool.Hashps = pool.Hashps
			existingPool.Ratio = pool.Ratio
			floatRadio, err := strconv.ParseFloat(existingPool.Ratio, 64)
			if err != nil {
				continue
			}

			//	pool.Price = pool.Ratio / (1 - pool.Profit)
			//existingPool.Price = floatRadio / (1 - existingPool.Profit)
			value := floatRadio / (1 - existingPool.Profit)
			formattedValue := strconv.FormatFloat(value, 'f', 4, 64)       // 格式化为 4 位小数的字符串
			existingPool.Price, _ = strconv.ParseFloat(formattedValue, 64) // 转回 float64

			if err := global.GormDB.Save(&existingPool).Error; err != nil {
				log.Println("❌ 更新失败:", err)
			} else {
				//log.Println("✅ 更新成功:", existingPool.PoolName)
			}
		}
	}

	//// 批量插入
	//result := global.GormDB.Model(&models.Pool{}).Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "pool_name"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"hashps", "ratio"}),
	//}).Create(&savedPools)
	//if result.Error != nil {
	//	return fmt.Errorf("❌ 数据库写入失败: %v", result.Error)
	//}
	//
	//log.Printf("✅ 成功插入 %d 条矿池数据", result.RowsAffected)

	return nil
}

func StartPoolScheduler(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	poolInfo, err := fetchPoolInfo()
	if err != nil {
		log.Println("❌ 获取矿池信息失败:", err)
	} else {
		err = savePoolInfo(poolInfo)
		if err != nil {
			log.Printf("❌ 保存矿池信息失败: %v", err)
		}
	}

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
