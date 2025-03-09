package api

import (
	"github.com/gin-gonic/gin"
	"pool/service"
	"pool/utils/rsp"
)

func HashRate(ctx *gin.Context) {
	//poolForm := &forms.PoolForm{}
	//err := ctx.ShouldBindJSON(&poolForm)
	//if err != nil {
	//	rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
	//	return
	//}

	// 传递给 service 处理
	hashRates, err := service.HashRate()
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	//var resp = UserResp{}
	//resp.Token = token
	//resp.Username = user.Username
	//resp.Avatar = user.Avatar

	// 返回数据给前端
	rsp.SuccessResponse(ctx, hashRates)
}

func GetPools(ctx *gin.Context) {
	_, pools, err := service.GetPools()
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, pools)
}
