package service

import (
	"pool/dao"
	"pool/models"
)

func GetRechargeRecordByUserId(userId, page, size int) (int64, []*models.RechargeRecord, error) {
	return dao.GetRechargeRecordByUserId(userId, page, size)
}
