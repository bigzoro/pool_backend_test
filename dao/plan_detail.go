package dao

import (
	"gorm.io/gorm"
	"pool/global"
	"pool/models"
)

func AddPlanDetail(planDetail *models.PlanDetail) error {
	return global.GormDB.Create(planDetail).Error
}

func BatchAddPlanDetail(planDetails []*models.PlanDetail) error {
	if len(planDetails) == 0 {
		return nil // 没有数据可插入，直接返回
	}

	err := global.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&planDetails).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}
