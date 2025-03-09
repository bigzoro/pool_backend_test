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
	result := map[string]interface{}{
		"blocks":         blocks,
		"purchases":      purchases,
		"pools":          pools,
		"totalBlocks":    totalBlocks,
		"totalPurchase":  totalPurchase,
		"totalPools":     totalPools,
		"availableCount": availableCount,
	}

	rsp.SuccessResponse(ctx, result)
}
