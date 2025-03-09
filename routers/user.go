package routers

import (
	"github.com/gin-gonic/gin"
	"pool/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		// 注册
		userRouter.POST("register", api.EmailRegister)
		// 登陆
		userRouter.POST("login", api.EmailLogin)
		// 邀请码
		userRouter.POST("invite_code", api.InviteCode)
		// 验证码
		//userRouter.GET("captcha", user.GetCaptcha)
		// 短信
		//userRouter.GET("sms", user.SendSms)
		userRouter.POST("update_password", api.UpdatePassword)
		userRouter.POST("update_withdraw_password", api.UpdateWithdrawPassword)
		userRouter.POST("user_info", api.GetUserInfo)
		userRouter.POST("update_email", api.UpdateEmail)
	}
}
