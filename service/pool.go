package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"pool/dao"
	"pool/forms"
	"pool/models"
	"strconv"
)

type HashRateResponse struct {
	Code int `json:"code"`
	Data struct {
		Code      int `json:"code"`
		Count     int `json:"count"`
		CurrPage  int `json:"curr_page"`
		Data      []HashRates
		HasNext   bool `json:"has_next"`
		Total     int  `json:"total"`
		TotalPage int  `json:"total_page"`
	} `json:"data"`
	Message string `json:"message"`
}

type HashRates struct {
	AvgFee         string `json:"avg_fee"`
	AvgSize        string `json:"avg_size"`
	Count          string `json:"count"`
	FeeRewardRatio string `json:"fee_reward_ratio"`
	Hashps         string `json:"hashps"`
	OrphanCount    string `json:"orphan_count"`
	OrphanRatio    string `json:"orphan_ratio"`
	PoolName       string `json:"pool_name"`
	Ratio          string `json:"ratio"`
	Web            bool   `json:"web"`
}

func HashRate() ([]forms.HashRateResp, error) {
	url := "https://explorer.coinex.com/res/btc/pools/distribution?page=1&limit=50&period=24h"

	// Create a new GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var response HashRateResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	var pools []models.Pool

	for _, pool := range response.Data.Data {
		intCount, err := strconv.ParseInt(pool.Count, 10, 64)
		if err != nil {
			panic(err)
		}
		newPool := models.Pool{
			AvgFee:         pool.AvgFee,
			AvgSize:        pool.AvgSize,
			Count:          strconv.FormatInt(intCount, 10),
			FeeRewardRatio: pool.FeeRewardRatio,
			Hashps:         pool.Hashps,
			OrphanCount:    pool.OrphanCount,
			OrphanRatio:    pool.OrphanRatio,
			PoolName:       pool.PoolName,
			Ratio:          pool.Ratio,
		}
		pools = append(pools, newPool)
	}

	err = dao.BatchAddPool(pools)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func GetPools() (int64, []*models.Pool, error) {
	return dao.GetPools()
}
