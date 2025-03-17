package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pool/common/constant"
	"pool/forms"
	"pool/service"
	"pool/utils/rsp"
)

func GetRechargeRecordByUserId(ctx *gin.Context) {
	userRechargeForm := &forms.UserRechargeForm{}
	err := ctx.ShouldBindJSON(&userRechargeForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	totalCount, rechargeRecords, err := service.GetRechargeRecordByUserId(userRechargeForm.UserId, userRechargeForm.Page, userRechargeForm.Size)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	result := map[string]interface{}{
		"total_count":      totalCount,
		"recharge_records": rechargeRecords,
	}

	rsp.SuccessResponse(ctx, result)
}
