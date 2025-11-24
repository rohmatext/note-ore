package entity

import (
	"time"

	"gorm.io/gorm"
)

type Source struct {
	ID          uint           `gorm:"column:id;primaryKey"`
	Name        string         `gorm:"column:name"`
	PhoneNumber *string        `gorm:"column:phone_number"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Source) TableName() string {
	return "sources"
}
