package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"pool/global"
	"pool/internal/database/redisx"
	"pool/utils"
	"time"
)

func SendVerificationCode(to string) error {
	// todo: 判断邮箱是否在黑名单中
	// 生成验证码
	code := generateVerificationCode()

	message, err := sendVerificationCode(to, code)
	if err != nil {
		return err
	}

	fmt.Println(message)

	// 保存到 redis 中
	err = redisx.SetEX(context.Background(), to, code, time.Duration(global.Config.RedisConfig.Expire)*time.Second)
	if err != nil {
		return errors.New(fmt.Sprintf("save sms code error: %v", err))
	}

	return nil
}

// sendVerificationCode 发送验证代码到指定的邮箱。
// 参数 to: 邮件接收人的邮箱地址。
// 参数 code: 需要发送的验证代码。
// 返回值 error: 发送过程中遇到的任何错误。
func sendVerificationCode(to string, code string) (string, error) {
	// 发送的代码
	return utils.AwsSendEmail(to, code)
}

// 随机生成一个6位数的验证码。
func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	return code
}
