package dao

import (
	"pool/global"
	"pool/models"
)

func BatchAddPool(pools []models.Pool) error {
	return global.GormDB.Create(pools).Error
}

func GetPools() (int64, []*models.Pool, error) {
	var pools []*models.Pool
	result := global.GormDB.Model(&models.Pool{}).Find(&pools)
	if result.Error != nil {
		return 0, pools, result.Error
	}

	total := result.RowsAffected
	return total, pools, result.Error
}
