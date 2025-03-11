package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pool/common/constant"
	"pool/forms"
	"pool/log"
	"pool/service"
	"pool/utils/rsp"
)

func AddPlan(ctx *gin.Context) {
	addPlanForm := forms.AddPlanForm{}
	if err := ctx.ShouldBind(&addPlanForm); err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	err := service.AddPlan(&addPlanForm)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, "ok")
}

func GetUserPlan(ctx *gin.Context) {
	getUserPlanForm := forms.GetUserPlanForm{}
	if err := ctx.ShouldBind(&getUserPlanForm); err != nil {
		log.SystemLog().Warnf("GetUserPlans err:%v", err)
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	plans, err := service.GetUserPlans(getUserPlanForm.UserId)
	if err != nil {
		log.SystemLog().Warnf("GetUserPlans err:%v", err)
		rsp.FailResponse(ctx, err)
		return
	}

	var resp []UserPlansResp
	for _, plan := range plans {
		resp = append(resp, UserPlansResp{
			PlanId:   int(plan.ID),
			PlanName: plan.PlanName,
		})
	}

	rsp.SuccessResponse(ctx, resp)
}

func GetUserPlanDetails(ctx *gin.Context) {
	getUserPlanDetailsForm := forms.GetUserPlanDetailsForm{}
	if err := ctx.ShouldBind(&getUserPlanDetailsForm); err != nil {
		log.SystemLog().Warnf("GetUserPlanDetails err:%v", err)
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	planDetails, err := service.GetUserPlanDetails(getUserPlanDetailsForm.PlanId)
	if err != nil {
		log.SystemLog().Warnf("GetUserPlanDetails err:%v", err)
		rsp.FailResponse(ctx, err)
		return
	}

	var resp []UserPlanDetailsResp
	for _, planDetail := range planDetails {
		resp = append(resp, UserPlanDetailsResp{
			PoolName: planDetail.PoolName,
			Count:    planDetail.Count,
		})
	}

	rsp.SuccessResponse(ctx, resp)
}

func DeletePlan(ctx *gin.Context) {
	deletePlansForm := forms.DeletePlansForm{}
	if err := ctx.ShouldBind(&deletePlansForm); err != nil {
		log.SystemLog().Warnf("GetUserPlanDetails err:%v", err)
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	err := service.DeletePlan(deletePlansForm.PlansId)
	if err != nil {
		log.SystemLog().Warnf("DeletePlan err:%v", err)
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, "ok")
}
