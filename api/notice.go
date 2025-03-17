package api

import (
	"github.com/gin-gonic/gin"
	"pool/service"
	"pool/utils/rsp"
)

func AddNotice(ctx *gin.Context) {

	err := service.AddNotice(nil)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}
}

func GetShowNotice(ctx *gin.Context) {
	notice, err := service.GetShowNotice()
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, notice)
}

func DeleteNotice(ctx *gin.Context) {
	err := service.DeleteNotice(nil)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, nil)
}
