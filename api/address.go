package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pool/common/constant"
	"pool/forms"
	"pool/models"
	"pool/service"
	"pool/utils/rsp"
)

func GetAddressByUserId(ctx *gin.Context) {
	addressForm := &forms.AddressForm{}
	err := ctx.ShouldBindJSON(&addressForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	addresses, err := service.GetAddressByUserId(uint(addressForm.UserId))
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, addresses)
}

func AddUserAddress(ctx *gin.Context) {
	addressForm := &forms.AddressForm{}
	err := ctx.ShouldBindJSON(&addressForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}
	err = service.AddUserAddress(&models.Addresses{})
	if err != nil {
		rsp.FailResponse(ctx, err)
		return
	}

	rsp.SuccessResponse(ctx, nil)
}

func DeleteAddress(ctx *gin.Context) {
	addressForm := &forms.AddressForm{}

	err := ctx.ShouldBindJSON(&addressForm)
	if err != nil {
		rsp.FailResponse(ctx, errors.New(constant.ErrorParameter))
		return
	}

	//err = service.DeleteAddress(addressForm.Id)
	//if err != nil {
	//	rsp.FailResponse(ctx, err)
	//	return
	//}
	//

	rsp.SuccessResponse(ctx, nil)
}
