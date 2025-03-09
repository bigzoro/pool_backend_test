package rsp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pool/common/constant"
	"reflect"
)

const (
	Code    = "code"
	Message = "message"
	Data    = "data"
	Detail  = "detail"
	Success = "success"
)

// Response 通用的返回结构
type Response struct {
	// 错误码
	Code ErrCode `json:"code"`
	// 消息
	Message string `json:"message"`
	// 错误细节
	Detail string `json:"detail,omitempty"`
	// 返回的结构体
	Data interface{} `json:"data,omitempty"`
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	retH := gin.H{
		Success: "true",
		Code:    ERR_CODE_OK,
		Message: ERR_CODE_OK.String(),
	}

	if isNil(data) {
		ctx.JSON(http.StatusOK, retH)
		return
	}

	retH[Data] = data

	ctx.JSON(http.StatusOK, retH)
}

func FailResponse(ctx *gin.Context, err error) {
	retH := gin.H{
		Code:    400,
		Message: getMsgFromCommonErr(err),
	}
	if len(err.Error()) != 0 {
		retH[Detail] = err.Error()
		ctx.JSON(http.StatusOK, retH)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, retH)
	ctx.Abort()
}

// 判断返回数据是否为空
func isNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}

func getMsgFromCommonErr(err error) string {
	errMsg := constant.ErrMap[err.Error()]
	if len(errMsg) == 0 {
		return constant.ErrMap[constant.ErrorUnknown]
	}

	return errMsg
}
