package dao

import (
	"pool/global"
	"pool/models"
)

func AddTransfer(transfer *models.Transfer) error {
	return global.GormDB.Create(transfer).Error
}

//func GetAddressByUserId(userId uint) ([]*models.Addresses, error) {
//	var addresses []*models.Addresses
//
//	result := global.GormDB.Where(&models.Addresses{UserId: userId}).First(&addresses)
//
//	if result.RowsAffected == 0 {
//		return nil, errors.New("用户不存在")
//	}
//	if result.Error != nil {
//		return nil, result.Error
//	}
//
//	return addresses, nil
//}

//func DeleteAddress(ids []int64) error {
//	var err error
//	tx := global.GormDB.Begin()
//	defer func() {
//		if err != nil {
//			tx.Rollback()
//		} else {
//			tx.Commit()
//		}
//	}()
//	err = tx.Where("id in (?)", ids).Delete(&models.Addresses{}).Error
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func GetUserIdByAddress(address string) (uint, error) {
//	var userId uint
//
//	result := global.GormDB.Where(&models.Addresses{Address: address}).First(&userId)
//
//	if result.RowsAffected == 0 {
//		return 0, errors.New("用户不存在")
//	}
//	if result.Error != nil {
//		return 0, result.Error
//	}
//
//	return userId, nil
//}
//
