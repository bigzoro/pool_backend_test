package dao

import (
	"errors"
	"gorm.io/gorm"
	"pool/global"
	"pool/models"
)

func AddPurchase(purchase *models.Purchase) error {
	return global.GormDB.Create(purchase).Error
}

func QueryPurchase() (int64, []*models.Purchase, error) {
	var purchases []*models.Purchase

	result := global.GormDB.Model(&models.Purchase{}).Find(&purchases)

	if result.Error != nil {
		return 0, nil, result.Error
	}

	total := result.RowsAffected

	return total, purchases, nil
}

func BatchAddPurchase(purchases []models.Purchase) error {
	if len(purchases) == 0 {
		return nil // 没有数据可插入，直接返回
	}

	// 这里假设 db 是你的数据库连接实例
	err := global.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&purchases).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func GetPurchaseByUserId(userId int) ([]*models.Purchase, error) {
	var purchases []*models.Purchase

	result := global.GormDB.Where(&models.Purchase{UserId: uint(userId)}).Find(&purchases)

	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return purchases, nil
}

func GetPurchaseByBlockNumber(blockNumber int) ([]*models.Purchase, error) {
	var purchases []*models.Purchase

	result := global.GormDB.Where("block_number = ?", blockNumber).Find(&purchases)

	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return purchases, nil
}

func GetOtherPurchaseByBlockNumber(blockNumber int, poolName string) ([]*models.Purchase, error) {
	var purchases []*models.Purchase

	//result := global.GormDB.Where("block_number = ?", blockNumber).Where("pool_name != ?", poolName).Find(&purchases)
	result := global.GormDB.Where("pool_name != ?", poolName).Find(&purchases)

	if result.Error != nil {
		return nil, result.Error
	}

	return purchases, nil
}
