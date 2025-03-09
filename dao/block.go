package dao

import (
	"errors"
	"pool/global"
	"pool/models"
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
