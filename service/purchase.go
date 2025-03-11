package service

import (
	"fmt"
	"pool/dao"
	"pool/forms"
	"pool/models"
	"strconv"
	"sync"
)

// 初始计算每个矿池的可买入的总额
//func AllPoolPurchaseNumber(poolName string) (map[string]float64, error) {
//	// 先获取所有池子的信息
//	_, pools, err := dao.GetPools()
//	if err != nil {
//		return nil, err
//	}
//
//	// 风险金需要动态变化
//	// 风险金怎么计算：
//	// 基本风险金 + 用户已经买入的金额
//	var total float64
//	purchases, err := dao.GetOtherPurchaseByBlockNumber(10, poolName)
//	if err != nil {
//		return nil, err
//	}
//	for _, purchase := range purchases {
//		total += purchase.Count
//	}
//
//	total += 5000
//
//	// 可用额度 = 价格 * 风险金
//	var availableNumber = make(map[string]float64)
//	for _, pool := range pools {
//		availableNumber[pool.PoolName] = pool.Price * total
//	}
//
//	return availableNumber, nil
//}

func AllPoolPurchaseNumber(poolName string) (map[string]float64, error) {
	// 先获取所有池子的信息
	_, pools, err := dao.GetPools()
	if err != nil {
		return nil, err
	}

	// 可用额度 = 价格 * 风险金
	var availableNumber = make(map[string]float64)

	for _, pool := range pools {
		var total float64

		total += 5000

		// 风险金需要动态变化
		// 风险金怎么计算：
		// 基本风险金 + 用户已经买入的金额
		poolCount, err := dao.GetPoolCount()
		if err != nil {
			return nil, err
		}

		for k, v := range poolCount {
			if k == pool.PoolName {
				total -= v / pool.Price
			} else {
				total += v
			}
		}

		availableNumber[pool.PoolName] = pool.Price * total
	}

	return availableNumber, nil
}

// 定义一个全局的互斥锁
var purchaseLock sync.Mutex

func AddPurchase(purchaseParam forms.PurchasesForm) (map[string][]string, error) {
	iUserId, err := strconv.Atoi(purchaseParam.UserId)
	if err != nil {
		return nil, err
	}

	purchaseResult := make(map[string][]string)

	// 买入的时候需要加锁
	purchaseLock.Lock()
	defer purchaseLock.Unlock()
	// 处理购买逻辑
	// 需要判断可买入的余额是否充足，充足就进行购买
	var purchasesToSave []models.Purchase
	for _, pool := range purchaseParam.Purchases {
		// 获取所有可以购买的数量
		poolPurchaseNumber, err := AllPoolPurchaseNumber(pool.PoolName)
		if err != nil {
			return nil, err
		}

		availableCount, exists := poolPurchaseNumber[pool.PoolName]
		if !exists {
			return nil, fmt.Errorf("pool %s not found", pool.PoolName)
		}
		if availableCount >= float64(pool.Count) {
			purchasesToSave = append(purchasesToSave, models.Purchase{
				UserId:      uint(iUserId),
				PoolName:    pool.PoolName,
				Count:       float64(pool.Count),
				Price:       float64(pool.Price),
				BlockNumber: uint64(pool.BlockNumber),
			})
			// 计入买入成功
			purchaseResult["success"] = append(purchaseResult["success"], pool.PoolName)
		} else {
			// 计入买入失败
			purchaseResult["fail"] = append(purchaseResult["fail"], pool.PoolName)
		}
	}

	// 统一写入数据库，减少事务提交次数
	if len(purchasesToSave) > 0 {
		err = dao.BatchAddPurchase(purchasesToSave)
		if err != nil {
			return nil, fmt.Errorf("failed to save purchases: %w", err)
		}
	}

	return purchaseResult, nil
}

func GetPurchaseByUserId(userId int) ([]*models.Purchase, error) {
	purchases, err := dao.GetPurchaseByUserId(userId)
	if err != nil {
	}

	return purchases, nil
}

// 获取所有的游戏记录
func GetAllPurchases() (int64, []*models.Purchase, error) {
	return dao.QueryPurchase()
}
