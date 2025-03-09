package service

import (
	"pool/dao"
	"pool/models"
	"pool/utils"
)

func GetGoogleAuthQR(username string) (string, string, error) {
	googleAuth := utils.NewGoogleAuth()
	secret := googleAuth.GetSecret()

	// 动态码(每隔30s会动态生成一个6位数的数字)
	code, err := googleAuth.GetCode(secret)
	if err != nil {
		return "", "", err
	}

	qrCode := googleAuth.GetQRCode(username, code)
	qrUrl := googleAuth.GetQRCodeUrl(username, secret)

	// 保存信息
	saveData := &models.GoogleAuth{
		Username: username,
		Secret:   secret,
		QrUrl:    qrUrl,
	}
	err = dao.AddGoogleAuth(saveData)
	if err != nil {
		return "", "", err
	}

	return qrCode, qrUrl, nil
}

func VerifyGoogleAuthCode(username, code string) (bool, error) {
	googleAuth := utils.NewGoogleAuth()

	// 获取用户的 secret，万一用户有多个 secret 呢
	unusedGoogleAuth, err := dao.GetUnusedGoogleAuth(username)
	if err != nil {
		return false, err
	}

	return googleAuth.VerifyCode(unusedGoogleAuth.Secret, code)
}
