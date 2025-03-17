package service

import (
	"errors"
	"gorm.io/gorm"
	"pool/dao"
	"pool/forms"
	"pool/models"
	"strconv"
)

type BlockInfoResponse struct {
	Code int `json:"code"`
	Data struct {
		Code      int `json:"code"`
		Count     int `json:"count"`
		CurrPage  int `json:"curr_page"`
		Data      []forms.BlockInfoResp
		HasNext   bool `json:"has_next"`
		Total     int  `json:"total"`
		TotalPage int  `json:"total_page"`
	} `json:"data"`
	Message string `json:"message"`
}

func BlockInfo() (int, []*models.Block, error) {
	return dao.GetBlocks()
}

func GetBlockByPage(num, size int) ([]*models.Block, error) {
	return dao.GetBlockByPage(num, size)
}

func GetCloseZeroBlockNumber() (int, error) {

	// 获取最接近0点的区块号
	zeroBlockNumber, err := dao.GetCloseZeroBlockNumber()

	// 找不到就获取数据库中最新的区块
	if errors.Is(err, gorm.ErrRecordNotFound) {
		latestBlockHeight, err := dao.GetLatestBlockHeight()
		if err != nil {
			return 0, err
		}

		return latestBlockHeight, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	// 转换成int类型
	intBlockHeight, err := strconv.Atoi(zeroBlockNumber)
	if err != nil {
		return 0, err
	}
	return intBlockHeight, nil
}
