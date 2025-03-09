package service

import (
	"github.com/gagliardetto/solana-go"
	"pool/dao"
	"pool/models"
)

func GetAddressByUserId(userId uint) ([]*models.Addresses, error) {
	return dao.GetAddressByUserId(userId)
}

func AddUserAddress(address *models.Addresses) error {
	// 为用户生成新的地址
	account := solana.NewWallet()
	privateKey := account.PrivateKey
	publicKey := account.PublicKey()
	saveAddress := &models.Addresses{
		UserId:     address.UserId,
		PrivateKey: privateKey.String(),
		PublicKey:  publicKey.String(),
	}

	return dao.AddUserAddress(saveAddress)
}

func DeleteAddress(ids []int64) error {
	return dao.DeleteAddress(ids)
}
