package dao

import (
	"errors"
	"pool/global"
	"pool/models"
)

func GetUnusedGoogleAuth(username string) (*models.GoogleAuth, error) {
	var googleAuth *models.GoogleAuth

	result := global.GormDB.Where(&models.GoogleAuth{Username: username}).Find(&googleAuth)

	if result.RowsAffected == 0 {
		return nil, errors.New("对应的谷歌授权码不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return googleAuth, nil
}

func AddGoogleAuth(googleAuth *models.GoogleAuth) error {
	return global.GormDB.Create(googleAuth).Error
}

func DeleteGoogleAuth(ids []int64) error {
	var err error
	tx := global.GormDB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Where("id in (?)", ids).Delete(&models.GoogleAuth{}).Error
	if err != nil {
		return err
	}

	return nil
}
