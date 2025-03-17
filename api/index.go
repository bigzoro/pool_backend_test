package api

import (
	"github.com/gin-gonic/gin"
	"pool/service"
	"pool/utils/rsp"
)

func IndexData(ctx *gin.Context) {
	totalPools, pools, err := service.GetPools()
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	totalPurchase, purchases, err := service.GetAllPurchases()
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	totalBlocks, blocks, err := service.BlockInfo()
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	availableCount, err := service.AllPoolPurchaseNumber("all")
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	// 获取每天最大可用的区块号
	// 获取当前0点的区块号
	zeroHeight, err := service.GetCloseZeroBlockNumber()
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}
	// 往后叠加120个
	dayMaxBlockHeight := zeroHeight + 120

	result := map[string]interface{}{
		"blocks":            blocks,
		"purchases":         purchases,
		"pools":             pools,
		"totalBlocks":       totalBlocks,
		"totalPurchase":     totalPurchase,
		"totalPools":        totalPools,
		"availableCount":    availableCount,
		"dayMaxBlockHeight": dayMaxBlockHeight,
	}

	rsp.SuccessResponse(ctx, result)
}
