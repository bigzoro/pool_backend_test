package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" gorm:"column:add_time"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}
