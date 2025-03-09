package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pool/common/constant"
	"pool/forms"
	"pool/service"
	"pool/utils/rsp"
	"strconv"
)

// InviteCode 获取当前用户的邀请码
func InviteCode(ctx *gin.Context) {
	inviteCodeForm := &forms.InviteCodeForm{}
	err := ctx.ShouldBind(&inviteCodeForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	intId, err := strconv.Atoi(inviteCodeForm.UserId)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}
	inviteCode, err := service.GetInviteCodeByUserId(uint(intId))
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, inviteCode)
}
