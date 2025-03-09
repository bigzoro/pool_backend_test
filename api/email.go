package api

import (
	"github.com/gin-gonic/gin"
	"pool/service"
	"pool/utils/rsp"
)

type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}

type SendVerificationCodeRequest struct {
	Email string `json:"email"`
}

type LoginByVerificationCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func SendVerificationCode(ctx *gin.Context) {
	var req SendVerificationCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	err := service.SendVerificationCode(req.Email)
	if err != nil {
		panic(err)
	}

	rsp.SuccessResponse(ctx, nil)
}
