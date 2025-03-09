package api

import (
	"github.com/gin-gonic/gin"
	"pool/forms"
	"pool/service"
	"pool/utils/rsp"
)

func PurchasePool(ctx *gin.Context) {
	var purchasesPoolForm forms.PurchasesForm
	err := ctx.ShouldBind(&purchasesPoolForm)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	err = service.AddPurchase(purchasesPoolForm)
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, nil)
}

func GetAllPurchases(ctx *gin.Context) {
	_, purchases, err := service.GetAllPurchases()
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, purchases)
}

func GetPurchaseByUserId(ctx *gin.Context) {
	//purchasePoolForm := &forms.PurchaseForm{}
	//err := ctx.ShouldBind(purchasePoolForm)
	//if err != nil {
	//	rsp.FailResponse(ctx, err)
	//	return
	//}
	//
	//purchases, err := service.GetPurchaseByUserId(purchasePoolForm.UserId)
	//if err != nil {
	//	rsp.FailResponse(ctx, err)
	//	return
	//}
	//
	//rsp.SuccessResponse(ctx, purchases)
}
