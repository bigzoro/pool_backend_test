package models

import "time"

// Referral 推荐关系表
type Referral struct {
	ID         uint `gorm:"primaryKey"`
	ReferrerID int  `gorm:"index"`       // 推荐人ID
	UserID     int  `gorm:"uniqueIndex"` // 被推荐用户ID
	CreatedAt  time.Time
}
