package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pool/common/constant"
	"pool/forms"
	"pool/service"
	"pool/utils/rsp"
)

func BlockInfo(ctx *gin.Context) {
	//poolForm := &forms.PoolForm{}
	//err := ctx.ShouldBindJSON(&poolForm)
	//if err != nil {
	//	rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
	//	return
	//}

	// 传递给 service 处理
	_, blockInfo, err := service.BlockInfo()
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	//var resp = UserResp{}
	//resp.Token = token
	//resp.Username = user.Username
	//resp.Avatar = user.Avatar

	// 返回数据给前端
	rsp.SuccessResponse(ctx, blockInfo)
}

func GetBlocksByPage(ctx *gin.Context) {
	params := forms.QueryBlockForms{}
	if err := ctx.ShouldBind(&params); err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	blocks, err := service.GetBlockByPage(params.Page, params.Size)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, blocks)
}
