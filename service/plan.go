package service

import (
	"pool/dao"
	"pool/forms"
	"pool/models"
)

func AddPlan(planForm *forms.AddPlanForm) error {
	// 添加计划
	plan := &models.Plan{
		UserId:   planForm.UserId,
		PlanName: planForm.PlanName,
	}
	err := dao.AddPlan(plan)
	if err != nil {
		return err
	}

	// 添加计划详情
	var planDetails []*models.PlanDetail
	for _, planDetail := range planForm.PlanDetails {
		planDetails = append(planDetails, &models.PlanDetail{
			PlanId:   int(plan.ID),
			PoolName: planDetail.PoolName,
			Count:    planDetail.Count,
		})
	}
	err = dao.BatchAddPlanDetail(planDetails)
	if err != nil {
		return err
	}

	return nil
}

func GetUserPlans(userId int) ([]*models.Plan, error) {
	plans, err := dao.GetUserPlans(userId)
	if err != nil {
		return nil, err
	}

	return plans, nil
}

func GetUserPlanDetails(planId int) ([]*models.PlanDetail, error) {
	planDetails, err := dao.GetUserPlanDetails(planId)
	if err != nil {
		return nil, err
	}

	return planDetails, nil
}

func DeletePlan(plantIds []int) error {
	return dao.DeletePlan(plantIds)
}
