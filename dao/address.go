package dao

import (
	"errors"
	"pool/global"
	"pool/models"
)

func GetAddressByUserId(userId uint) ([]*models.Addresses, error) {
	var addresses []*models.Addresses

	result := global.GormDB.Where(&models.Addresses{UserId: userId}).First(&addresses)

	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return addresses, nil
}

func AddUserAddress(address *models.Addresses) error {
	return global.GormDB.Create(address).Error
}

func DeleteAddress(ids []int64) error {
	var err error
	tx := global.GormDB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Where("id in (?)", ids).Delete(&models.Addresses{}).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUserIdByAddress(address string) (int, error) {
	var userId int

	result := global.GormDB.Model(&models.Addresses{}).Where(&models.Addresses{Address: address}).Select("user_id").First(&userId)

	if result.RowsAffected == 0 {
		return 0, errors.New("用户不存在")
	}
	if result.Error != nil {
		return 0, result.Error
	}

	return userId, nil
}
