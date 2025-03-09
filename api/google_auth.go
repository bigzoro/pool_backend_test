package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pool/common/constant"
	"pool/forms"
	"pool/service"
	"pool/utils/rsp"
)

func GetGoogleAuthQR(ctx *gin.Context) {
	qrForm := &forms.GetQRForm{}
	err := ctx.ShouldBindJSON(&qrForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	qrCode, qrUrl, err := service.GetGoogleAuthQR(qrForm.Username)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	var result = map[string]string{
		"qrCode": qrCode,
		"qrUrl":  qrUrl,
	}

	rsp.SuccessResponse(ctx, result)
}

func VerifyGoogleAuthCode(ctx *gin.Context) {
	verifyForm := &forms.VerifyForm{}
	err := ctx.ShouldBindJSON(&verifyForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	verify, err := service.VerifyGoogleAuthCode(verifyForm.Username, verifyForm.Code)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, verify)
}
