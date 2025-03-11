package dao

import (
	"pool/global"
	"pool/models"
)

func AddPlan(plan *models.Plan) error {
	return global.GormDB.Create(plan).Error
}

func GetUserPlans(userId int) ([]*models.Plan, error) {
	var plans []*models.Plan
	result := global.GormDB.Model(&models.Plan{}).Where("user_id = ?", userId).Find(&plans)
	if result.Error != nil {
		return plans, result.Error
	}

	return plans, nil
}

func GetUserPlanDetails(planId int) ([]*models.PlanDetail, error) {
	var planDetails []*models.PlanDetail
	result := global.GormDB.Model(&models.PlanDetail{}).Where("plan_id = ?", planId).Find(&planDetails)
	if result.Error != nil {
		return planDetails, result.Error
	}

	return planDetails, nil
}

// 包括刪除计划和计划详情
func DeletePlan(ids []int) error {
	var err error
	tx := global.GormDB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Where("id in (?)", ids).Delete(&models.Plan{}).Error
	if err != nil {
		return err
	}

	err = tx.Where("plan_id in (?)", ids).Delete(&models.PlanDetail{}).Error
	if err != nil {
		return err
	}

	return nil
}
