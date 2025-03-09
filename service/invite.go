package service

import (
	"errors"
	"gorm.io/gorm"
	"pool/dao"
	"pool/utils"
	"strings"
)

// 获取用户的邀请码
func GetInviteCodeByUserId(userId uint) (string, error) {
	// 获取邀请码
	inviteCode, err := dao.GetInviteCodeByUserId(userId)
	var newInviteCode string
	// 没有邀请码，就创建一个
	if errors.Is(err, gorm.ErrRecordNotFound) {
		for {
			newInviteCode = strings.ToUpper(utils.RandString(8))
			// 判断是否有相同的邀请码
			exist, err := dao.InviteCodeExist(newInviteCode)
			if err != nil {
				return "", err
			}
			if !exist {
				break
			}
		}
		// 保存用户的邀请码
		err := dao.AddInviteCode(userId, newInviteCode)
		if err != nil {
			return "", err
		}

		inviteCode = newInviteCode
	}

	return inviteCode, nil
}

// 获取用户的邀请记录
func InviteRecord() {

}

// 添加邀请码
func AddInviteCode(userId uint, inviteCode string) error {
	return dao.AddInviteCode(userId, inviteCode)
}
