package dao

import (
	"errors"
	"pool/global"
	"pool/models"
	"time"
)

func GetBlocks() (int, []*models.Block, error) {
	var blocks []*models.Block

	result := global.GormDB.Where(&models.Block{}).Order("height desc").Find(&blocks)

	if result.RowsAffected == 0 {
		return 0, nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return 0, nil, result.Error
	}

	total := len(blocks)
	return total, blocks[:10], nil
}

// GetBlockByPage 分页获取区块数据
func GetBlockByPage(num, size int) ([]*models.Block, error) {
	var blocks []*models.Block
	result := global.GormDB.Scopes(Paginate(num, size)).Order("height desc").Find(&blocks)
	if result.RowsAffected == 0 {
		return nil, errors.New("暂无区块信息")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return blocks, nil
}

// 获取最接近0点的区块号
func GetCloseZeroBlockNumber() (string, error) {
	now := time.Now()
	zeroTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	var closestBlock models.Block
	result := global.GormDB.Model(&models.Block{}).Where("time > ?", zeroTime).Order("time ASC").First(&closestBlock)
	if result.Error != nil {
		return "", result.Error
	}

	return closestBlock.Height, nil
}

// 获取最新的区块
func GetLatestBlockHeight() (int, error) {
	var latestHeight int
	result := global.GormDB.Model(&models.Block{}).Order("height ASC").Select("height").First(&latestHeight)
	if result.Error != nil {
		return 0, result.Error
	}

	return latestHeight, nil
}
