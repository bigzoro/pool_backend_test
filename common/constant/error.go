package constant

const (
	ErrorInternal            = "internal error"
	ErrorParameter           = "error_parameter"
	ErrorUserNotExist        = "error_user_not_exist"
	ErrorUserExist           = "error_user_exist"
	ErrorPassword            = "error_password"
	ErrorUnknown             = "error_unknown"
	ErrorInsufficientBalance = "error_insufficient_balance"
)

var ErrMap = map[string]string{
	ErrorParameter:           "参数错误",
	ErrorUserNotExist:        "用户不存在",
	ErrorUserExist:           "用户已存在",
	ErrorInternal:            "系统错误",
	ErrorPassword:            "密码错误",
	ErrorUnknown:             "未知类型错误",
	ErrorInsufficientBalance: "可用份额不足",
}
