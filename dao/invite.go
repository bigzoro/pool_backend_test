package dao

import (
	"errors"
	"gorm.io/gorm"
	"pool/global"
	"pool/models"
)

func GetInviteCodeByUserId(userId uint) (string, error) {
	var inviteCode string
	result := global.GormDB.Model(&models.InviteCode{}).Where("user_id = ?", userId).Select("code").First(&inviteCode)
	if result.Error != nil {
		return "", result.Error
	}

	return inviteCode, nil
}

func InviteCodeExist(inviteCode string) (bool, error) {
	var user models.User

	result := global.GormDB.Where("code = ?", inviteCode).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}

	return true, nil
}

func AddInviteCode(userId uint, inviteCode string) error {
	return global.GormDB.Create(&models.InviteCode{UserId: userId, Code: inviteCode}).Error
}
