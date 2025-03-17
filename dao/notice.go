package dao

import (
	"pool/global"
	"pool/models"
)

func AddNotice(notice *models.Notice) error {
	return global.GormDB.Create(notice).Error
}

func GetShowNotice() (*models.Notice, error) {
	var notice *models.Notice

	result := global.GormDB.Where(&models.Notice{IsShow: true}).First(&notice)

	if result.Error != nil {
		return nil, result.Error
	}

	return notice, nil
}

func DeleteNotice(ids []int64) error {
	var err error
	tx := global.GormDB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Where("id in (?)", ids).Delete(&models.Notice{}).Error
	if err != nil {
		return err
	}

	return nil
}
