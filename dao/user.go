package dao

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"pool/global"
	"pool/models"
)

func GetUserByMobile(mobile string) (*models.User, error) {
	var user models.User

	result := global.GormDB.Where(&models.User{Mobile: mobile}).First(&user)

	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByName(name string) (*models.User, error) {
	var user models.User

	result := global.GormDB.Where(&models.User{Username: name}).First(&user)

	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByInviteCode(inviteCode string) (*models.User, error) {
	var user models.User
	result := global.GormDB.Where(&models.User{InviteCode: inviteCode}).First(&user)
	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := global.GormDB.Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func UserExist(nickName string) (bool, error) {
	var user models.User

	result := global.GormDB.Where("username = ?", nickName).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}

	return true, nil
}

func AddUser(newUser *models.User) error {
	return global.GormDB.Create(newUser).Error
}

//func GetUserById(id int64) (*models.User, error) {
//	var user models.User
//	result := global.GormDB.Where(&models.User{ID: int(id)}).First(&user)
//	if result.RowsAffected == 0 {
//		return nil, errors.New("用户不存在")
//	}
//	if result.Error != nil {
//		return nil, result.Error
//	}
//
//	return &user, nil
//}

func GetAllUserByPage(num, size int) ([]*models.User, error) {
	var users []*models.User
	result := global.GormDB.Scopes(Paginate(num, size)).Find(&users)
	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func CreateUser(username string) error {
	var user models.User
	result := global.GormDB.Where(&models.User{Username: username}).First(&user)
	if result.RowsAffected == 1 {
		return status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	result = global.GormDB.Create(&user)
	if result.Error != nil {
		return status.Errorf(codes.Internal, result.Error.Error())
	}

	return nil
}

func UpdatePassword(userId int, newPassword string) error {
	result := global.GormDB.Model(&models.User{}).Where("id = ?", userId).Update("password", newPassword)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateWithdrawPassword(userId int, newWithdrawPassword string) error {
	result := global.GormDB.Model(&models.User{}).Where("id = ?", userId).Update("withdraw_password", newWithdrawPassword)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateEmail(userId int, newEmail string) error {
	result := global.GormDB.Model(&models.User{}).Where("id = ?", userId).Update("email", newEmail)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func QueryUserById(userId int) (*models.User, error) {
	var user models.User
	result := global.GormDB.Model(&models.User{}).Where("id = ?", userId).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
