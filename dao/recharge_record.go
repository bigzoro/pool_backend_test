package dao

import (
	"gorm.io/gorm/clause"
	"pool/global"
	"pool/models"
)

func AddRecharge(transfer *models.RechargeRecord) error {
	return global.GormDB.Create(transfer).Error
}

// 不存在则添加
func AddRechargeUnique(recharge *models.RechargeRecord) error {
	return global.GormDB.Clauses(clause.OnConflict{DoNothing: true}).Create(recharge).Error
}

//func GetBlockByPage(num, size int) ([]*models.Block, error) {
//	var blocks []*models.Block
//	result := global.GormDB.Scopes(Paginate(num, size)).Order("height desc").Find(&blocks)
//	if result.RowsAffected == 0 {
//		return nil, errors.New("暂无区块信息")
//	}
//	if result.Error != nil {
//		return nil, result.Error
//	}
//
//	return blocks, nil
//}

func GetRechargeRecordByUserId(userId, page, size int) (int64, []*models.RechargeRecord, error) {
	var recharges []*models.RechargeRecord
	result := global.GormDB.Model(&models.RechargeRecord{}).Scopes(Paginate(page, size)).Where("user_id = ?", userId).Find(&recharges)
	if result.Error != nil {
		return 0, nil, result.Error
	}

	totalCount := result.RowsAffected
	return totalCount, recharges, nil
}
