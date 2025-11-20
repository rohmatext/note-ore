package entity

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Token     string    `gorm:"column:token"`
	UserID    uint      `gorm:"column:user_id"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
	ExpiredAt time.Time `gorm:"column:expired_at"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
