package dao

import (
	"pool/global"
	"pool/models"
)

// 增
func AddReferral(referrerId, beReferrerId int) error {
	referral := &models.Referral{
		ReferrerID: referrerId,
		UserID:     beReferrerId,
	}
	return global.GormDB.Create(referral).Error
}

// 查
//
//	func GetUserById(id int) (*plum_model.PlumUser, error) {
//		var user plum_model.PlumUser
//
//		result := global.GormDB.Where("id = ?", id).First(&user)
//		if result.Error != nil {
//			return nil, result.Error
//		}
//
//		return &user, nil
//	}
func GetReferralByBeReferral(beReferral string) {

}

//func GetUserByNickName(nickName string) (*plum_model.PlumUser, error) {
//	var user plum_model.PlumUser
//
//	result := global.GormDB.Where("nick_name = ?", nickName).First(&user)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//
//	return &user, nil
//}
//
//func GetUserByUsername(username string) (*plum_model.PlumUser, error) {
//	var user plum_model.PlumUser
//
//	username = "admin"
//	result := global.GormDB.Where("username = ?", username).First(&user)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//
//	return &user, nil
//}
//
//func UserExist(nickName string) (bool, error) {
//	var user plum_model.PlumUser
//
//	result := global.GormDB.Where("username = ?", nickName).First(&user)
//	if result.Error == gorm.ErrRecordNotFound {
//		return false, nil
//	}
//
//	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
//		return false, result.Error
//	}
//
//	return true, nil
//}
//
//func GetUserByPage(page, size int) (int64, []*plum_model.PlumUser, error) {
//	var users []*plum_model.PlumUser
//
//	result := global.GormDB.Find(&users)
//	if result.Error != nil {
//		return 0, nil, result.Error
//	}
//
//	total := result.RowsAffected
//
//	result = global.GormDB.Scopes(dao.Paginate(page, size)).Find(&users)
//	if result.Error != nil {
//		return 0, nil, result.Error
//	}
//
//	return total, users, nil
//}
//

//
//// 改
//func UpdateUser(newUser *plum_model.PlumUser) error {
//	var err error
//	tx := global.GormDB.Begin()
//	defer func() {
//		if err != nil {
//			tx.Rollback()
//		} else {
//			tx.Commit()
//		}
//	}()
//
//	err = tx.Where("id = ?", newUser.ID).Delete(&plum_model.PlumUser{}).Error
//	if err != nil {
//		return err
//	}
//
//	err = tx.Save(newUser).Error
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// 删
//func DeleteUser(ids []int64) error {
//	var err error
//	tx := global.GormDB.Begin()
//	defer func() {
//		if err != nil {
//			tx.Rollback()
//		} else {
//			tx.Commit()
//		}
//	}()
//	err = tx.Where("id in (?)", ids).Delete(&plum_model.PlumUser{}).Error
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// 查询用户总数
//func GetUserCount() (int64, error) {
//	var totalCount int64
//	err := global.GormDB.Table("users").Count(&totalCount).Error
//
//	return totalCount, err
//}
//
//// 查询今日新增的用户数
//func GetDateUserCount(startTime, endTime time.Time) (int64, error) {
//	var totalCount int64
//	err := global.GormDB.Where("create_time > ?", startTime).Where("create_time < ?", endTime).Error
//
//	return totalCount, err
//}
