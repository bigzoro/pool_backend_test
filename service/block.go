package service

import (
	"pool/dao"
	"pool/forms"
	"pool/models"
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
