package service

import (
	"pool/dao"
	"pool/models"
)

func AddNotice(notice *models.Notice) error {
	return dao.AddNotice(notice)
}

func GetShowNotice() (*models.Notice, error) {
	return dao.GetShowNotice()
}

func DeleteNotice(ids []int64) error {
	return dao.DeleteNotice(ids)
}
