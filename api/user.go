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

func EmailRegister(ctx *gin.Context) {
	registerForm := &forms.EmailRegisterForm{}
	err := ctx.ShouldBindJSON(registerForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	err = service.EmailRegister(registerForm)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, "ok")
}

func EmailLogin(ctx *gin.Context) {
	loginForm := &forms.EmailLoginForm{}
	err := ctx.ShouldBindJSON(&loginForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	// 传递给 service 处理
	user, token, err := service.Login(loginForm)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	var resp = UserResp{}
	resp.Token = token
	resp.Username = user.Username
	resp.Avatar = user.Avatar
	resp.UserId = strconv.Itoa(int(user.ID))

	// 返回数据给前端
	rsp.SuccessResponse(ctx, resp)
}

func UpdatePassword(ctx *gin.Context) {
	updateForm := &forms.UpdatePasswordForm{}
	err := ctx.ShouldBindJSON(&updateForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	err = service.UpdatePassword(updateForm)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, "ok")
}

func UpdateWithdrawPassword(ctx *gin.Context) {
	updateWithdrawForm := &forms.UpdateWithdrawPasswordForm{}
	err := ctx.ShouldBindJSON(&updateWithdrawForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	err = service.UpdateWithdrawPassword(updateWithdrawForm)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, "ok")
}

func UpdateEmail(ctx *gin.Context) {
	updateForm := &forms.UpdateEmailForm{}
	err := ctx.ShouldBindJSON(&updateForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	err = service.UpdateEmail(updateForm)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, "ok")
}

func GetUserInfo(ctx *gin.Context) {
	userInfoForm := &forms.UserInfoForm{}
	err := ctx.ShouldBindJSON(&userInfoForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	user, err := service.GetUserInfo(userInfoForm.UserId)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	var resp = UserInfoResp{}
	resp.Username = user.Username
	resp.Avatar = user.Avatar
	resp.Id = int(user.ID)
	resp.Username = user.Username
	resp.Password = user.Password
	resp.Avatar = user.Avatar
	resp.Sex = user.Sex
	resp.Mobile = user.Mobile
	resp.NickName = user.NickName
	resp.Email = user.Email
	resp.InviteCode = user.InviteCode

	// 返回数据给前端
	rsp.SuccessResponse(ctx, resp)
}
