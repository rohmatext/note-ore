package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"column:id;primaryKey"`
	Name        string         `gorm:"column:name"`
	Username    string         `gorm:"column:username"`
	Password    string         `gorm:"column:password"`
	RoleID      uint16         `gorm:"column:role_id"`
	Role        Role           `gorm:"foreignKey:RoleID;references:ID;constraint:;"`
	Tokens      []RefreshToken `gorm:"foreignKey:UserID"`
	Productions []Production   `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (User) TableName() string {
	return "users"
}
